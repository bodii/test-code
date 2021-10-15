package main

import (
	"fmt"
)

func minTimeToType(word string) int {
	result := 0

	prev := 0
	for i := 0; i < len(word); i++ {
		cur := int(word[i] - 'a')
		result += min(abs(cur-prev), 26-abs(cur-prev)) + 1

		prev = cur
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func test01() {
	word := "abc"
	succResult := 5
	fmt.Println("test01 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minTimeToType(word))
	fmt.Println()
}

func test02() {
	word := "bza"
	succResult := 7
	fmt.Println("test02 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minTimeToType(word))
	fmt.Println()
}

func test03() {
	word := "zjpc"
	succResult := 34
	fmt.Println("test03 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minTimeToType(word))
	fmt.Println()
}

func test04() {
	word := "pdy"
	succResult := 31
	fmt.Println("test04 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minTimeToType(word))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
