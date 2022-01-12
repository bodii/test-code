package main

import (
	"fmt"
	"strings"
)

func checkOnesSegment(s string) bool {
	return strings.Contains(s, "01")
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
