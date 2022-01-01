package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	maxWordNum := 0

	for _, s := range sentences {
		maxWordNum = max(maxWordNum, strings.Count(s, " "))
	}

	return maxWordNum + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	succResult := 6
	fmt.Println("test01 sentences:=", sentences)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", mostWordsFound(sentences))
	fmt.Println()
}

func test02() {
	sentences := []string{"please wait", "continue to fight", "continue to win"}
	succResult := 3
	fmt.Println("test02 sentences:=", sentences)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", mostWordsFound(sentences))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
