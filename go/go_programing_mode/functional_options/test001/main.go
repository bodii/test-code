package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

// Server struct
type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// NewDefaultServer new default server
func NewDefaultServer(addr string, port int) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
}

// NewTLSServer new tls server
func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
	return &Server{addr, port, "tcp", 30 * time.Second, 100, tls}, nil
}

// NewServerWithTimeout new server with timeout
func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
	return &Server{addr, port, "tcp", timeout, 100, nil}, nil
}

// NewTLSServerWithMaxConnAndTimeout new tls server with maxConns and timeout
func NewTLSServerWithMaxConnAndTimeout(addr string, port int, timeout time.Duration, maxConns int, tls *tls.Config) (*Server, error) {
	return &Server{addr, port, "tcp", timeout, maxConns, tls}, nil
}

func test01() {
	serv, _ := NewDefaultServer("localhost", 9200)
	fmt.Printf("%v\n", serv)
}

func main() {
	test01()
}
