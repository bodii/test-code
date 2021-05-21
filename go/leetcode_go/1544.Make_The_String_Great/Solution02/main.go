package main

import (
	"fmt"
)

func makeGood(s string) string {
	sLen := len(s)
	resultBytes := make([]byte, sLen)

	for i := 0; i < sLen; i++ {
		resultBytesLen := len(resultBytes)
		if resultBytesLen > 0 && toUpper(resultBytes[resultBytesLen-1]) == toUpper(s[i]) && resultBytes[resultBytesLen-1] != s[i] {
			resultBytes = resultBytes[0 : resultBytesLen-1]
		} else {
			resultBytes = append(resultBytes, s[i])
		}
	}

	return string(resultBytes)
}

func toUpper(a byte) byte {
	if isLowerChar(a) {
		return a - 32
	}

	return a
}

func isLowerChar(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}

	return false
}

func test01() {
	s := "leEeetcode"
	succResult := "leetcode"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeGood(s))
	fmt.Println()
}

func test02() {
	s := "abBAcC"
	succResult := ""
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeGood(s))
	fmt.Println()
}

func test03() {
	s := "s"
	succResult := "s"
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeGood(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
