package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()
	router.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			// 设置cookie
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Println("cookie value: %s \n", cookie)
	})

	router.Run(":8888")
}
