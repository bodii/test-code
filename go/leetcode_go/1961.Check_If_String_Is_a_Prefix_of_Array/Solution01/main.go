package main

import (
	"fmt"
	"strings"
)

func isPrefixString(s string, words []string) bool {
	wordAdd := strings.Builder{}
	sLen := len(s)

	for i := 0; i < len(words); i++ {
		if sLen <= wordAdd.Len() {
			break
		}

		wordAdd.WriteString(words[i])
		if s == wordAdd.String() {
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
