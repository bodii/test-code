package main

import (
	"fmt"
	"net"

	nettools "github.com/bodii/test-code/go/learning01/test/net"
)

func client() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed. err:", err)
		return
	}

	defer conn.Close()

	for i := 0; i < 20; i++ {
		msg := "Hello, Hello. How are you?"
		data, err := nettools.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}

		conn.Write(data)
	}
}

func main() {
	client()
}
