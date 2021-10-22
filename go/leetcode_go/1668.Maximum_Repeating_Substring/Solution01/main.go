package main

import (
	"fmt"
)

func maxRepeating(sequence string, word string) int {
	num, result := 0, 0
	wLen := len(word)

	step, i := 0, 0
	for i <= len(sequence)-wLen {
		if sequence[i:i+wLen] == word {
			num++
			i += wLen
			result = max(result, num)
		} else {
			step++
			i, num = step, 0
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	sequence := "ababc"
	word := "ab"
	succResult := 2
	fmt.Println("test01 sequence:=", sequence, "word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxRepeating(sequence, word))
	fmt.Println()
}

func test02() {
	sequence := "ababc"
	word := "ba"
	succResult := 1
	fmt.Println("test01 sequence:=", sequence, "word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxRepeating(sequence, word))
	fmt.Println()
}

func test03() {
	sequence := "ababc"
	word := "ac"
	succResult := 0
	fmt.Println("test01 sequence:=", sequence, "word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxRepeating(sequence, word))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
