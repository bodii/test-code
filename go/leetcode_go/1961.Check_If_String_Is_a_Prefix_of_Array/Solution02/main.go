package main

import (
	"fmt"
	"strings"
)

func isPrefixString(s string, words []string) bool {
	for i := 1; i <= len(words); i++ {
		if s == strings.Join(words[0:i], "") {
			return true
		}
	}

	return false
}

func test01() {
	s := "iloveleetcode"
	words := []string{"i", "love", "leetcode", "apples"}
	succResult := true
	fmt.Println("test01 s:=", s, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixString(s, words))
	fmt.Println()
}

func test02() {
	s := "iloveleetcode"
	words := []string{"apples", "i", "love", "leetcode"}
	succResult := false
	fmt.Println("test02 s:=", s, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixString(s, words))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
