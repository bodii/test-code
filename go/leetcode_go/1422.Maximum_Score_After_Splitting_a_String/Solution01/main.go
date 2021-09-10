package main

import (
	"fmt"
)

func maxScore(s string) int {
	l, r := 1, len(s)

	max := 0
	for ; l < r; l++ {
		sum := findSubstrNum(s, 0, r-l, '0') + findSubstrNum(s, r-l, r, '1')
		if max < sum {
			max = sum
		}
	}

	return max
}

func findSubstrNum(s string, l, r int, sub byte) int {
	count := 0

	for ; l < r; l++ {
		if s[l] == sub {
			count++
		}
	}

	return count
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
