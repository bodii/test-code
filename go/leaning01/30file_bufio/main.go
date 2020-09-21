package main

import (
	"fmt"
	"os"
	"bufio"
)


func main() {
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed! err:", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == os.EOF {
			fmt.Println("read end!")
			break
		}

		if err != nil {
			fmt.Println("read file failed！err:", err)
			return
		}

		fmt.Println(line)
	}
}
