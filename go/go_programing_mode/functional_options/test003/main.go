package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

// Server server struct
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// ServerBuilder server builder struct
type ServerBuilder struct {
	Server
}

func (s *ServerBuilder) create(addr string, port int) *ServerBuilder {
	s.Server.Addr = addr
	s.Server.Port = port
	return s
}

func (s *ServerBuilder) withProtocol(protocol string) *ServerBuilder {
	s.Server.Protocol = protocol
	return s
}

func (s *ServerBuilder) withMaxConn(maxConns int) *ServerBuilder {
	s.Server.MaxConns = maxConns
	return s
}

func (s *ServerBuilder) withTimeOut(timeout time.Duration) *ServerBuilder {
	s.Server.Timeout = timeout
	return s
}

func (s *ServerBuilder) withTLS(tls *tls.Config) *ServerBuilder {
	s.TLS = tls
	return s
}

func (s *ServerBuilder) build() Server {
	return s.Server
}

func test01() {
	s := ServerBuilder{}
	server := s.create("127.0.0.1", 8000).
		withProtocol("udp").
		withMaxConn(1024).
		withTimeOut(30 * time.Second).
		build()

	fmt.Printf("%#v\n", server)
}

func main() {
	test01()
}
