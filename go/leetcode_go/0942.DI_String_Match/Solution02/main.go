package main

import (
	"fmt"
)

func diStringMatch(s string) []int {
	sLen := len(s)

	left, right := 0, sLen
	result := make([]int, sLen+1)
	for k, v := range s {
		if v == 'I' {
			result[k] = left
			left++
		} else {
			result[k] = right
			right--
		}
	}
	result[sLen] = left

	return result
}

func test01() {
	s := "IDID"
	success := []int{0, 4, 1, 3, 2}

	fmt.Println("test01  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", diStringMatch(s))
	fmt.Println()
}

func test02() {
	s := "III"
	success := []int{0, 1, 2, 3}

	fmt.Println("test02  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", diStringMatch(s))
	fmt.Println()
}

func test03() {
	s := "DDI"
	success := []int{3, 2, 0, 1}

	fmt.Println("test02  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", diStringMatch(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
