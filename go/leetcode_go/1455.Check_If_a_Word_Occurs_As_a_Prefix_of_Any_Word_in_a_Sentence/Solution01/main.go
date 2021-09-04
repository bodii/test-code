package main

import (
	"fmt"
	"strings"
)

func isPrefixOfWord(sentence string, searchWord string) int {
	sentences := strings.Split(sentence, " ")
	for k, v := range sentences {
		if strings.HasPrefix(v, searchWord) {
			return k + 1
		}
	}

	return -1
}

func test01() {
	sentence, searchWord := "i love eating burger", "burg"
	succResult := 4
	fmt.Println("test01 sentence:=", sentence)
	fmt.Println("test01 searchWord:=", searchWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixOfWord(sentence, searchWord))
	fmt.Println()
}

func test02() {
	sentence, searchWord := "this problem is an easy problem", "pro"
	succResult := 2
	fmt.Println("test02 sentence:=", sentence)
	fmt.Println("test02 searchWord:=", searchWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixOfWord(sentence, searchWord))
	fmt.Println()
}

func test03() {
	sentence, searchWord := "i am tired", "you"
	succResult := -1
	fmt.Println("test03 sentence:=", sentence)
	fmt.Println("test03 searchWord:=", searchWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixOfWord(sentence, searchWord))
	fmt.Println()
}

func test04() {
	sentence, searchWord := "i use triple pillow", "pill"
	succResult := 4
	fmt.Println("test04 sentence:=", sentence)
	fmt.Println("test04 searchWord:=", searchWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixOfWord(sentence, searchWord))
	fmt.Println()
}

func test05() {
	sentence, searchWord := "hello from the other side", "they"
	succResult := -1
	fmt.Println("test05 sentence:=", sentence)
	fmt.Println("test05 searchWord:=", searchWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixOfWord(sentence, searchWord))
	fmt.Println()
}

func test06() {
	sentence, searchWord := "hellohello hellohellohello", "ell"
	succResult := -1
	fmt.Println("test06 sentence:=", sentence)
	fmt.Println("test06 searchWord:=", searchWord)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isPrefixOfWord(sentence, searchWord))
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
