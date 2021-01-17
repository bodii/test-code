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

// Option func
type Option func(*Server)

// timeout setting func
func timeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

// maxConns setting func
func maxConns(maxConns int) Option {
	return func(s *Server) {
		s.MaxConns = maxConns
	}
}

func protocol(protocol string) Option {
	return func(s *Server) {
		s.Protocol = protocol
	}
}

// TLS setting func
func TLS(tls *tls.Config) Option {
	return func(s *Server) {
		s.TLS = tls
	}
}

func newServer(addr string, port int, options ...Option) (*Server, error) {
	srv := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConns: 1000,
		TLS:      nil,
	}

	for _, option := range options {
		option(&srv)
	}

	return &srv, nil
}

func test01() {
	s1, _ := newServer("localhost", 1024)
	s2, _ := newServer("localhost", 2048, protocol("udp"))
	s3, _ := newServer("0.0.0.0", 8080, timeout(300*time.Second))

	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", s2)
	fmt.Printf("%v\n", s3)
}

func main() {
	test01()
}
