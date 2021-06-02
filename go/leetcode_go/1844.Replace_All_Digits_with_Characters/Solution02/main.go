package main

import (
	"fmt"
)

func replaceDigits(s string) string {
	sLen := len(s)
	resultBytes := []byte(s)
	for i := 0; i < sLen-1; i += 2 {
		resultBytes[i+1] = s[i] + (s[i+1] - '0')
	}

	return string(resultBytes)
}

func test01() {
	s := "a1c1e1"
	succResult := "abcdef"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceDigits(s))
	fmt.Println()
}

func test02() {
	s := "a1b2c3d4e"
	succResult := "abbdcfdhe"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceDigits(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
