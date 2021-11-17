package main

import (
	"fmt"
	"strings"
)

func generateTheString(n int) string {
	if isOdd(n) {
		return strings.Repeat("a", n)
	}

	return strings.Repeat("a", n-1) + "b"
}

func isOdd(i int) bool {
	return i&1 == 1
}

func test01() {
	n := 4
	succResult := "pppz"
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", generateTheString(n))
	fmt.Println()
}

func test02() {
	n := 2
	succResult := "xy"
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", generateTheString(n))
	fmt.Println()
}

func test03() {
	n := 7
	succResult := "holasss"
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", generateTheString(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
