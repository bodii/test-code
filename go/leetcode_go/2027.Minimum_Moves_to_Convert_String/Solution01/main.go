package main

import (
	"fmt"
)

func minimumMoves(s string) int {
	sLen := len(s)

	count := 0
	for i := 0; i < sLen; {
		if s[i] == 'O' {
			i++
			continue
		}

		count++
		i += 3
	}

	return count
}

func test01() {
	s := "XXX"
	succResult := 1
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumMoves(s))
	fmt.Println()
}

func test02() {
	s := "XXOX"
	succResult := 2
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumMoves(s))
	fmt.Println()
}

func test03() {
	s := "OOOO"
	succResult := 0
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumMoves(s))
	fmt.Println()
}

func test04() {
	s := "OXOX"
	succResult := 1
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumMoves(s))
	fmt.Println()
}

func test05() {
	s := "OOXOOOX"
	succResult := 2
	fmt.Println("test05 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumMoves(s))
	fmt.Println()
}

func test06() {
	s := "XXXOOXXX"
	succResult := 2
	fmt.Println("test06 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumMoves(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
