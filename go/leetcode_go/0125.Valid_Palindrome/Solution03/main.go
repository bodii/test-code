package main

import (
	"fmt"
)

func isPalindrome(s string) bool {
	len := len(s)
	left, right := 0, len-1

	for i := 0; i < len; i++ {
		if isValidChar(s[left]) && isValidChar(s[right]) {
			if !isEqual(s[left], s[right]) {
				return false
			}
			left++
			right--
		} else if isValidChar(s[left]) && !isValidChar(s[right]) {
			right--
		} else if !isValidChar(s[left]) && isValidChar(s[right]) {
			left++
		} else if !isValidChar(s[left]) && !isValidChar(s[right]) {
			left++
			right--
			i++
		}
	}
	return true
}

func isEqual(left, right byte) bool {
	if isUpper(left) {
		left += 32
	}

	if isUpper(right) {
		right += 32
	}
	return left == right
}

func isUpper(c byte) bool {
	if 65 <= c && c <= 90 {
		return true
	}

	return false
}

func isValidChar(c byte) bool {
	if isUpper(c) || (97 <= c && c <= 122) || (48 <= c && c <= 57) {
		return true
	}

	return false
}

func test01() {
	s := "Z man, a plan, a canal: Panama"
	fmt.Println("test01 is ture result: ", isPalindrome(s))
}

func test02() {
	s := "race a car"
	fmt.Println("test02 is false result: ", isPalindrome(s))
}

func test03() {
	s := " "
	fmt.Println("test03 is true result: ", isPalindrome(s))
}

func test04() {
	s := "0P119"
	fmt.Println("test04 is false result: ", isPalindrome(s))
}

func test05() {
	s := ".,"
	fmt.Println("test05 is true result: ", isPalindrome(s))
}

func test06() {
	s := "ab"
	fmt.Println("test06 is false result: ", isPalindrome(s))
}

func test07() {
	s := "......a....."
	fmt.Println("test07 is true result: ", isPalindrome(s))
}

func test08() {
	s := "a.b,."
	fmt.Println("test08 is false result: ", isPalindrome(s))
}

func test09() {
	s := ",,,,,,,,,,,,acva"
	fmt.Println("test09 is false result: ", isPalindrome(s))
}

func test10() {
	s := "a"
	fmt.Println("test10 is true result: ", isPalindrome(s))
}

func test11() {
	s := "a."
	fmt.Println("test11 is true result: ", isPalindrome(s))
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
	test10()
	test11()
}
