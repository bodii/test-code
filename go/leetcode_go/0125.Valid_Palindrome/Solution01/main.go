package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	len := len(s)
	newS := strings.Builder{}

	for i := 0; i < len; i++ {
		if (s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57) {
			newS.WriteByte(s[i])
		}
	}

	newStrLen := newS.Len()
	newStr := newS.String()
	for i := 0; i < newStrLen; i++ {
		if newStr[i] != newStr[newStrLen-1-i] {
			return false
		}
	}

	return true
}

func test01() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println("test01 result: ", isPalindrome(s))
}

func test02() {
	s := "race a car"
	fmt.Println("test02 result: ", isPalindrome(s))
}

func test03() {
	s := " "
	fmt.Println("test03 result: ", isPalindrome(s))
}

func test04() {
	s := "0P119"
	fmt.Println("test04 result: ", isPalindrome(s))
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
