package main

import (
	"fmt"
	"sort"
)

func checkAlmostEquivalent(word1 string, word2 string) bool {
	wordSet := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		wordSet[word1[i]-'a']++
		wordSet[word2[i]-'a']--
	}

	sort.Ints(wordSet)

	return wordSet[0] >= -3 && wordSet[25] <= 3
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
