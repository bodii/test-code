package main

import (
	"fmt"
)

func countSegments(s string) int {
	count := 0
	for k, v := range s {
		if (k == 0 || s[k-1] == ' ') && v != ' ' {
			count++
		}
	}

	return count
}

func test01() {
	s := "Hello, my name is John"
	succResult := 5
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countSegments(s))
	fmt.Println()
}

func test02() {
	s := "Hello"
	succResult := 1
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countSegments(s))
	fmt.Println()
}

func test03() {
	s := "love live! mu'sic forever"
	succResult := 4
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countSegments(s))
	fmt.Println()
}

func test04() {
	s := ""
	succResult := 0
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countSegments(s))
	fmt.Println()
}

func test05() {
	s := ", , , ,        a, eaefa"
	succResult := 6
	fmt.Println("test05 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countSegments(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
