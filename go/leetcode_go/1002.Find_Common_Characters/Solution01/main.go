package main

import (
	"fmt"
	"strings"
)

func commonChars(words []string) []string {
	result := []string{}
	resMap := map[rune]bool{}

	for _, v := range words[0] {
		vStr := string(v)
		if resMap[v] {
			continue
		}
		resMap[v] = true

		count := strings.Count(words[0], vStr)
		min := count
		for i := 1; i < len(words); i++ {
			otherCount := strings.Count(words[i], vStr)
			if min > otherCount {
				min = otherCount
			}

			if otherCount == 0 {
				break
			}

		}

		if min == 0 {
			continue
		}

		for i := 0; i < min; i++ {
			result = append(result, vStr)
		}
	}

	return result
}

func test01() {
	words := []string{"bella", "label", "roller"}
	succResult := []string{"e", "l", "l"}
	fmt.Println("test01 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", commonChars(words))
	fmt.Println()
}

func test02() {
	words := []string{"cool", "lock", "cook"}
	succResult := []string{"c", "o"}
	fmt.Println("test02 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", commonChars(words))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
