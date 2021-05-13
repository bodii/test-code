package main

import (
	"fmt"
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return merge(word1) == merge(word2)
}

func merge(strs []string) string {
	s := strings.Builder{}

	for _, v := range strs {
		s.WriteString(v)
	}

	return s.String()
}

func test01() {
	word1, word2 := []string{"ab", "c"}, []string{"a", "bc"}
	succResult := true
	fmt.Println("test01 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arrayStringsAreEqual(word1, word2))
	fmt.Println()
}

func test02() {
	word1, word2 := []string{"a", "cb"}, []string{"ab", "c"}
	succResult := false
	fmt.Println("test02 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arrayStringsAreEqual(word1, word2))
	fmt.Println()
}

func test03() {
	word1, word2 := []string{"abc", "d", "defg"}, []string{"abcddefg"}
	succResult := true
	fmt.Println("test03 word1:=", word1, " word2:=", word2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arrayStringsAreEqual(word1, word2))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
