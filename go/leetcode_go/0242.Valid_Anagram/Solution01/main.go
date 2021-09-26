package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	bMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		bMap[s[i]]++
		bMap[t[i]]--
	}

	for k := range bMap {
		if bMap[k] != 0 {
			return false
		}
	}

	return true
}

func test01() {
	s, t := "anagram", "nagaram"
	succResult := true
	fmt.Println("test01 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isAnagram(s, t))
	fmt.Println()
}

func test02() {
	s, t := "rat", "car"
	succResult := false
	fmt.Println("test02 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isAnagram(s, t))
	fmt.Println()
}

func test03() {
	s, t := "a", "ab"
	succResult := false
	fmt.Println("test03 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isAnagram(s, t))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
