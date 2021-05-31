package main

import (
	"fmt"
)

func reverseWords(s string) string {
	revStr := make([]byte, 0)
	left := 0

	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			word := s[left:i]
			revStr = append(revStr, reverseWord(word)...)
			revStr = append(revStr, ' ')
			left = i + 1
		}
	}

	word := s[left:]
	revStr = append(revStr, reverseWord(word)...)

	return string(revStr)
}

func reverseWord(word string) []byte {
	wordLen := len(word)
	revWord := make([]byte, wordLen)
	i := 0
	for j := len(word) - 1; j >= 0; j-- {
		revWord[i] = word[j]
		i++
	}

	return revWord
}

func test01() {
	s := "Let's take LeetCode contest"
	successResult := "s'teL ekat edoCteeL tsetnoc"
	fmt.Println("test01 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", reverseWords(s))
	fmt.Println()
}

func test02() {
	s := "God Ding"
	successResult := "doG gniD"
	fmt.Println("test02 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", reverseWords(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
