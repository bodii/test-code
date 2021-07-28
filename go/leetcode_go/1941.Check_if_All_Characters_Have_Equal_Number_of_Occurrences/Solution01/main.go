package main

import (
	"fmt"
)

func areOccurrencesEqual(s string) bool {
	sMap := make(map[rune]int)
	for _, v := range s {
		sMap[v]++
	}

	current := 0
	for _, v := range sMap {
		if current == 0 {
			current = v
		} else if current != v {
			return false
		}
	}

	return true
}

func test01() {
	s := "abacbc"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", areOccurrencesEqual(s))
	fmt.Println()
}

func test02() {
	s := "aaabb"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", areOccurrencesEqual(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
