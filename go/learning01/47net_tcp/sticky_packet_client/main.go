package main

import (
	"fmt"
	"net"
)

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed. err:", err)
		return
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "hello, Hello.How are you?"
		conn.Write([]byte(msg))
	}
}

func main() {
	client()
}
