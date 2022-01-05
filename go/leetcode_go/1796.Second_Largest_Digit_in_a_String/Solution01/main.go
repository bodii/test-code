package main

import (
	"fmt"
)

func secondHighest(s string) int {
	first, second := -1, -1
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			first, second = findHighests(first, second, int(s[i]-'0'))
		}
	}

	return second
}

func findHighests(a, b, c int) (int, int) {
	if a == b && b != c {
		if a > c {
			return a, c
		}
		return c, a
	} else if a == c && b != c {
		if a > b {
			return a, b
		}
		return b, a
	} else if a > b && a > c {
		if b > c {
			return a, b
		}
		return a, c
	} else if b > a && b > c {
		if a > c {
			return b, a
		}
		return b, c
	}

	if a > b {
		return c, a
	}

	return c, b
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
