package main

import (
	"fmt"
)

func maxNumberOfBalloons(text string) int {
	b := map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0}
	for _, v := range text {
		if _, ok := b[v]; ok {
			b[v]++
		}
	}

	min := 0
	if b['l'] < b['o'] {
		min = b['l'] / 2
	} else {
		min = b['o'] / 2
	}

	if min < 0 {
		return 0
	}

	if min < b['b'] && min < b['a'] {
		return min
	} else if b['b'] < min && b['a'] >= min {
		return b['b']
	} else if b['a'] < min && b['b'] >= min {
		return b['a']
	} else if b['a'] < min && b['b'] < min {
		if b['a'] < b['b'] {
			return b['a']
		} else {
			return b['b']
		}
	}

	return min
}

func test01() {
	text := "nlaebolko"
	success := 1

	fmt.Println("test01  text:=", text)
	fmt.Println("success result:", success)
	fmt.Println("result:", maxNumberOfBalloons(text))
	fmt.Println()
}

func test02() {
	text := "loonbalxballpoon"
	success := 2

	fmt.Println("test02  text:=", text)
	fmt.Println("success result:", success)
	fmt.Println("result:", maxNumberOfBalloons(text))
	fmt.Println()
}

func test03() {
	text := "leetcode"
	success := 0

	fmt.Println("test03  text:=", text)
	fmt.Println("success result:", success)
	fmt.Println("result:", maxNumberOfBalloons(text))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
