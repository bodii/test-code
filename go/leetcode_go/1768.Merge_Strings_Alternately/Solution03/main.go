package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	word1Len := len(word1)
	word2Len := len(word2)

	maxLen := word1Len
	if word2Len > word1Len {
		maxLen = word2Len
	}

	mergeBytes := make([]byte, word1Len+word2Len)
	j := 0
	for i := 0; i < maxLen; i++ {
		if i < word1Len {
			mergeBytes[j] = word1[i]
			j++
		}

		if i < word2Len {
			mergeBytes[j] = word2[i]
			j++
		}
	}

	return string(mergeBytes)
}

func test01() {
	word1 := "abc"
	word2 := "pqr"
	succResult := "apbqcr"
	fmt.Println("test01 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", mergeAlternately(word1, word2))
	fmt.Println()
}

func test02() {
	word1 := "ab"
	word2 := "pqrs"
	succResult := "apbqrs"
	fmt.Println("test02 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", mergeAlternately(word1, word2))
	fmt.Println()
}

func test03() {
	word1 := "abcd"
	word2 := "pq"
	succResult := "apbqcd"
	fmt.Println("test03 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", mergeAlternately(word1, word2))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
