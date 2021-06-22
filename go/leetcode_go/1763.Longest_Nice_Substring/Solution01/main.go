package main

import (
	"fmt"
)

func longestNiceSubstring(s string) string {
	sLen := len(s)
	result := ""
	for i := 0; i < sLen-1; i++ {
		for j := i + 1; j < sLen; j++ {
			if isNice(s, i, j) && len(result) < j-i+1 {
				result = s[i : j+1]
			}
		}
	}
	return result
}

func isNice(s string, start, end int) bool {
	lowers := make(map[byte]bool)
	uppers := make(map[byte]bool)

	for i := start; i <= end; i++ {
		if s[i] >= 97 {
			lowers[s[i]] = true
		} else {
			uppers[s[i]] = true
		}
	}

	for i := start; i <= end; i++ {
		if (s[i] >= 97 && !uppers[s[i]-32]) || s[i] < 97 && !lowers[s[i]+32] {
			return false
		}
	}
	return true
}

func test01() {
	s := "YazaAay"
	succResult := "aAa"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test02() {
	s := "Bb"
	succResult := "Bb"
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test03() {
	s := "c"
	succResult := ""
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test04() {
	s := "dDzeE"
	succResult := "dD"
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test05() {
	s := "cChH"
	succResult := "cChH"
	fmt.Println("test05 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test06() {
	s := "dcChHfFKkv"
	succResult := "cChHfFKk"
	fmt.Println("test06 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test07() {
	s := "dcChHfFKkvaABbEeJjoOLl"
	succResult := "aABbEeJjoOLl"
	fmt.Println("test07 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test08() {
	s := "BebjJE"
	succResult := "BebjJE"
	fmt.Println("test08 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
	fmt.Println()
}

func test09() {
	s := "qQUjJ"
	succResult := "qQ"
	fmt.Println("test09 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", longestNiceSubstring(s))
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
