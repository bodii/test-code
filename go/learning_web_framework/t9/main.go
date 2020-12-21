package main

import (
	"fmt"
	"html/template"
	"learning_web_framework"
	"net/http"
	"os"
	"time"
)

type student struct {
	Name string
	Age  int8
}

// FormatAsDate format date return Y-m-d
func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := learning_web_framework.New()
	r.Use(learning_web_framework.Logger())
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	wd, _ := os.Getwd()
	r.LoadHTMLGlob(wd + "/t9/templates/*")
	r.Static("/assets", wd+"/t9/static")

	stu1 := &student{Name: "web", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	r.GET("/", func(c *learning_web_framework.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *learning_web_framework.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", learning_web_framework.H{
			"title":  "gee",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *learning_web_framework.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", learning_web_framework.H{
			"title": "web",
			"now":   time.Date(2020, 10, 01, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
