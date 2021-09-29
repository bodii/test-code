package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) || s == t {
		return false
	}

	byteList := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		byteList[s[i]-'a']++
		byteList[t[i]-'a']--
	}

	for _, v := range byteList {
		if v != 0 {
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

func test04() {
	s, t := "a", "a"
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
	test04()
}
