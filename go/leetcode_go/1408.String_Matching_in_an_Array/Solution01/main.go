package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	wordsLen := len(words)
	wordsMap := make(map[string]bool)
	for i := 0; i < wordsLen-1; i++ {
		for j := i + 1; j < wordsLen; j++ {
			if strings.Contains(words[i], words[j]) {
				wordsMap[words[j]] = true
			} else if strings.Contains(words[j], words[i]) {
				wordsMap[words[i]] = true
			}
		}
	}

	result := make([]string, 0)
	for k := range wordsMap {
		result = append(result, k)
	}

	return result
}

func test01() {
	words := []string{"mass", "as", "hero", "superhero"}
	succResult := []string{"as", "hero"}
	fmt.Println("test01 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", stringMatching(words))
	fmt.Println()
}

func test02() {
	words := []string{"leetcode", "et", "code"}
	succResult := []string{"et", "code"}
	fmt.Println("test02 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", stringMatching(words))
	fmt.Println()
}

func test03() {
	words := []string{"blue", "green", "bu"}
	succResult := []string{}
	fmt.Println("test03 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", stringMatching(words))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
