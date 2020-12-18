package main

import (
	"learning_web_framework"
	"log"
	"net/http"
	"time"
)

func onlyForV2() learning_web_framework.HandlerFunc {
	return func(c *learning_web_framework.Context) {
		t := time.Now()
		// c.Fail(500, "Internal Server Error")
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := learning_web_framework.New()
	r.Use(learning_web_framework.Logger())
	r.GET("/", func(c *learning_web_framework.Context) {
		c.HTML(http.StatusOK, "<h1>hello web</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *learning_web_framework.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":99999")
}
