package main

import (
	"fmt"
)

func firstUniqChar(s string) byte {
	sList := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		sList[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if sList[s[i]-'a'] == 1 {
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
