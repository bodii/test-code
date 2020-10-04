package main

import (
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func ginRestfulAPIGet() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "GET",
		})
	})
	r.Run(":8000")
}

func ginRestfulAPIPost() {
	r := gin.Default()
	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "POST",
		})
	})
	r.Run(":8000")
}

func ginRestfulAPIPut() {
	r := gin.Default()
	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "PUT",
		})
	})
	r.Run(":8000")
}

func ginRestfulAPIDelete() {
	r := gin.Default()
	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "DELETE",
		})
	})
	r.Run(":8000")
}

func ginLoadHTMLDemo() {
	r := gin.Default()
	path, _ := os.Getwd()
	path += "/48gin/"
	r.LoadHTMLGlob(path + "/templates/**/*")
	// r.LoadHTMLFiles(path+"/templates/posts/index.html", path+"/templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8000")
}

func customizeTemplateDemo() {
	router := gin.Default()
	path, _ := os.Getwd()
	path += "/48gin/"
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	router.Static("/static", path+"/static")
	router.LoadHTMLFiles(path + "/templates/customize/index.tmpl")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", "<a href='http://baidu.com'>百度</a>")
	})

	router.Run(":8000")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.tmpl")
	if err != nil {
		panic(err.Error())
	}

	// templates map
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}

	return r
}

func renderXMLDemo() {
	r := gin.Default()
	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})

	r.GET("/moreXML", func(c *gin.Context) {
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}

		var msg MessageRecord
		msg.Name = "jack"
		msg.Message = "Hello, world."
		msg.Age = 18

		c.XML(http.StatusOK, msg)
	})

	r.Run(":8000")
}

func renderYAMLDemo() {
	r := gin.Default()

	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "YAML", "status": http.StatusOK})
	})

	r.Run(":8000")
}

func main() {
	renderYAMLDemo()
}
