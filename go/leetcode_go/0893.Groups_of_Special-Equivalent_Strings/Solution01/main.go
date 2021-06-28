package main

import (
	"fmt"
	"sort"
)

func numSpecialEquivGroups(words []string) int {
	groups := make(map[string]int)

	for _, word := range words {
		equivs := [2][]rune{}
		for k, c := range word {
			equivs[k%2] = append(equivs[k%2], c)
		}

		sort.Slice(equivs[0], func(i, j int) bool {
			return equivs[0][i] < equivs[0][j]
		})

		sort.Slice(equivs[1], func(i, j int) bool {
			return equivs[1][i] < equivs[1][j]
		})

		groups[string(append(equivs[0], equivs[1]...))] = 1
	}

	return len(groups)
}

func test01() {
	words := []string{"abcd", "cdab", "cbad", "xyzz", "zzxy", "zzyx"}
	succResult := 3
	fmt.Println("test01 prices:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numSpecialEquivGroups(words))
	fmt.Println()
}

func test02() {
	words := []string{"abc", "acb", "bac", "bca", "cab", "cba"}
	succResult := 3
	fmt.Println("test01 prices:=", words)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numSpecialEquivGroups(words))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
