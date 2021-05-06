package main

import (
	"fmt"
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	aStrs := strings.Split(A, " ")
	bStrs := strings.Split(B, " ")

	aLen := len(aStrs)
	bLen := len(bStrs)

	len := aLen
	if bLen > aLen {
		len = bLen
	}

	strMap := make(map[string]int)
	for i := 0; i < len; i++ {
		if i < aLen {
			strMap[aStrs[i]]++
		}

		if i < bLen {
			strMap[bStrs[i]]++
		}
	}

	rest := make([]string, 0)
	for k, v := range strMap {
		if v == 1 {
			rest = append(rest, k)
		}
	}

	return rest
}

func test01() {
	a := "this apple is sweet"
	b := "this apple is sour"

	fmt.Printf(
		"test01 a:%s b:%s is [\"sweet\",\"sour\"] result: %v\n",
		a, b, uncommonFromSentences(a, b))
}

func test02() {
	a := "apple apple"
	b := "banana"

	fmt.Printf(
		"test02 a:%s b:%s is [\"banana\"] result: %v\n",
		a, b, uncommonFromSentences(a, b))
}

func main() {
	test01()
	test02()
}
