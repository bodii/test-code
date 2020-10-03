package main

import (
	"fmt"
	"os"
)


func main() {
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)	
	if err != nil {
		fmt.Println("open file failed! err:", err)
		return 
	}

	defer file.Close();	

	str := "hello world\n"
	file.Write([]byte(str))
	file.WriteString(str)



}
