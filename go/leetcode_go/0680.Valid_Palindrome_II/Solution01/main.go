package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
		l++
		r--
	}

	return true
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func test01() {
	s := "aba"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", validPalindrome(s))
	fmt.Println()
}

func test02() {
	s := "abca"
	succResult := true
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", validPalindrome(s))
	fmt.Println()
}

func test03() {
	s := "abc"
	succResult := false
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", validPalindrome(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
