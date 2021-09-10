package main

import (
	"fmt"
)

func isFlipedString(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	s1Len, s2Len := len(s1), len(s2)

	if s1Len != s2Len {
		return false
	}

	s2Bytes := []byte(s2)
	for i := 0; i < s1Len-1; i++ {
		if string(append(s2Bytes[i+1:s2Len], s2Bytes[0:i+1]...)) == s1 {
			return true
		}
	}

	return false
}

func test01() {
	s1, s2 := "waterbottle", "erbottlewat"
	succResult := true
	fmt.Println("test01 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isFlipedString(s1, s2))
	fmt.Println()
}

func test02() {
	s1, s2 := "aa", "aba"
	succResult := false
	fmt.Println("test02 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isFlipedString(s1, s2))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
