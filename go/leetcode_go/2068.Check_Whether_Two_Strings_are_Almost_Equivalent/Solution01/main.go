package main

import (
	"fmt"
)

func checkAlmostEquivalent(word1 string, word2 string) bool {
	wordSet := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		wordSet[word1[i]-'a']++
		wordSet[word2[i]-'a']--
	}

	for _, v := range wordSet {
		if v > 3 || v < -3 {
			return false
		}
	}

	return true
}

func test01() {
	word1, word2 := "aaaa", "bccb"
	succResult := false
	fmt.Println("test01 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkAlmostEquivalent(word1, word2))
	fmt.Println()
}

func test02() {
	word1, word2 := "abcdeef", "abaaacc"
	succResult := true
	fmt.Println("test02 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkAlmostEquivalent(word1, word2))
	fmt.Println()
}

func test03() {
	word1, word2 := "cccddabba", "babababab"
	succResult := true
	fmt.Println("test03 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkAlmostEquivalent(word1, word2))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
