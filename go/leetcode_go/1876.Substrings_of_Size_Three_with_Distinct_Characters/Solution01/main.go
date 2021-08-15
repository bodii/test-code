package main

import (
	"fmt"
)

func countGoodSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s)-2; i++ {
		if s[i] != s[i+1] && s[i] != s[i+2] && s[i+1] != s[i+2] {
			count++
		}
	}

	return count
}

func test01() {
	s := "xyzzaz"
	succResult := 1
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countGoodSubstrings(s))
	fmt.Println()
}

func test02() {
	s := "aababcabc"
	succResult := 4
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countGoodSubstrings(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
