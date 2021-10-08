package main

import (
	"fmt"
)

func findString(words []string, s string) int {
	left, right := 0, len(words)-1

	for left <= right {
		if words[left] == s {
			return left
		}

		if words[right] == s {
			return right
		}
		left++
		right--
	}

	return -1
}

func test01() {
	words := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	s := "ta"
	succResult := -1
	fmt.Println("test01 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findString(words, s))
	fmt.Println()
}

func test02() {
	words := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	s := "ball"
	succResult := 4
	fmt.Println("test02 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findString(words, s))
	fmt.Println()
}

func test03() {
	words := []string{"at", "ball", "car", "", "dad"}
	s := "car"
	succResult := 2
	fmt.Println("test03 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findString(words, s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
