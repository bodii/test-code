package main

import (
	"fmt"
	"strings"
)

func halvesAreAlike(s string) bool {
	vowels := "aeiouAEIOU"
	vowelNums := 0
	mid := len(s) / 2
	for i := 0; i < mid; i++ {
		if strings.Contains(vowels, string(s[i])) {
			vowelNums++
		}

		if strings.Contains(vowels, string(s[mid+i])) {
			vowelNums--
		}
	}

	return vowelNums == 0
}

func test01() {
	s := "book"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", halvesAreAlike(s))
	fmt.Println()
}

func test02() {
	s := "textbook"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", halvesAreAlike(s))
	fmt.Println()
}

func test03() {
	s := "MerryChristmas"
	succResult := false
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", halvesAreAlike(s))
	fmt.Println()
}

func test04() {
	s := "AbCdEfGh"
	succResult := true
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", halvesAreAlike(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
