package main

import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	return gcd(str1, str2)
}

func gcd(str1, str2 string) string {
	if str1 == str2 {
		return str1
	}

	str1Len, str2Len := len(str1), len(str2)
	minLen, max, min := str2Len, str1, str2
	if str1Len < minLen {
		minLen, max, min = str1Len, str2, str1
	}

	if max[:minLen] != min {
		return ""
	}

	return gcd(min, max[minLen:])
}

func test01() {
	str1 := "ABCABC"
	str2 := "ABC"
	succResult := "ABC"
	fmt.Println("test01 str1:=", str1, " str2:=", str2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", gcdOfStrings(str1, str2))
	fmt.Println()
}

func test02() {
	str1 := "ABABAB"
	str2 := "ABAB"
	succResult := "AB"
	fmt.Println("test02 str1:=", str1, " str2:=", str2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", gcdOfStrings(str1, str2))
	fmt.Println()
}

func test03() {
	str1 := "LEET"
	str2 := "CODE"
	succResult := ""
	fmt.Println("test03 str1:=", str1, " str2:=", str2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", gcdOfStrings(str1, str2))
	fmt.Println()
}

func test04() {
	str1 := "ABCDEF"
	str2 := "ABC"
	succResult := ""
	fmt.Println("test04 str1:=", str1, " str2:=", str2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", gcdOfStrings(str1, str2))
	fmt.Println()
}

func test05() {
	str1 := "ABABABAB"
	str2 := "ABABAB"
	succResult := "AB"
	fmt.Println("test05 str1:=", str1, " str2:=", str2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", gcdOfStrings(str1, str2))
	fmt.Println()
}

func test06() {
	str1 := "ABCCBAABCCBAABCCBA"
	str2 := "ABCCBAABCCBA"
	succResult := "ABCCBA"
	fmt.Println("test06 str1:=", str1, " str2:=", str2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", gcdOfStrings(str1, str2))
	fmt.Println()
}

func test07() {
	str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"
	succResult := "TAUXX"
	fmt.Println("test07 str1:=", str1, " str2:=", str2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", gcdOfStrings(str1, str2))
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
