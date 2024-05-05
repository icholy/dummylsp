package main

import (
	"context"
	"flag"
	"io"
	"log/slog"
	"os"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

type Server struct {
	Log *slog.Logger
	DefaultServer
}

func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	s.Log.Info("Initialize", "params", params)
	return &protocol.InitializeResult{
		ServerInfo: &protocol.ServerInfo{
			Name:    "dummy",
			Version: "v0.0.0",
		},
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncOptions{
				OpenClose: true,
				Change:    protocol.TextDocumentSyncKindIncremental,
			},
		},
	}, nil
}

func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	s.Log.Debug("Initialized", "params", params)
	return nil
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	s.Log.Debug("DidChange", "params", params)
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.Log.Debug("Shutdown")
	return nil
}

func (s *Server) Exit(ctx context.Context) error {
	s.Log.Debug("Exit")
	return nil
}

func main() {
	var logfile string
	level := slog.LevelInfo
	flag.Func("log.level", "log level", func(s string) error {
		return level.UnmarshalText([]byte(s))
	})
	flag.StringVar(&logfile, "log.file", "", "log file to write to")
	flag.Parse()
	output := os.Stderr
	if logfile != "" {
		var err error
		output, err = os.OpenFile(logfile, os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			print(err)
			os.Exit(0)
		}
	}
	log := slog.New(slog.NewTextHandler(output, &slog.HandlerOptions{}))
	log.Info("Starting LSP server...")
	handler := protocol.ServerHandler(&Server{Log: log}, nil)
	stream := jsonrpc2.NewStream(struct {
		io.Reader
		io.Writer
		io.Closer
	}{
		Reader: os.Stdin,
		Writer: os.Stdout,
		Closer: io.NopCloser(nil),
	})
	conn := jsonrpc2.NewConn(stream)
	conn.Go(context.Background(), handler)
	<-conn.Done()
	if err := conn.Err(); err != nil {
		print(err)
		os.Exit(1)
	}
}
