package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"
)

/// SSE(服务器发送事件)特点:
/// 单向数据推送，基于HTTP，实现简单。
/// 适用场景:新闻推送、股票更新、系统通知。
/// 优点:实现简单，易于集成，基于HTTP。
/// 缺点:单向通信，断线后需重连。

type SseHandler struct {
	clients map[chan string]struct{}
}

func NewSseHandler() *SseHandler {
	return &SseHandler{
		clients: make(map[chan string]struct{}),
	}
}

func (h *SseHandler) Serve(w http.ResponseWriter, r *http.Request) {
	// 设置SSE必需的HTTP头
	w.Header().Add("Content-Type", "text/event-stream")
	w.Header().Add("Cache-Control", "no-cache")
	w.Header().Add("Connection", "keep-alive")

	// 为每个客户端创建一个channel
	clientChan := make(chan string)
	h.clients[clientChan] = struct{}{}

	// 客户端断开时清理
	defer func() {
		delete(h.clients, clientChan)
		close(clientChan)
	}()

	// 持续监听并推送事件
	for {
		select {
		case msg := <-clientChan:
			// 发送事件数据
			fmt.Fprintf(w, "data: %s\n\n", msg)
			w.(http.Flusher).Flush()
		case <-r.Context().Done():
			// 客户端断开连接
			return
		}
	}
}

// SimulateEvents 模拟周期事件
func (h *SseHandler) SimulateEvents() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for range ticker.C {
		message := fmt.Sprintf("Server time: %s",
			time.Now().Format(time.RFC3339))
		// 广播给所有客户端
		for clientChan := range h.clients {
			select {
			case clientChan <- message:
			default:
				// 跳过阻塞的channel
			}
		}
	}
}

func (h *SseHandler) Home(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./static/index.html")
	err := t.Execute(w, nil)
	if err != nil {
		fmt.Printf("execute file failed err := %v", err)
	}
}

func main() {
	mux := http.NewServeMux()

	sseHandler := NewSseHandler()
	mux.HandleFunc("/", sseHandler.Home)
	mux.HandleFunc("/sse", sseHandler.Serve)

	// 在单独的goroutine中模拟事件
	go sseHandler.SimulateEvents()

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
