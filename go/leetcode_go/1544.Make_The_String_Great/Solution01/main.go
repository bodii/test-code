package main

import (
	"fmt"
)

func makeGood(s string) string {
	tLen := len(s)
	sBytes := []byte(s)

	needMake := true
	for needMake && tLen > 1 {
		for i := 1; i < tLen; i++ {
			fmt.Println(string(sBytes))
			if (isLowerChar(sBytes[i-1]) && isUpperChar(sBytes[i]) && hasEqualChar(sBytes[i-1], sBytes[i])) ||
				(isUpperChar(sBytes[i-1]) && isLowerChar(sBytes[i]) && hasEqualChar(sBytes[i-1], sBytes[i])) {
				if i-1 == 0 {
					sBytes = sBytes[2:]
				} else {
					tmp := sBytes[i+1:]
					sBytes = sBytes[0 : i-1]
					sBytes = append(sBytes, tmp...)
				}
				tLen -= 2
				break
			}

			if i == tLen-1 {
				needMake = false
				break
			}
		}
	}

	return string(sBytes)
}

func isUpperChar(c byte) bool {
	if c >= 'A' && c <= 'Z' {
		return true
	}

	return false
}

func isLowerChar(c byte) bool {
	if c >= 'a' && c <= 'z' {
		return true
	}

	return false
}

func hasEqualChar(a, b byte) bool {
	if a-b == 32 || b-a == 32 {
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
