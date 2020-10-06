package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// statTimeCost
func statTimeCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 通过在请求上下文中设置值，后续的处理函数能取到该值
		c.Set("name", "jack")
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		// c.Abort()

		cost := time.Since(start)
		log.Println(cost)

	}
}

func ginHookDemo() {
	// 新建一个没有任何默认中间件的路由
	router := gin.New()
	// 注册一个全局中间件
	router.Use(statTimeCost())

	router.GET("/test", func(c *gin.Context) {
		name := c.MustGet("name").(string)
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	router.GET("/test2", statTimeCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string)
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})

	// shopGroup := router.Group("/shop", statTimeCost())
	// {
	// 	shopGroup.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "ok"}) })
	// }

	shopGroup := router.Group("/shop")
	shopGroup.Use(statTimeCost())
	{
		shopGroup.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "ok"}) })

	}

	router.Run(":8000")
}

func main() {
	ginHookDemo()
}
