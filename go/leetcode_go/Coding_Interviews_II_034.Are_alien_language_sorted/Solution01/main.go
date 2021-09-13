package main

import (
	"fmt"
)

func isAlienSorted(words []string, order string) bool {
	orderMap := make(map[byte]int)

	for k := 0; k < len(order); k++ {
		orderMap[order[k]] = k + 1
	}

	for i := 0; i < len(words)-1; i++ {
		minLen, l2 := len(words[i]), len(words[i+1])

		if minLen > l2 {
			minLen = l2
		}

		isEqual := true
		for j := 0; j < minLen; j++ {
			if orderMap[words[i][j]] > orderMap[words[i+1][j]] {
				return false
			} else if orderMap[words[i][j]] < orderMap[words[i+1][j]] {
				isEqual = false
				if j == 0 {
					break
				}
			}
		}

		if isEqual && len(words[i]) > minLen {
			return false
		}
	}

	return true
}

func test01() {
	words := []string{"hello", "leetcode"}
	order := "hlabcdefgijkmnopqrstuvwxyz"
	succResult := true
	fmt.Println("test01 words:=", words, " order:=", order)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isAlienSorted(words, order))
	fmt.Println()
}

func test02() {
	words := []string{"word", "world", "row"}
	order := "worldabcefghijkmnpqstuvxyz"
	succResult := false
	fmt.Println("test02 words:=", words, " order:=", order)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isAlienSorted(words, order))
	fmt.Println()
}

func test03() {
	words := []string{"apple", "app"}
	order := "abcdefghijklmnopqrstuvwxyz"
	succResult := false
	fmt.Println("test03 words:=", words, " order:=", order)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isAlienSorted(words, order))
	fmt.Println()
}

func test04() {
	words := []string{"kuvp", "q"}
	order := "ngxlkthsjuoqcpavbfdermiywz"
	succResult := true
	fmt.Println("test04 words:=", words, " order:=", order)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isAlienSorted(words, order))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
