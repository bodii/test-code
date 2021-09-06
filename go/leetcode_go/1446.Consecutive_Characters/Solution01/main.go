package main

import (
	"fmt"
)

func maxPower(s string) int {
	len := len(s)
	max := 0

	for i := 0; i < len; i++ {
		j := i + 1
		for ; j < len; j++ {
			if s[i] != s[j] {
				break
			}
		}

		if j-i > max {
			max = j - i
		}

		if j == len {
			return max
		}

		i = j - 1
	}

	return max
}

func test01() {
	s := "leetcode"
	succResult := 2
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxPower(s))
	fmt.Println()
}

func test02() {
	s := "abbcccddddeeeeedcba"
	succResult := 5
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxPower(s))
	fmt.Println()
}

func test03() {
	s := "triplepillooooow"
	succResult := 5
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxPower(s))
	fmt.Println()
}

func test04() {
	s := "hooraaaaaaaaaaay"
	succResult := 11
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxPower(s))
	fmt.Println()
}

func test05() {
	s := "tourist"
	succResult := 1
	fmt.Println("test05 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxPower(s))
	fmt.Println()
}

func test06() {
	s := "hooraaaaaaaaaaaa"
	succResult := 12
	fmt.Println("test06 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxPower(s))
	fmt.Println()
}

func test07() {
	s := "j"
	succResult := 1
	fmt.Println("test07 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxPower(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
}
