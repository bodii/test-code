package main

import (
	"learning_web_framework"
	"net/http"
)

func main() {
	r := learning_web_framework.New()
	r.GET("/", func(c *learning_web_framework.Context) {
		c.HTML(http.StatusOK, "<h1>hello web framework</h1>")
	})

	r.GET("/hello", func(c *learning_web_framework.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *learning_web_framework.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *learning_web_framework.Context) {
		c.JSON(http.StatusOK, learning_web_framework.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")
}
