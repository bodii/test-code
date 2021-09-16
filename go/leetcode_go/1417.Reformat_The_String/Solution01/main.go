package main

import (
	"fmt"
)

func reformat(s string) string {
	numbers := make([]rune, 0)
	strings := make([]rune, 0)

	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			strings = append(strings, v)
		} else {
			numbers = append(numbers, v)
		}
	}

	stringsLen, numbersLen := len(strings), len(numbers)
	if stringsLen < numbersLen-1 || numbersLen < stringsLen-1 {
		return ""
	}

	result := make([]rune, 0)
	maxLen, minLen := stringsLen, numbersLen
	maxs, mins := strings, numbers
	if numbersLen > maxLen {
		maxLen, minLen = minLen, maxLen
		maxs, mins = mins, maxs
	}

	for k, v := range maxs {
		if k < maxLen {
			result = append(result, v)
		}
		if k < minLen {
			result = append(result, mins[k])
		}
	}

	return string(result)
}

func test01() {
	s := "a0b1c2"
	success := "0a1b2c"

	fmt.Println("test01  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", reformat(s))
	fmt.Println()
}

func test02() {
	s := "leetcode"
	success := ""

	fmt.Println("test02  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", reformat(s))
	fmt.Println()
}

func test03() {
	s := "1229857369"
	success := ""

	fmt.Println("test03  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", reformat(s))
	fmt.Println()
}

func test04() {
	s := "covid2019"
	success := "c2o0v1i9d"

	fmt.Println("test04  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", reformat(s))
	fmt.Println()
}

func test05() {
	s := "ab123"
	success := "1a2b3"

	fmt.Println("test05  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", reformat(s))
	fmt.Println()
}

func test06() {
	s := "a0b1c2"
	success := "a0b1c2"

	fmt.Println("test06  s:=", s)
	fmt.Println("success result:", success)
	fmt.Println("result:", reformat(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
