package main

import (
	"learning_web_framework"
	"net/http"
)

func main() {
	r := learning_web_framework.Default()
	r.GET("/", func(c *learning_web_framework.Context) {
		c.String(http.StatusOK, "hello web\n")
	})

	r.GET("/panic", func(c *learning_web_framework.Context) {
		names := []string{"web"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
