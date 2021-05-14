package main

import (
	"fmt"
)

func halvesAreAlike(s string) bool {
	vowelNums := 0
	mid := len(s) / 2
	vowels := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'A': 1, 'E': 1,
		'I': 1, 'O': 1, 'U': 1}
	fmt.Println('a', 'e', 'i', 'o', 'u')
	for i := 0; i < mid; i++ {
		if _, ok := vowels[s[i]]; ok {
			vowelNums++
		}

		if _, ok := vowels[s[mid+i]]; ok {
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
