package main

import (
	"fmt"
)

func detectCapitalUse(word string) bool {
	isUpFirstStr := isUpperStr(word[0])

	isUp := false
	for i := 1; i < len(word); i++ {
		if i == 1 {
			isUp = isUpperStr(word[i])
			if !isUpFirstStr && isUp {
				return false
			}
		} else {
			if isUp != isUpperStr(word[i]) {
				return false
			}
		}
	}

	return true
}

func isUpperStr(s byte) bool {
	return s >= 'A' && s <= 'Z'
}

func test01() {
	word := "USA"
	succResult := true
	fmt.Println("test01 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", detectCapitalUse(word))
	fmt.Println()
}

func test02() {
	word := "FlaG"
	succResult := false
	fmt.Println("test02 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", detectCapitalUse(word))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
