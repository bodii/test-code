package main

import (
	"fmt"
	"strings"
)

func checkString(s string) bool {
	return !strings.Contains(s, "ba")
}

func test01() {
	s := "aaabbb"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkString(s))
	fmt.Println()
}

func test02() {
	s := "abab"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkString(s))
	fmt.Println()
}

func test03() {
	s := "bbb"
	succResult := true
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkString(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
