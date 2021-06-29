package main

import (
	"fmt"
)

func removeOuterParentheses(s string) string {
	sum := 0
	parentheses := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			sum++
			if sum > 1 {
				parentheses = append(parentheses, s[i])
			}
		} else {
			sum--
			if sum > 0 {
				parentheses = append(parentheses, s[i])
			}
		}
	}

	return string(parentheses)
}

func test01() {
	s := "(()())(())"
	succResult := "()()()"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", removeOuterParentheses(s))
	fmt.Println()
}

func test02() {
	s := "(()())(())(()(()))"
	succResult := "()()()()(())"
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", removeOuterParentheses(s))
	fmt.Println()
}

func test03() {
	s := "()()"
	succResult := ""
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", removeOuterParentheses(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
