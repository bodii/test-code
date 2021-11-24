package main

import (
	"fmt"
	"strconv"
)

func areNumbersAscending(s string) bool {
	prev, start := -1, 0
	for i := 1; i < len(s)-1; i++ {
		if s[i] == ' ' {
			if s[start]-'0' <= 9 {
				n, _ := strconv.Atoi(s[start:i])
				if n <= prev {
					return false
				}
				prev = n
			}
			start = i + 1
		}
	}

	if s[start]-'0' <= 9 {
		n, _ := strconv.Atoi(s[start:])
		if n <= prev {
			return false
		}
	}

	return true
}

func test01() {
	s := "1 box has 3 blue 4 red 6 green and 12 yellow marbles"
	succResult := true
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", areNumbersAscending(s))
	fmt.Println()
}

func test02() {
	s := "hello world 5 x 5"
	succResult := false
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", areNumbersAscending(s))
	fmt.Println()
}

func test03() {
	s := "sunset is at 7 51 pm overnight lows will be in the low 50 and 60 s"
	succResult := false
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", areNumbersAscending(s))
	fmt.Println()
}

func test04() {
	s := "4 5 11 26"
	succResult := true
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", areNumbersAscending(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
