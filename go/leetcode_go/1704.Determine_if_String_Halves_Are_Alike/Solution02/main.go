package main

import (
	"fmt"
)

func halvesAreAlike(s string) bool {
	vowelNums := 0
	mid := len(s) / 2

	isVowel := func(b byte) bool {
		result := false
		switch b {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			result = true
		}

		return result
	}

	for i := 0; i < mid; i++ {
		if isVowel(s[i]) {
			vowelNums++
		}

		if isVowel(s[mid+i]) {
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
