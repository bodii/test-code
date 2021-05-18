package main

import (
	"fmt"
)

func toLowerCase(str string) string {
	strBytes := []byte(str)
	for i := 0; i < len(strBytes); i++ {
		if isUpper(strBytes[i]) {
			strBytes[i] += 32
		}
	}

	return string(strBytes)
}

func isUpper(v byte) bool {
	return 65 <= v && v <= 90
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
