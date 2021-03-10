package main

import (
	"fmt"
	"strings"
)

func sortString(s string) string {
	var result strings.Builder
	mapBox := make(map[int]int, 26)
	strCount := 0

	for _, c := range s {
		mapBox[int(c-'a')]++
		strCount++
	}

	i, flag := 0, 1
	for strCount > 0 {
		if mapBox[i] > 0 {
			result.WriteByte(byte(i) + 'a')
			strCount--
			mapBox[i]--
		}

		i += flag
		if i == 26 || i == -1 {
			flag *= -1
			i += flag
		}
	}

	return result.String()
}

func test01() {
	s := "aaaabbbbcccc"
	fmt.Println("result:", sortString(s))
}

func test02() {
	s := "rat"
	fmt.Println("result:", sortString(s))
}

func test03() {
	s := "leetcode"
	fmt.Println("result:", sortString(s))
}

func test04() {
	s := "ggggggg"
	fmt.Println("result:", sortString(s))
}

func test05() {
	s := "spo"
	fmt.Println("result:", sortString(s))
}

func main() {
	// test01()

	// test02()

	// test03()

	// test04()

	test05()
}
