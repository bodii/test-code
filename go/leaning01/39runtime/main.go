package main

import (
	"fmt"
	"runtime"
	"path"
)


func runFunc() {
	// 获取自己
	// pc, file, line, ok := runtime.Caller(0)
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("runtime.Caller failed.")
		return 
	}

	fmt.Println(runtime.FuncForPC(pc).Name())
	fmt.Println(file)
	fmt.Println(path.Base(file))
	fmt.Println(line)
	fmt.Println(ok)
}

func main() {
	runFunc()
}
