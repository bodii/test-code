package main

import (
	"fmt"
	"strings"
)

func firstUniqChar(s string) byte {
	for i := 0; i < len(s); i++ {
		if strings.LastIndexByte(s, s[i]) == i && strings.IndexByte(s, s[i]) == i {
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

func test03() {
	s := "cc"
	succResult := ' '
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", firstUniqChar(s))
	fmt.Println()
}

func test04() {
	s := "z"
	succResult := 'z'
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", firstUniqChar(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
