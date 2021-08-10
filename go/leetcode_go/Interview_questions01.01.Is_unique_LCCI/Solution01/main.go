package main

import (
	"fmt"
)

func isUnique(astr string) bool {
	astrLen := len(astr)

	for i := 0; i < astrLen-1; i++ {
		for j := i + 1; j < astrLen; j++ {
			if astr[i] == astr[j] {
				return false
			}
		}
	}

	return true
}

func test01() {
	s := "leetcode"
	succResult := false
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUnique(s))
	fmt.Println()
}

func test02() {
	s := "abc"
	succResult := true
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUnique(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
