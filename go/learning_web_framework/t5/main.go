package main

import (
	"learning_web_framework"
	"net/http"
)

func main() {
	r := learning_web_framework.New()
	r.GET("/", func(c *learning_web_framework.Context) {
		c.HTML(http.StatusOK, "<h1>Hello learning web framework</h1>")
	})

	r.GET("/hello", func(c *learning_web_framework.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *learning_web_framework.Context) {
		c.JSON(http.StatusOK, learning_web_framework.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
