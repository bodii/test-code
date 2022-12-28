package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

var port *int = flag.Int("p", 23456, "Port to listen.")

// copyServer echoes back messages sent from client using io.Copy.
func copyServer(ws *websocket.Conn) {
	fmt.Printf("copyServer %#v\n", ws.Config())
	io.Copy(ws, ws)
	fmt.Println("copyServer finished")
}

// readWriteServer echoes back messages sent from client using Read and Write.
func readWriteServer(ws *websocket.Conn) {
	fmt.Printf("readWriteServer %#v\n", ws.Config())
	for {
		buf := make([]byte, 100)
		// Read at most 100 bytes.  If client sends a message more than
		// 100 bytes, first Read just reads first 100 bytes.
		// Next Read will read next at most 100 bytes.
		n, err := ws.Read(buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf[:n])
		// Write send a message to the client.
		n, err = ws.Write(buf[:n])
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%q\n", buf[:n])
	}
	fmt.Println("readWriteServer finished")
}

// sendRecvServer echoes back text messages sent from client
// using websocket.Message.
func sendRecvServer(ws *websocket.Conn) {
	fmt.Printf("sendRecvServer %#v\n", ws)
	for {
		var buf string
		// Receive receives a text message from client, since buf is string.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%q\n", buf)
		// Send sends a text message to client, since buf is string.
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%q\n", buf)
	}
	fmt.Println("sendRecvServer finished")
}

// sendRecvBinaryServer echoes back binary messages sent from clent
// using websocket.Message.
// Note that chrome supports binary messaging in 15.0.874.* or later.
func sendRecvBinaryServer(ws *websocket.Conn) {
	fmt.Printf("sendRecvBinaryServer %#v\n", ws)
	for {
		var buf []byte
		// Receive receives a binary message from client, since buf is []byte.
		err := websocket.Message.Receive(ws, &buf)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("recv:%#v\n", buf)
		// Send sends a binary message to client, since buf is []byte.
		err = websocket.Message.Send(ws, buf)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", buf)
	}
	fmt.Println("sendRecvBinaryServer finished")
}

type T struct {
	Msg  string
	Path string
}

// jsonServer echoes back json string sent from client using websocket.JSON.
func jsonServer(ws *websocket.Conn) {
	fmt.Printf("jsonServer %#v\n", ws.Config())
	for {
		var msg T
		// Receive receives a text message serialized T as JSON.
		err := websocket.JSON.Receive(ws, &msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("recv:%#v\n", msg)
		// Send send a text message serialized T as JSON.
		err = websocket.JSON.Send(ws, msg)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("send:%#v\n", msg)
	}
}

func MainServer(w http.ResponseWriter, req *http.Request) {
	html, err := template.ParseFiles("./view.html")
	if err != nil {
		panic(err.Error())
	}

	html.Execute(w, nil)
}

func main() {
	flag.Parse()
	http.Handle("/copy", websocket.Handler(copyServer))
	http.Handle("/readWrite", websocket.Handler(readWriteServer))
	http.Handle("/sendRecvText", websocket.Handler(sendRecvServer))
	http.Handle("/sendRecvArrayBuffer", websocket.Handler(sendRecvBinaryServer))
	http.Handle("/sendRecvBlob", websocket.Handler(sendRecvBinaryServer))
	http.Handle("/json", websocket.Handler(jsonServer))
	http.HandleFunc("/", MainServer)
	fmt.Printf("http://localhost:%d/\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		panic("ListenANdServe: " + err.Error())
	}
}
