package main

import (
	"fmt"
)

func removeDuplicates(s string) string {
	stack := []byte{}
	for i := range s {
		if len(stack) > 0 && stack[len(stack)-1] == s[i] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}

func test01() {
	s := "abbaca"
	success := "ca"

	fmt.Println("test01  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", removeDuplicates(s))
	fmt.Println()
}

func test02() {
	s := "azxxzy"
	success := "ay"

	fmt.Println("test02  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", removeDuplicates(s))
	fmt.Println()
}

func test03() {
	s := "aaaaaaaa"
	success := ""

	fmt.Println("test03  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", removeDuplicates(s))
	fmt.Println()
}

func test04() {
	s := "aababaab"
	success := "ba"

	fmt.Println("test04  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", removeDuplicates(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
