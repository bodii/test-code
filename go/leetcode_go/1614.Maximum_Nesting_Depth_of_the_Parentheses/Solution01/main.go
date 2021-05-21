package main

import (
	"fmt"
)

func maxDepth(s string) int {
	maxDepth, depth := 0, 0

	for _, v := range s {
		if v == '(' {
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
		} else if v == ')' {
			depth--
		}
	}

	return maxDepth
}

func test01() {
	s := "(1+(2*3)+((8)/4))+1"
	succResult := 3
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test02() {
	s := "(1)+((2))+(((3)))"
	succResult := 3
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test03() {
	s := "1+(2*3)/(2-1)"
	succResult := 1
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test04() {
	s := "1"
	succResult := 0
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func test05() {
	s := "8*((1*(5+6))*(8/6))"
	succResult := 3
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDepth(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
