package main

import (
	"fmt"
)

func CheckPermutation(s1 string, s2 string) bool {
	sMap := make(map[rune]int)
	for _, v := range s1 {
		sMap[v]++
	}

	for _, v := range s2 {
		sMap[v]--
		if sMap[v] == 0 {
			delete(sMap, v)
		}
	}

	return len(sMap) == 0
}

func test01() {
	s1, s2 := "abc", "bca"
	successResult := true
	fmt.Println("test01 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", CheckPermutation(s1, s2))
	fmt.Println()
}

func test02() {
	s1, s2 := "abc", "bad"
	successResult := false
	fmt.Println("test02 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", CheckPermutation(s1, s2))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
