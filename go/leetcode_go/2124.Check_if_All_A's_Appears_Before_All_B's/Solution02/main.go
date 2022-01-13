package main

import (
	"fmt"
)

func checkString(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] == 'b' && s[i] == 'a' {
			return false
		}
	}

	return true
}

func test01() {
	s := "aaabbb"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkString(s))
	fmt.Println()
}

func test02() {
	s := "abab"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkString(s))
	fmt.Println()
}

func test03() {
	s := "bbb"
	succResult := true
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkString(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
