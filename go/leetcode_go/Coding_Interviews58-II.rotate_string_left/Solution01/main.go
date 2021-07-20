package main

import (
	"fmt"
)

func reverseLeftWords(s string, n int) string {
	result := []byte{}

	result = append(result, s[n:]...)
	result = append(result, s[:n]...)

	return string(result)
}

func test01() {
	s := "abcdefg"
	k := 2
	succResult := "cdefgab"
	fmt.Println("test01 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reverseLeftWords(s, k))
	fmt.Println()
}

func test02() {
	s := "lrloseumgh"
	k := 6
	succResult := "umghlrlose"
	fmt.Println("test01 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reverseLeftWords(s, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
