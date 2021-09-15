package main

import (
	"fmt"
)

func makeEqual(words []string) bool {
	wordMap := make(map[rune]int)
	wordsLen := len(words)

	for _, word := range words {
		for _, s := range word {
			wordMap[s]++
		}
	}

	for _, w := range wordMap {
		if w != wordsLen {
			return false
		}
	}

	return true
}

func test01() {
	words := []string{"abc", "aabc", "bc"}
	succResult := true
	fmt.Println("test01 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeEqual(words))
	fmt.Println()
}

func test02() {
	words := []string{"ab", "a"}
	succResult := false
	fmt.Println("test02 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeEqual(words))
	fmt.Println()
}

func test03() {
	words := []string{"aabbccdde", "e"}
	succResult := true
	fmt.Println("test03 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeEqual(words))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
