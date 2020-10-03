package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var s string
	fmt.Print("请输入内容:")
	reader := bufio.NewReader(os.Stdin)
	s, _ = reader.ReadString('\n')
	fmt.Println("你输入的内容是:", s)
}
