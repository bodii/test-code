package main

import (
	"fmt"
	"strings"
)

func findOcurrences(text string, first string, second string) []string {
	result := make([]string, 0)
	textArr := strings.Split(text, " ")
	for i := 0; i < len(textArr)-2; i++ {
		if textArr[i] == first && textArr[i+1] == second {
			result = append(result, textArr[i+2])
		}
	}
	return result
}

func test01() {
	text := "alice is a good girl she is a good student"
	first, second := "a", "good"
	succResult := []string{"girl", "student"}

	fmt.Printf("test01: [%v] is first=\"%s\",vsecond=\"%s\"\n",
		text, first, second)
	fmt.Println("success result:", succResult)
	fmt.Println(findOcurrences(text, first, second))
	fmt.Println()
}

func test02() {
	text := "we will we will rock you"
	first, second := "we", "will"
	succResult := []string{"we", "rock"}

	fmt.Printf("test02: [%v] is first=\"%s\", second=\"%s\"\n",
		text, first, second)
	fmt.Println("success result:", succResult)
	fmt.Println(findOcurrences(text, first, second))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
