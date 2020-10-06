package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func redirectDemo() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})

	router.Run(":8000")
}

func routerRedirectDemo() {
	router := gin.Default()
	router.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		router.HandleContext(c)
	})

	router.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})

	router.Run(":8000")
}

func routerAllDemo() {
	router := gin.Default()
	router.Any("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"hello": "world"})
	})
	router.Run(":8000")
}

func noRouterDemo() {
	router := gin.Default()
	path, _ := os.Getwd()
	path += "/48gin_router/"
	router.LoadHTMLFiles(path + "views/404.html")

	router.NoRoute(func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{"hello": "ok"})
		c.HTML(http.StatusNotFound, "404.html", nil)
	})

	router.Run(":8000")
}

func routerGroupDemo() {
	router := gin.Default()
	userGroup := router.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"page": "user/index"}) })
		userGroup.GET("/login", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"page": "user/login"}) })
	}

	shopGroup := router.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"page": "shop/index"}) })
		shopGroup.GET("/cart", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"page": "shop/cart"}) })
	}

	router.Run(":8000")
}

func routerGroupNestedDemo() {
	router := gin.Default()
	shopGroup := router.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"page": "shop/index"}) })

		// nested
		xx := shopGroup.Group("/xx")
		xx.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"page": "shop/xx/index"}) })
	}

	router.Run(":8000")
}

func main() {
	// redirectDemo()
	// routerRedirectDemo()
	// routerAllDemo()
	// noRouterDemo()
	// routerGroupDemo()
	routerGroupNestedDemo()
}
