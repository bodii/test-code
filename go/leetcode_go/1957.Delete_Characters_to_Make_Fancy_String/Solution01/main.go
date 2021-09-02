package main

import (
	"fmt"
)

func makeFancyString(s string) string {
	sLen := len(s)
	result := make([]byte, 0)

	for i := 0; i < sLen; i++ {
		result = append(result, s[i])

		next := i + 1
		for ; next < sLen; next++ {
			if s[i] != s[next] {
				break
			}
		}

		if next-i > 1 {
			result = append(result, s[i])
			i = next - 1
		}

	}

	return string(result)
}

func test01() {
	s := "leeetcode"
	succResult := "leetcode"
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeFancyString(s))
	fmt.Println()
}

func test02() {
	s := "aaabaaaa"
	succResult := "aabaa"
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeFancyString(s))
	fmt.Println()
}

func test03() {
	s := "aab"
	succResult := "aab"
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", makeFancyString(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
