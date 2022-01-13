package main

import (
	"fmt"
	"strings"
)

func capitalizeTitle(title string) string {
	words := strings.Split(title, " ")

	for k, v := range words {
		words[k] = strings.ToLower(v)
		if len(v) > 2 {
			words[k] = strings.Title(words[k])
		}
	}

	return strings.Join(words, " ")
}

func test01() {
	s := "capiTalIze tHe titLe"
	successResult := "Capitalize The Title"
	fmt.Println("test01 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func test02() {
	s := "First leTTeR of EACH Word"
	successResult := "First Letter of Each Word"
	fmt.Println("test02 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func test03() {
	s := "i lOve leetcode"
	successResult := "i Love Leetcode"
	fmt.Println("test03 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", capitalizeTitle(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
