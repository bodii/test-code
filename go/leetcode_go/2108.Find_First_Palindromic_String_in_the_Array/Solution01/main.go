package main

import (
	"fmt"
)

func firstPalindrome(words []string) string {
	for _, word := range words {
		if isPalindrome(word) {
			return word
		}
	}

	return ""
}

func isPalindrome(word string) bool {
	len := len(word)

	for i := 0; i < len/2; i++ {
		if word[i] != word[len-1-i] {
			return false
		}
	}

	return true
}

func test01() {
	words := []string{"abc", "car", "ada", "racecar", "cool"}
	succResult := "ada"
	fmt.Println("test01 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", firstPalindrome(words))
	fmt.Println()
}

func test02() {
	words := []string{"notapalindrome", "racecar"}
	succResult := "racecar"
	fmt.Println("test02 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", firstPalindrome(words))
	fmt.Println()
}

func test03() {
	words := []string{"def", "ghi"}
	succResult := ""
	fmt.Println("test03 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", firstPalindrome(words))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
