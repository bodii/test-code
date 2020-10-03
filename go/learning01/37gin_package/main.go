package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/form_post", func(c *gin.Context) {
		id := c.Query("id")
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "annoymous")

		c.JSON(200, gin.H{
			"id": id,
			"status": "posted",
			"message": message,
			"nick": nick,
		})
	})

	router.Run(":8888")
}

