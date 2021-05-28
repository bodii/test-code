package main

import (
	"fmt"
)

func removePalindromeSub(s string) int {
	sLen := len(s)
	if sLen < 2 {
		return sLen
	}

	for i, j := 0, sLen-1; i < j; i++ {
		if s[i] != s[j] {
			return 2
		}
		j--
	}

	return 1
}

func test01() {
	s := "ababa"
	succResult := 1
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", removePalindromeSub(s))
	fmt.Println()
}

func test02() {
	s := "abb"
	succResult := 2
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", removePalindromeSub(s))
	fmt.Println()
}

func test03() {
	s := "baabb"
	succResult := 2
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", removePalindromeSub(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
