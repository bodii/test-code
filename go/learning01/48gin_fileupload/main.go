package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func singleFileUpload() {
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		// single file
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Println(file.Filename)
		dst := fmt.Sprintf("/tmp/%s", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})

	})

	path, _ := os.Getwd()
	path += "/48gin_fileupload/"
	router.LoadHTMLFiles(path + "upload_file.html")

	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_file.html", gin.H{})
	})

	router.Run(":8000")
}

func multipartFileUploadDemo() {
	router := gin.Default()
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("/tmp/%s_%d", file.Filename, index)
			c.SaveUploadedFile(file, dst)
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})

	path, _ := os.Getwd()
	path += "/48gin_fileupload/"

	router.LoadHTMLFiles(path + "upload_files.html")
	router.GET("/uploads", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload_files.html", gin.H{})
	})

	router.Run(":8000")

}

func main() {
	// singleFileUpload()
	multipartFileUploadDemo()
}
