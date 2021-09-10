package main

import (
	"fmt"
	"strings"
)

func maxScore(s string) int {
	zeros, ones := 0, 0
	max := 0

	ones = strings.Count(s, "1")

	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			zeros++
		} else {
			ones--
		}

		num := ones + zeros
		if max < num {
			max = num
		}
	}

	return max
}

func test01() {
	s := "011101"
	success := 5

	fmt.Println("test01  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", maxScore(s))
	fmt.Println()
}

func test02() {
	s := "00111"
	success := 5

	fmt.Println("test02  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", maxScore(s))
	fmt.Println()
}

func test03() {
	s := "1111"
	success := 3

	fmt.Println("test03  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", maxScore(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
