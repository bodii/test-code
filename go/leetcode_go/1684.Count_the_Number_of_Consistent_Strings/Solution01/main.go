package main

import "fmt"

func countConsistentStrings(allowed string, words []string) int {
	nums := 0
	for _, v := range words {
		flag := true
		for _, c := range v {
			if !contains(allowed, c) {
				flag = false
				break
			}
		}

		if flag {
			nums++
		}
	}

	return nums
}

func contains(allowed string, s rune) bool {
	for _, v := range allowed {
		if v == s {
			return true
		}
	}

	return false
}

func test01() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}
	succResult := 2
	fmt.Println("test01 allowedv:=", allowed, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countConsistentStrings(allowed, words))
	fmt.Println()
}

func test02() {
	allowed := "abc"
	words := []string{"a", "b", "c", "ab", "ac", "bc", "abc"}
	succResult := 7
	fmt.Println("test02 allowedv:=", allowed, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countConsistentStrings(allowed, words))
	fmt.Println()
}

func test03() {
	allowed := "cad"
	words := []string{"cc", "acd", "b", "ba", "bac", "bad", "ac", "d"}
	succResult := 4
	fmt.Println("test03 allowedv:=", allowed, " words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countConsistentStrings(allowed, words))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
