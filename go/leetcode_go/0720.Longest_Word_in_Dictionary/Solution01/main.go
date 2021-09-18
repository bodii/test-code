package main

import (
	"fmt"
	"sort"
)

func longestWord(words []string) string {
	sort.Strings(words)
	wordsMap := make(map[string]bool)
	result := ""
	for _, w := range words {
		wLen, resultLen := len(w), len(result)
		if wLen == 1 || wordsMap[w[0:wLen-1]] {
			wordsMap[w] = true
			if resultLen == 0 || resultLen < wLen {
				result = w
			}
		}
	}

	return result
}

func test01() {
	words := []string{"w", "wo", "wor", "worl", "world"}
	succResult := "world"
	fmt.Println("test01 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestWord(words))
	fmt.Println()
}

func test02() {
	words := []string{"a", "banana", "app", "appl", "ap", "apply", "apple"}
	succResult := "apple"
	fmt.Println("test02 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestWord(words))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
