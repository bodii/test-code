package main

import (
	"fmt"
)

func firstUniqChar(s string) byte {
	sMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		sMap[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if sMap[s[i]] == 1 {
			return s[i]
		}
	}

	return ' '
}

func test01() {
	s := "abaccdeff"
	succResult := 'b'
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", firstUniqChar(s))
	fmt.Println()
}

func test02() {
	s := ""
	succResult := ' '
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", firstUniqChar(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
