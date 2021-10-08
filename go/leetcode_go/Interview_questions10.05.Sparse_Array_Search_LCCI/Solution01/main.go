package main

import (
	"fmt"
)

func findString(words []string, s string) int {
	for k, v := range words {
		if v == s {
			return k + 1
		}
	}

	return -1
}

func test01() {
	words := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	s := "ta"
	succResult := -1
	fmt.Println("test01 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findString(words, s))
	fmt.Println()
}

func test02() {
	words := []string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}
	s := "ball"
	succResult := 4
	fmt.Println("test02 words:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findString(words, s))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
