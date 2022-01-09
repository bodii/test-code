package main

import (
	"fmt"
)

func calculateTime(keyboard string, word string) int {
	keyInds := make([]int, 26)
	for i := 0; i < len(keyboard); i++ {
		keyInds[keyboard[i]-'a'] = i
	}

	times := keyInds[word[0]-'a']
	for i := 1; i < len(word); i++ {
		times += abs(keyInds[word[i]-'a'] - keyInds[word[i-1]-'a'])
	}

	return times
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func test01() {
	keyboard := "abcdefghijklmnopqrstuvwxyz"
	word := "cba"
	succResult := 4
	fmt.Println("test01 keyboard:=", keyboard, "word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", calculateTime(keyboard, word))
	fmt.Println()
}

func test02() {
	keyboard := "pqrstuvwxyzabcdefghijklmno"
	word := "leetcode"
	succResult := 73
	fmt.Println("test01 keyboard:=", keyboard, "word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", calculateTime(keyboard, word))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
