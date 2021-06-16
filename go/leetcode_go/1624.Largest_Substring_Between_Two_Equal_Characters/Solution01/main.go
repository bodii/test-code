package main

import (
	"fmt"
)

func maxLengthBetweenEqualCharacters(s string) int {
	max := -1

	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] && j-i-1 > max {
				max = j - i - 1
			}
		}
	}

	return max
}

func test01() {
	s := "aa"
	succResult := 0
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxLengthBetweenEqualCharacters(s))
	fmt.Println()
}

func test02() {
	s := "abca"
	succResult := 2
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxLengthBetweenEqualCharacters(s))
	fmt.Println()
}

func test03() {
	s := "cbzxy"
	succResult := -1
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxLengthBetweenEqualCharacters(s))
	fmt.Println()
}

func test04() {
	s := "cabbac"
	succResult := 4
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxLengthBetweenEqualCharacters(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
