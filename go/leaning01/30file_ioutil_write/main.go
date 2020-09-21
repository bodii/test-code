package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	err := ioutil.WriteFile("test.txt", []byte("hello go!\n"), 0666)
	if err != nil {
		fmt.Println("file write failed! err:", err)
		return
	}
}

