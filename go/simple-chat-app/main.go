package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	log.Println("Server started successfully.")
	// html页面
	http.HandleFunc("/", index)

	// 接受websocket的路由地址
	http.Handle("/ws", websocket.Handler(serveWs))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("static/index.html")
	if err != nil {
		panic(err.Error())
	}

	html.Execute(w, nil)
}

var clients = make(map[*websocket.Conn]struct{})

type Message struct {
	From    string `json:"from"`
	Message string `json:"message"`
}

func serveWs(ws *websocket.Conn) {
	go handleClient(ws)

	select {}
}

func handleClient(ws *websocket.Conn) {
	defer func() {
		delete(clients, ws)
		log.Println("Closing Websocket")
		ws.Close()
	}()

	clients[ws] = struct{}{}

	for {
		var msg Message
		if err := ReadJSON(ws, &msg); err != nil {
			log.Printf("Error in reading json message. Error: %v", err)
			return
		}

		broadcast(ws, msg)
	}
}

// broadcast process the message
func broadcast(ws *websocket.Conn, msg Message) {
	for ws := range clients {
		WriteJSON(ws, msg)
	}
}

func ReadJSON(ws *websocket.Conn, msg *Message) error {
	var buf string
	// Receive receives a text message from client, since buf is string.
	err := websocket.Message.Receive(ws, &buf)
	if err != nil {
		log.Printf("Error in receive message. Error: %v", err)
		return err
	}

	fmt.Printf("recv:%q\n", buf)

	if err := json.Unmarshal([]byte(buf), msg); err != nil {
		log.Printf("Error in receive message json Unmarshal. Error: %v", err)
		return err
	}

	return nil
}

func WriteJSON(ws *websocket.Conn, msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Error in send message json marshal. Error: %v", err)
		return err
	}

	// Send sends a text message to client, since buf is string.
	err = websocket.Message.Send(ws, string(data))
	if err != nil {
		log.Printf("Error in send message. Error: %v", err)
		return err
	}
	fmt.Printf("send:%q\n", data)

	return nil
}
