package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	return strings.Index((s + s)[1:len(s)*2-1], s) != -1
}

func test01() {
	s := "abab"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedSubstringPattern(s))
	fmt.Println()
}

func test02() {
	s := "aba"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedSubstringPattern(s))
	fmt.Println()
}

func test03() {
	s := "abcabcabcabc"
	succResult := true
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedSubstringPattern(s))
	fmt.Println()
}

func test04() {
	s := "abaababaab"
	succResult := true
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedSubstringPattern(s))
	fmt.Println()
}

func test05() {
	s := "abcabcabc"
	succResult := true
	fmt.Println("test05 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedSubstringPattern(s))
	fmt.Println()
}

func test06() {
	s := "abaaba"
	fmt.Println(s + s)
	succResult := true
	fmt.Println("test06 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedSubstringPattern(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
