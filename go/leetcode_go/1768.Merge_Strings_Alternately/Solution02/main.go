package main

import (
	"fmt"
)

func mergeAlternately(word1 string, word2 string) string {
	word1Len := len(word1)
	word2Len := len(word2)

	minLen := word1Len
	minLenIsWord1 := true
	if word2Len < word1Len {
		minLen = word2Len
		minLenIsWord1 = false
	}

	mergeBytes := make([]byte, minLen*2)
	j := 0
	for i := 0; i < minLen; i++ {
		mergeBytes[j] = word1[i]
		mergeBytes[j+1] = word2[i]
		j += 2
	}

	if word1Len != word2Len {
		if minLenIsWord1 {
			mergeBytes = append(mergeBytes, word2[minLen:]...)
		} else {
			mergeBytes = append(mergeBytes, word1[minLen:]...)
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
