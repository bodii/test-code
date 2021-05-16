package main

import (
	"fmt"
)

func truncateSentence(s string, k int) string {
	spaces := 0
	for i := 0; i < len(s); i++ {
		// if s[i] == 32 {
		if s[i] == ' ' {
			spaces++
		}

		if spaces == k {
			return s[0:i]
		}
	}
	return s
}

func test01() {
	s := "Hello how are you Contestant"
	k := 4
	succResult := "Hello how are you"
	fmt.Println("test01 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", truncateSentence(s, k))
	fmt.Println()
}

func test02() {
	s := "What is the solution to this problem"
	k := 4
	succResult := "What is the solution"
	fmt.Println("test02 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", truncateSentence(s, k))
	fmt.Println()
}

func test03() {
	s := "chopper is not a tanuki"
	k := 5
	succResult := "chopper is not a tanuki"
	fmt.Println("test03 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", truncateSentence(s, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
