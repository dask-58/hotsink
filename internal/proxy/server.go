// Package proxy provides a simple TCP proxy server.
package proxy

import (
	"context"
	"log"
	"net"
	"sync"
)

// Server is a TCP proxy server.
type Server struct {
	addr     string
	listener net.Listener
	wg       sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewServer creates a new Server for the given address.
func NewServer(addr string) (*Server, error) {
	ctx, cancel := context.WithCancel(context.Background())

	server := &Server{
		addr:   addr,
		ctx:    ctx,
		cancel: cancel,
	}
	return server, nil
}

// Start runs the server and begins accepting connections.
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}
	s.listener = listener

	log.Printf("Hot Key Proxy Server started on %s", s.addr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			select {
			case <-s.ctx.Done():
				return nil
			default:
				log.Printf("Error accepting connection: %v", err)
				continue
			}
		}

		s.wg.Add(1)
		go s.handleConnection(conn)
	}
}

// Addr returns the address the server is currently listening on.
// It returns nil if the server hasn't started listening yet.
func (s *Server) Addr() net.Addr {
	if s.listener == nil {
		return nil
	}
	return s.listener.Addr()
}

// handleConnection handles a client connection.
func (s *Server) handleConnection(clientConn net.Conn) {
	defer s.wg.Done()
	defer func() {
		if err := clientConn.Close(); err != nil {
			log.Printf("failed to close client connection: %v", err)
		}
	}()

	log.Printf("New connection from %s", clientConn.RemoteAddr())
}

// Shutdown stops the server and closes all resources.
func (s *Server) Shutdown() {
	s.cancel()
	if s.listener != nil {
		if err := s.listener.Close(); err != nil {
			log.Printf("failed to close listener: %v", err)
		}
	}
	s.wg.Wait()
	log.Println("Hot Key Proxy Server stopped")
}
