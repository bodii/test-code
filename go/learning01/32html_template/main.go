package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)


func sayHello(w http.ResponseWriter, r *http.Request) {
	// 当前目录
	wd, _ := os.Getwd()
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles(wd + "/32html_template/hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	tmpl.Execute(w, "world")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("HTTP server failed, err:", err)
		return
	}
}
