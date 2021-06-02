package main

import (
	"fmt"
	"math"
)

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	return stringToInt(firstWord)+stringToInt(secondWord) == stringToInt(targetWord)
}

func stringToInt(s string) int {
	result := 0
	sLen := len(s)
	j := sLen - 1
	for i := 0; i < sLen; i++ {
		num := int(uint8(s[i] - 'a'))
		if j > 0 {
			num *= int(math.Pow10(j))
		}
		j--
		result += num
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
