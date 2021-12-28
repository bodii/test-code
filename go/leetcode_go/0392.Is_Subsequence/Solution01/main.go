package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	index, sLen := 0, len(s)

	for i := 0; i < len(t) && index != sLen; i++ {
		if s[index] == t[i] {
			index++
		}
	}

	return index == sLen
}

func test01() {
	s, t := "abc", "ahbgdc"
	succResult := true
	fmt.Println("test01 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isSubsequence(s, t))
	fmt.Println()
}

func test02() {
	s, t := "axc", "ahbgdc"
	succResult := false
	fmt.Println("test02 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isSubsequence(s, t))
	fmt.Println()
}

func test03() {
	s, t := "abc", "ahbgdcdffd"
	succResult := true
	fmt.Println("test03 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isSubsequence(s, t))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
