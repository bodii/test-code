package main

import (
	"fmt"
	"strings"
)

func repeatedStringMatch(a string, b string) int {
	aLen, bLen := len(a), len(b)
	maxLen := bLen + aLen*2
	repA := strings.Builder{}
	repA.WriteString(a)

	nums := 1
	for repA.Len() <= maxLen {
		if strings.Contains(repA.String(), b) {
			return nums
		}
		repA.WriteString(a)
		nums++
	}

	return -1
}

func test01() {
	a := "abcd"
	b := "cdabcdab"
	succResult := 3
	fmt.Println("test01 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedStringMatch(a, b))
	fmt.Println()
}

func test02() {
	a := "a"
	b := "aa"
	succResult := 2
	fmt.Println("test02 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedStringMatch(a, b))
	fmt.Println()
}

func test03() {
	a := "a"
	b := "a"
	succResult := 1
	fmt.Println("test03 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedStringMatch(a, b))
	fmt.Println()
}

func test04() {
	a := "abc"
	b := "wxyz"
	succResult := -1
	fmt.Println("test04 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedStringMatch(a, b))
	fmt.Println()
}

func test05() {
	a := "abc"
	b := "cabcabca"
	succResult := 4
	fmt.Println("test05 a:=", a, " b:=", b)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", repeatedStringMatch(a, b))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
