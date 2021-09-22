package main

import (
	"fmt"
	"strings"
)

func reversePrefix(word string, ch byte) string {
	index := strings.IndexByte(word, ch)
	if index == -1 {
		return word
	}

	wordBytes := []byte(word)
	wordBytes = append(reverse(wordBytes[0:index+1]), wordBytes[index+1:]...)
	return string(wordBytes)
}

func reverse(sBytes []byte) []byte {
	sLen := len(sBytes)
	for i := 0; i < sLen/2; i++ {
		sBytes[i], sBytes[sLen-1-i] = sBytes[sLen-1-i], sBytes[i]
	}

	return sBytes
}

func test01() {
	word := "abcdefd"
	ch := byte('d')
	successResult := "dcbaefd"

	fmt.Println("func test01 word:=", word, " ch:=", ch)
	fmt.Println("success result:=", successResult)
	fmt.Println("result: ", reversePrefix(word, ch))
	fmt.Println()
}

func test02() {
	word := "xyxzxe"
	ch := byte('z')
	successResult := "zxyxxe"

	fmt.Println("func test02 word:=", word, " ch:=", ch)
	fmt.Println("success result:=", successResult)
	fmt.Println("result: ", reversePrefix(word, ch))
	fmt.Println()
}

func test03() {
	word := "abcd"
	ch := byte('z')
	successResult := "abcd"

	fmt.Println("func test03 word:=", word, " ch:=", ch)
	fmt.Println("success result:=", successResult)
	fmt.Println("result: ", reversePrefix(word, ch))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
