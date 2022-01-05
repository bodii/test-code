package main

import (
	"fmt"
	"unicode"
)

func secondHighest(s string) int {
	first, second := -1, -1
	for _, v := range s {
		if unicode.IsDigit(v) {
			val := int(v - '0')
			if first == -1 {
				first = val
			} else if val > first {
				second, first = first, val
			} else if val < first && val > second {
				second = val
			}
		}
	}

	return second
}

func test01() {
	s := "dfa12321afd"
	succResult := 2
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", secondHighest(s))
	fmt.Println()
}

func test02() {
	s := "abc1111"
	succResult := -1
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", secondHighest(s))
	fmt.Println()
}

func test03() {
	s := "ck077"
	succResult := 0
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", secondHighest(s))
	fmt.Println()
}

func test04() {
	s := "1"
	succResult := -1
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", secondHighest(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
