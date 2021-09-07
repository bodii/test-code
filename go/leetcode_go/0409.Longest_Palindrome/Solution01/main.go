package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	sLen := len(s)

	if sLen == 1 {
		return 1
	}

	count := 0
	sMap := make(map[byte]int)
	for i := 0; i < sLen; i++ {
		sMap[s[i]]++
	}

	for _, v := range sMap {
		if v >= 2 {
			count += (v / 2) * 2
		}
	}

	if count < sLen {
		count++
	}

	return count
}

func test01() {
	s := "abccccdd"
	succResult := 7
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestPalindrome(s))
	fmt.Println()
}

func test02() {
	s := "a"
	succResult := 1
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestPalindrome(s))
	fmt.Println()
}

func test03() {
	s := "bb"
	succResult := 2
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestPalindrome(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
