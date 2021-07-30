package main

import (
	"fmt"
)

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	if len(a) > len(b) {
		return len(a)
	}

	return len(b)
}

func test01() {
	a := "aba"
	b := "cdc"
	succResult := 3
	fmt.Println("test01 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findLUSlength(a, b))
	fmt.Println()
}

func test02() {
	a := "aaa"
	b := "bbb"
	succResult := 3
	fmt.Println("test02 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findLUSlength(a, b))
	fmt.Println()
}

func test03() {
	a := "aaa"
	b := "aaa"
	succResult := -1
	fmt.Println("test03 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findLUSlength(a, b))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
