package main

import (
	"fmt"
)

func minOperations(s string) int {
	c1, c2 := 0, 0
	for i := 0; i < len(s); i++ {
		if (s[i] == '0' && 1&i == 1) || (s[i] == '1' && 1&i == 0) {
			c1++
		} else {
			c2++
		}
	}

	if c1 < c2 {
		return c1
	}

	return c2
}

func test01() {
	s := "0100"
	succResult := 1
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(s))
	fmt.Println()
}

func test02() {
	s := "10"
	succResult := 0
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(s))
	fmt.Println()
}

func test03() {
	s := "1111"
	succResult := 2
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(s))
	fmt.Println()
}

func test04() {
	s := "10010100"
	succResult := 3
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(s))
	fmt.Println()
}

func test05() {
	s := "1100010111"
	succResult := 3
	fmt.Println("test05 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
