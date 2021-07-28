package main

import (
	"fmt"
)

func replaceSpace(s string) string {
	result := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			result = append(result, '%', '2', '0')
		} else {
			result = append(result, s[i])
		}
	}

	return string(result)
}

func test01() {
	s := "We are happy."
	succResult := "We%20are%20happy."
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceSpace(s))
	fmt.Println()
}

func main() {
	test01()
}
