package main

import (
	"fmt"
	"strings"
)

func numOfStrings(patterns []string, word string) int {
	count := 0
	for _, v := range patterns {
		if strings.Contains(word, v) {
			count++
		}
	}

	return count
}

func test01() {
	patterns := []string{"a", "abc", "bc", "d"}
	word := "abc"
	succResult := 3
	fmt.Println("test01 patterns:=", patterns, " word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numOfStrings(patterns, word))
	fmt.Println()
}

func test02() {
	patterns := []string{"a", "b", "c"}
	word := "aaaaabbbbb"
	succResult := 2
	fmt.Println("test02 patterns:=", patterns, " word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numOfStrings(patterns, word))
	fmt.Println()
}

func test03() {
	patterns := []string{"a", "a", "a"}
	word := "ab"
	succResult := 3
	fmt.Println("test03 patterns:=", patterns, " word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numOfStrings(patterns, word))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
