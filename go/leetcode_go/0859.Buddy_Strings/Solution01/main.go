package main

import (
	"fmt"
)

func buddyStrings(s string, goal string) bool {
	s1Len, s2Len := len(s), len(goal)
	if s1Len != s2Len {
		return false
	}

	for i := 0; i < s1Len-1; i++ {
		for j := i + 1; j < s1Len; j++ {
			if swap(s, i, j) == goal {
				return true
			}
		}
	}

	return false
}

func swap(s string, i, j int) string {
	sBytes := []byte(s)

	sBytes[i], sBytes[j] = sBytes[j], sBytes[i]

	return string(sBytes)
}

func test01() {
	s, t := "ab", "ba"
	succResult := true
	fmt.Println("test01 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func test02() {
	s, t := "ab", "ab"
	succResult := false
	fmt.Println("test02 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func test03() {
	s, t := "aa", "aa"
	succResult := true
	fmt.Println("test03 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func test04() {
	s, t := "aaaaaaabc", "aaaaaaacb"
	succResult := true
	fmt.Println("test04 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func test05() {
	s, t := "", "aa"
	succResult := false
	fmt.Println("test05 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
