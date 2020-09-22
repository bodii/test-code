package main

import (
	"fmt"
	"os"
)

func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open file failed! err:", err)
		return 
	}
	defer src.Close()

	dst, err := os.Open(dstName)
	if err != nil {
		fmt.Println("open file failed! err:", err)
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}


func main() {
	_, err := CopyFile("dst.txt", "src.txt") 
	if err != nil {
		fmt.Println("copy file failed, err:", err)
		return
	}

	fmt.Println("copy success!")
}
