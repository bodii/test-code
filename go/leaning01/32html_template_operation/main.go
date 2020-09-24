package main


import (
	"fmt"
	"net/http"
	"html/template"
	"os"
)

type UserInfo struct {
	Name string
	Gender string
	Age int
}


func sayHello(w http.ResponseWriter, r *http.Request) {
	// 获取当前目录
	dir, _ := os.Getwd()
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles( dir + "/32html_template_operation/")
	if err != nil {
		fmt.Println("create template failed! err:", err)
		return
	}

	user := UserInfo {
		Name: "Tom",
		Gender: "man",
		Age: 18,
	}

	tepl.Execute(w, user)
}
