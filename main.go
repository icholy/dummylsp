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
	return nil, ErrNotImplemented
}

func (s *Server) Initialized(ctx context.Context, params *protocol.InitializedParams) error {
	return ErrNotImplemented
}

func (s *Server) Shutdown(ctx context.Context) error {
	return ErrNotImplemented
}

func (s *Server) Exit(ctx context.Context) error {
	return ErrNotImplemented
}

func main() {
	log.Println("Starting LSP server...")
	handler := protocol.ServerHandler(&DefaultServer{}, nil)
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

	conn.Done()

	if err := conn.Run(context.Background(), handler); err != nil {
		log.Fatal(err)
	}
}
