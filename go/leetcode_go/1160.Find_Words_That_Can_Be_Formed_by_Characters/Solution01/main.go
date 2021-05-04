package main

import "fmt"

func countCharacters(words []string, chars string) int {
	nums := 0
	charMap := countWorld(chars)
	for _, v := range words {
		wordMap := countWorld(v)
		if containWord(wordMap, charMap) {
			nums += len(v)
		}
	}
	return nums
}

func countWorld(s string) map[rune]int {
	rest := make(map[rune]int)
	for _, v := range s {
		rest[v]++
	}
	return rest
}

func containWord(a, b map[rune]int) bool {
	for k, v := range a {
		v2, ok := b[k]
		fmt.Println(v, v2, k)
		if !ok || v2 < v {
			return false
		}
	}
	return true
}

func test01() {
	words := []string{"cat", "bt", "hat", "tree"}
	chars := "atach"
	fmt.Printf("test01 words:=%v, chars=%s, result=[6]:%d\n", words, chars, countCharacters(words, chars))
}

func test02() {
	words := []string{"hello", "world", "leetcode"}
	chars := "welldonehoneyr"
	fmt.Printf("test01 words:=%v, chars=%s, result=[10]:%d\n", words, chars, countCharacters(words, chars))
}

func main() {
	test01()
	test02()
}
