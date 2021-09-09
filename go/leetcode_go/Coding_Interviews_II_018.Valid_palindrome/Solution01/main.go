package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	len := len(s)

	bytes := make([]byte, 0)
	bytesLen := 0
	isAllSpace := true
	for i := 0; i < len; i++ {
		if 'A' <= s[i] && 'Z' >= s[i] {
			bytes = append(bytes, s[i]+('a'-'A'))
			bytesLen++
		} else if ('a' <= s[i] && 'z' >= s[i]) || ('0' <= s[i] && '9' >= s[i]) {
			bytes = append(bytes, s[i])
			bytesLen++
		}

		if s[i] != ' ' {
			isAllSpace = false
		}
	}

	if isAllSpace {
		return true
	}

	l, r := 0, bytesLen-1
	for l < r {
		if bytes[l] != bytes[r] {
			return false
		}
		l++
		r--
	}

	return true
}

func test01() {
	s := "A man, a plan, a canal: Panama"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPalindrome(s))
	fmt.Println()
}

func test02() {
	s := "race a car"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPalindrome(s))
	fmt.Println()
}

func test03() {
	s := " "
	succResult := true
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPalindrome(s))
	fmt.Println()
}

func test04() {
	s := "0P119"
	succResult := false
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPalindrome(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
