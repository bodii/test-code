package main

import (
	"fmt"
)

func countBinarySubstrings(s string) int {
	sLen, counts := len(s), 0
	prevNum, curNum := 0, 1

	for i := 1; i < sLen; i++ {
		if s[i] == s[i-1] {
			curNum++
		} else {
			prevNum, curNum = curNum, 1
		}

		if prevNum >= curNum {
			counts++
		}
	}

	return counts
}

func test01() {
	s := "00110011"
	succResult := 6
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBinarySubstrings(s))
	fmt.Println()
}

func test02() {
	s := "10101"
	succResult := 4
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBinarySubstrings(s))
	fmt.Println()
}

func test03() {
	s := "111100011000"
	succResult := 7
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBinarySubstrings(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
