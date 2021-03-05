package main

import (
	"fmt"
	"strings"
)

func sortString(s string) string {
	var result strings.Builder
	mapBox := make(map[int]int, 26)

	for _, i := range s {
		mapBox[int(i-'a')]++
	}

	len := len(s)
	for i := 0; i < len; i++ {
		for j := 0; j < 26; j++ {
			if mapBox[j] > 0 {
				result.WriteByte(byte(j + 'a'))
				mapBox[j]--
			}
		}

		for j := 25; j >= 0; j-- {
			if mapBox[j] > 0 {
				result.WriteByte(byte(j + 'a'))
				mapBox[j]--
			}
		}
	}

	return result.String()
}

func test01() {
	s := "aaaabbbbcccc"
	// fmt.Println([]rune(s))
	fmt.Println("result", sortString(s))
}

func test02() {
	s := "rat"
	fmt.Println("result", sortString(s))
}

func test03() {
	s := "leetcode"
	fmt.Println("result", sortString(s))
}

func test04() {
	s := "ggggggg"
	fmt.Println("result", sortString(s))
}

func test05() {
	s := "spo"
	fmt.Println("result", sortString(s))
}

func main() {
	// test01()

	// test02()

	// test03()

	// test04()

	test05()
}
