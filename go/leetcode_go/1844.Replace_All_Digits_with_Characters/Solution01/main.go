package main

import (
	"fmt"
	"strconv"
)

func replaceDigits(s string) string {
	sLen := len(s)
	resultBytes := []byte(s)
	for i := 0; i < sLen-1; i += 2 {
		num, _ := strconv.Atoi(string(s[i+1]))
		resultBytes[i+1] = shift(s[i], uint8(num))
	}

	return string(resultBytes)
}

func shift(b byte, n uint8) byte {
	return b + n
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
