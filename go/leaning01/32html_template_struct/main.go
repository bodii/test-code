package main

import (
	"fmt"
	"html/template"
	"net/http"
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
	tmpl, err := template.ParseFiles(dir + "/32html_template_struct/hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	user := UserInfo {
		Name: "Jack",
		Gender: "man",
		Age: 18,
	}

	tmpl.Execute(w, user)
}


func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}
