package main

import (
	"fmt"
	"strings"
)

func toLowerCase(str string) string {
	result := strings.Builder{}

	for _, v := range str {
		if isUpper(v) {
			v += 32
		}

		result.WriteRune(v)
	}

	return result.String()
}

func isUpper(c rune) bool {
	if 65 <= c && c <= 90 {
		return true
	}

	return false
}

func test01() {
	str := "Hello"
	succResult := "hello"
	fmt.Println("test01 str:=", str)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", toLowerCase(str))
	fmt.Println()
}

func test02() {
	str := "here"
	succResult := "here"
	fmt.Println("test02 str:=", str)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", toLowerCase(str))
	fmt.Println()
}

func test03() {
	str := "LOVELY"
	succResult := "lovely"
	fmt.Println("test03 str:=", str)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", toLowerCase(str))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
