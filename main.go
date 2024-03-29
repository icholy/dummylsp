package main

import (
	"context"
	"io"
	"log"
	"os"

	"go.lsp.dev/jsonrpc2"
	"go.lsp.dev/protocol"
)

type Server struct {
	DefaultServer
}

func (s *Server) Initialize(ctx context.Context, params *protocol.InitializeParams) (*protocol.InitializeResult, error) {
	log.Printf("Initialize: %#v\n", params)
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
	log.Printf("Initialized: %#v\n", params)
	return nil
}

func (s *Server) DidChange(ctx context.Context, params *protocol.DidChangeTextDocumentParams) error {
	log.Printf("DidChange: %#v\n", params)
	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutdown")
	return nil
}

func (s *Server) Exit(ctx context.Context) error {
	log.Println("Exit")
	return nil
}

func main() {
	log.Println("Starting LSP server...")
	handler := protocol.ServerHandler(&Server{}, nil)
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
		log.Fatal(err)
	}
}
