package main

import (
	"fmt"
)

func checkOnesSegment(s string) bool {
	sLen := len(s)
	if sLen <= 2 {
		return true
	}

	for i := 1; i < sLen-1; i++ {
		if s[i] == '0' && s[i+1] == '1' {
			return false
		}
	}

	return true
}

func test01() {
	s := "1001"
	succResult := false
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkOnesSegment(s))
	fmt.Println()
}

func test02() {
	s := "110"
	succResult := true
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkOnesSegment(s))
	fmt.Println()
}

func test03() {
	s := "10000"
	succResult := true
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkOnesSegment(s))
	fmt.Println()
}

func test04() {
	s := "1100111"
	succResult := false
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkOnesSegment(s))
	fmt.Println()
}

func test05() {
	s := "1100111"
	succResult := false
	fmt.Println("test05 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkOnesSegment(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
