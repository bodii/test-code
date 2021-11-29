package main

import (
	"fmt"
	"unicode"
)

func numDifferentIntegers(word string) int {
	nums := make(map[string]bool)

	l, r := 0, len(word)
	for l < r {
		if !unicode.IsDigit(rune(word[l])) {
			l++
			continue
		}

		j := l + 1
		for j < r && unicode.IsDigit(rune(word[j])) {
			if word[l] == '0' {
				l++
			}
			j++
		}

		nums[word[l:j]] = true
		l = j
	}

	return len(nums)
}

func test01() {
	word := "a123bc34d8ef34"
	succResult := 3
	fmt.Println("test01 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numDifferentIntegers(word))
	fmt.Println()
}

func test02() {
	word := "leet1234code234"
	succResult := 2
	fmt.Println("test02 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numDifferentIntegers(word))
	fmt.Println()
}

func test03() {
	word := "a1b01c001"
	succResult := 1
	fmt.Println("test03 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numDifferentIntegers(word))
	fmt.Println()
}

func test04() {
	word := "a1b01c001a3"
	succResult := 2
	fmt.Println("test04 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numDifferentIntegers(word))
	fmt.Println()
}

func test05() {
	word := "2393706880236110407059624696967828762752651982730115221690437821508229419410771541532394006597463715513741725852432559057224478815116557380260390432211227579663571046845842281704281749571110076974264971989893607137140456254346955633455446057823738757323149856858154529105301197388177242583658641529908583934918768953462557716z97438020429952944646288084173334701047574188936201324845149110176716130267041674438237608038734431519439828191344238609567530399189316846359766256507371240530620697102864238792350289978450509162697068948604722646739174590530336510475061521094503850598453536706982695212493902968251702853203929616930291257062173c79487281900662343830648295410"
	succResult := 3
	fmt.Println("test05 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numDifferentIntegers(word))
	fmt.Println()
}

func test06() {
	word := "0a0"
	succResult := 1
	fmt.Println("test06 word:=", word)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numDifferentIntegers(word))
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
