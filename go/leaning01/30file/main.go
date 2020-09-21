package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("main.go")
	if err != nil {
		fmt.Println("Open file failed! error:", err)
		return
	}
	// 关闭文件
	defer file.Close()

	for {
		// data := make([]byte, 128)
		// count, err := file.Read(data)
		// if err != nil {
		// 	fmt.Println("Read file failed! error:", err)
		// 	return
		// }
		// fmt.Println(count)
		// ==
		var data [128]byte
		count, err := file.Read(data[:])
		if err == io.EOF {
			fmt.Println("read end!")
			return
		}

		if err != nil {
			fmt.Println("Read file failed! error:", err)
			return
		}
		fmt.Printf("读取了%d个字节\n", count)
		fmt.Println(string(data[:count]))
	}
}
