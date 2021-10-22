package main

import (
	"fmt"
	"strings"
)

func maxRepeating(sequence string, word string) int {
	num := 0

	for repWord := word; strings.Contains(sequence, repWord); repWord += word {
		num++
	}

	return num
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
