package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("file open failed! err:", err)
		return 
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0;i < 10;i++ {
		// 将数据先写入缓存
		writer.WriteString(fmt.Sprintf("hello, %d\n", i))
	}

	// 将缓存中的内容写入文件
	writer.Flush()
}
