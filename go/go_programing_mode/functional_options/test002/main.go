package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

// ServerConfig server config struct
type ServerConfig struct {
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}

// Server server struct
type Server struct {
	Addr string
	Port int
	Conf *ServerConfig
}

// NewServer new Server
func NewServer(add string, port int, conf *ServerConfig) (*Server, error) {
	return &Server{add, port, conf}, nil
}

func test01() {
	serv1, _ := NewServer("localhost", 9000, nil)
	fmt.Printf("%v\n", serv1)

	conf := ServerConfig{Protocol: "tcp", Timeout: 60 * time.Second}
	serv2, _ := NewServer("localhost", 9000, &conf)
	fmt.Printf("%v\n", serv2)
}

func main() {
	test01()
}
