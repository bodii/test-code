package main

import (
	"fmt"
	"sort"
	"strings"
)

func stringMatching(words []string) []string {
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	result := make([]string, 0)
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if strings.Contains(words[j], words[i]) {
				result = append(result, words[i])
				break
			}
		}
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
