package main

import (
	"fmt"
)

func minTimeToType(word string) int {
	result := 0

	prev := 'a'
	for _, v := range word {
		cur := abs(int(v - prev))
		result += min(cur, 26-cur) + 1

		prev = v
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
