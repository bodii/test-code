package main

import (
	"fmt"
)

func buddyStrings(s string, goal string) bool {
	s1Len, s2Len := len(s), len(goal)
	if s1Len != s2Len || s1Len < 2 {
		return false
	}

	if s1Len == 2 {
		return s[0] == goal[1] && s[1] == goal[0]
	}

	d1, d2 := make([]byte, 0), make([]byte, 0)
	d1Map := make(map[byte]int)
	for i := 0; i < s1Len; i++ {
		d1Map[s[i]]++
		if s[i] != goal[i] {
			d1 = append(d1, s[i])
			d2 = append(d2, goal[i])
		}
	}

	if len(d1) == 0 {
		for _, v := range d1Map {
			if v >= 2 {
				return true
			}
		}
		return false
	} else if len(d1) != 2 {
		return false
	}

	return d1[0] == d2[1] && d1[1] == d2[0]
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

func test06() {
	s, t := "abab", "abab"
	succResult := true
	fmt.Println("test06 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func test07() {
	s, t := "abcd", "abcd"
	succResult := false
	fmt.Println("test07 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func test08() {
	s, t := "aaab", "aaab"
	succResult := true
	fmt.Println("test08 s:=", s, " t:=", t)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buddyStrings(s, t))
	fmt.Println()
}

func test09() {
	s, t := "abcabc", "abcabc"
	succResult := true
	fmt.Println("test09 s:=", s, " t:=", t)
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
	test06()
	test07()
	test08()
	test09()
}
