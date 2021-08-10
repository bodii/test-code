package main

import (
	"fmt"
)

func isUnique(astr string) bool {
	strMap := make(map[rune]int)

	for _, s := range astr {
		strMap[s]++
		if strMap[s] > 1 {
			return false
		}
	}

	return true
}

func test01() {
	s := "leetcode"
	succResult := false
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUnique(s))
	fmt.Println()
}

func test02() {
	s := "abc"
	succResult := true
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isUnique(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
