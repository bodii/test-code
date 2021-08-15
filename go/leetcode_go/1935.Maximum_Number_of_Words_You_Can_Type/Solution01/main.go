package main

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	words := strings.Split(text, " ")

	count := 0
	for _, w := range words {
		isContain := false
		for _, b := range brokenLetters {
			if strings.ContainsRune(w, b) {
				isContain = true
				break
			}
		}

		if !isContain {
			count++
		}
	}

	return count
}

func test01() {
	text, brokenLetters := "hello world", "ad"
	succResult := 1
	fmt.Println("test01 text:=", text, " brokenLetters:=", brokenLetters)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeTypedWords(text, brokenLetters))
	fmt.Println()
}

func test02() {
	text, brokenLetters := "leet code", "lt"
	succResult := 1
	fmt.Println("test02 text:=", text, " brokenLetters:=", brokenLetters)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeTypedWords(text, brokenLetters))
	fmt.Println()
}

func test03() {
	text, brokenLetters := "leet code", "e"
	succResult := 0
	fmt.Println("test03 text:=", text, " brokenLetters:=", brokenLetters)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeTypedWords(text, brokenLetters))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
