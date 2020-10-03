package main

import (
	"fmt"
	"os"
	"log"
)


func main() {
	log.Println("这是一条日志")
	
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	/*
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
	*/

	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条普通的日志。")
	// 设置前缀
	log.SetPrefix("[log prefix]")
	log.Println("这是一条普通的日志。")

	dir, _ := os.Getwd()
	logFile, err := os.OpenFile(dir + "/36log/test.log", os.O_CREATE | os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}

	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds |log.Ldate)
	log.Println("这是一条很普通的日志.")
	log.SetPrefix("[日志]")
	log.Println("这是一条很普通的日志。")

	logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
	logger.Println("这是自定义的logger记录的日志")
}
