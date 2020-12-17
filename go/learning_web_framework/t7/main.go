package main

import (
	"learning_web_framework"
	"net/http"
)

func main() {
	r := learning_web_framework.New()

	r.GET("/index", func(c *learning_web_framework.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *learning_web_framework.Context) {
			c.HTML(http.StatusOK, "<h1>hello web</h1>")
		})

		v1.GET("/hello", func(c *learning_web_framework.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *learning_web_framework.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})

		v2.POST("/login", func(c *learning_web_framework.Context) {
			c.JSON(http.StatusOK, learning_web_framework.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})
	}

	r.Run(":9999")
}
