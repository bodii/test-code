package main

import (
	"fmt"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return str2Int(firstWord)+str2Int(secondWord) == str2Int(targetWord)
}

func str2Int(s string) int {
	result := 0
	for _, v := range s {
		result = result*10 + int(v-'a')
	}
	return result
}

func test01() {
	firstWord, secondWord, targetWord := "acb", "cba", "cdb"
	succResult := true
	fmt.Println("test01 firstWord:=", firstWord, " secondWord:=", secondWord, " targetWord:=", targetWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isSumEqual(firstWord, secondWord, targetWord))
	fmt.Println()
}

func test02() {
	firstWord, secondWord, targetWord := "aaa", "a", "aab"
	succResult := false
	fmt.Println("test02 firstWord:=", firstWord, " secondWord:=", secondWord, " targetWord:=", targetWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isSumEqual(firstWord, secondWord, targetWord))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
