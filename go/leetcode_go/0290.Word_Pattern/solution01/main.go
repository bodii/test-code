package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	patternLen := len(pattern)
	words := strings.Split(s, " ")
	if 0 == patternLen || patternLen != len(words) {
		return false
	}

	if 1 == patternLen {
		return true
	}

	patternBox := make(map[byte]int)
	wordBox := make(map[string]int)
	for i := 0; i < patternLen; i++ {
		if patternBox[pattern[i]] < 1 {
			patternBox[pattern[i]] = i + 1
		}

		if wordBox[words[i]] < 1 {
			wordBox[words[i]] = i + 1
		}

		// fmt.Println(patternBox[pattern[i]], wordBox[words[i]])
		if patternBox[pattern[i]] != wordBox[words[i]] {
			return false
		}
	}

	return true
}

func test01() {
	pattern := "abba"
	s := "dog cat cat dog"
	fmt.Printf("result: %t\n", wordPattern(pattern, s))
}

func test02() {
	pattern := "abba"
	s := "dog cat cat fish"
	fmt.Printf("result: %t\n", wordPattern(pattern, s))
}

func test03() {
	pattern := "aaaa"
	s := "dog cat cat dog"
	fmt.Printf("result: %t\n", wordPattern(pattern, s))
}

func test04() {
	pattern := "abba"
	s := "dog dog dog dog"
	fmt.Printf("result: %t\n", wordPattern(pattern, s))
}

func test05() {
	pattern := "e"
	s := "eukera"

	fmt.Printf("result: %t\n", wordPattern(pattern, s))
}

func test06() {
	pattern := "aba"
	s := "dog cat cat"

	fmt.Printf("result: %t\n", wordPattern(pattern, s))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
