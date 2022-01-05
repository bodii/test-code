package main

import (
	"fmt"
)

func validPalindrome(s string) bool {
	for l, r := 0, len(s)-1; l < r; l++ {
		if s[l] != s[r] {
			return isPalindrome(s, l+1, r) || isPalindrome(s, l, r-1)
		}
		r--
	}
	return true
}

func isPalindrome(s string, l, r int) bool {
	for ; l < r; l++ {
		if s[l] != s[r] {
			return false
		}
		r--
	}

	return true
}

func test01() {
	s := "aba"
	succResult := true
	fmt.Println("test01:")
	fmt.Println("s:=", s)
	fmt.Println("success result: ", succResult)
	fmt.Println("result: ", validPalindrome(s))
	fmt.Println()
}

func test02() {
	s := "abca"
	succResult := true
	fmt.Println("test02:")
	fmt.Println("s:=", s)
	fmt.Println("success result: ", succResult)
	fmt.Println("result: ", validPalindrome(s))
	fmt.Println()
}

func test03() {
	s := "abc"
	succResult := false
	fmt.Println("test03:")
	fmt.Println("s:=", s)
	fmt.Println("success result: ", succResult)
	fmt.Println("result: ", validPalindrome(s))
	fmt.Println()
}

func test04() {
	s := "deeee"
	succResult := true
	fmt.Println("test04:")
	fmt.Println("s:=", s)
	fmt.Println("success result: ", succResult)
	fmt.Println("result: ", validPalindrome(s))
	fmt.Println()
}

func test05() {
	s := "cbbcc"
	succResult := true
	fmt.Println("test05:")
	fmt.Println("s:=", s)
	fmt.Println("success result: ", succResult)
	fmt.Println("result: ", validPalindrome(s))
	fmt.Println()
}

func test06() {
	s := "aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga"
	succResult := true
	fmt.Println("test06:")
	fmt.Println("s:=", s)
	fmt.Println("success result: ", succResult)
	fmt.Println("result: ", validPalindrome(s))
	fmt.Println()
}

func test07() {
	s := "acxcybycxcxa"
	succResult := true
	fmt.Println("test07:")
	fmt.Println("s:=", s)
	fmt.Println("success result: ", succResult)
	fmt.Println("result: ", validPalindrome(s))
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
}
