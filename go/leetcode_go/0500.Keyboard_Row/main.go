package main

import (
	"fmt"
	"strings"
)

var keyLists = []string{
	"qwertyuiop",
	"asdfghjkl",
	"zxcvbnm",
}

func findWords(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		index := -1
		for j := 0; j < len(words[i]); j++ {
			index = find(index, string(words[i][j]))
			if -1 == index {
				break
			}
		}
		if -1 < index {
			result = append(result, words[i])
		}
	}

	return result
}

func find(index int, b string) int {
	b = strings.ToLower(b)
	result := -1
	if -1 == index {
		for i := 0; i < len(keyLists); i++ {
			if true == strings.Contains(keyLists[i], b) {
				return i
			}
		}
	} else {
		if true == strings.Contains(keyLists[index], b) {
			return index
		}
	}
	return result
}

func test1() {
	words := []string{"Hello", "Alaska", "Dad", "Peace"}
	v := findWords(words)
	fmt.Println(v)
}

func test2() {
	words := []string{"omk"}
	v := findWords(words)
	fmt.Println(v)
}

func test3() {
	words := []string{"adsdf", "sfd"}
	v := findWords(words)
	fmt.Println(v)
}

func main() {
	// test1()
	// test2()
	test3()
}
