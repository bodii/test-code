package main

import (
	"fmt"
	"sort"
	"strings"
)

func mostCommonWord(paragraph string, banned []string) string {
	stringNums := make(map[string]int)
	bannedNums := make(map[string]int)
	char := strings.Builder{}

	for k, v := range paragraph {
		if isEnglishChar(v) {
			if isUpperChar(v) {
				v = toLowerChar(v)
			}
			char.WriteRune(v)
		} else if char.Len() > 0 {
			stringNums[char.String()]++
			char = strings.Builder{}
		}

		if k == len(paragraph)-1 && char.Len() > 0 {
			stringNums[char.String()]++
		}
	}

	if len(stringNums) == 0 {
		return ""
	}

	for _, v := range banned {
		bannedNums[v]++
	}

	stringNumValus := make(map[int]string)
	stringNumSlice := make([]int, 0)
	for k, v := range stringNums {
		if _, ok := bannedNums[k]; ok {
			delete(stringNums, k)
			continue
		}

		stringNumValus[v] = k
		stringNumSlice = append(stringNumSlice, v)
	}

	if len(stringNumValus) == 0 {
		return ""
	}

	sort.Ints(stringNumSlice)
	return stringNumValus[stringNumSlice[len(stringNumSlice)-1]]
}

func isEnglishChar(e rune) bool {
	if (97 <= e && e <= 122) || (65 <= e && e <= 90) {
		return true
	}

	return false
}

func isUpperChar(e rune) bool {
	if 65 <= e && e <= 90 {
		return true
	}

	return false
}

func toLowerChar(e rune) rune {
	return e + 32
}

func test01() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	success := "ball"

	fmt.Println("test01 func paragraph:=", paragraph, " banned:=", banned)
	fmt.Println("success result:", success)
	fmt.Println("result:", mostCommonWord(paragraph, banned))
	fmt.Println()
}

func test02() {
	paragraph := "a."
	banned := []string{""}
	success := "a"

	fmt.Println("test02 func paragraph:=", paragraph, " banned:=", banned)
	fmt.Println("success result:", success)
	fmt.Println("result:", mostCommonWord(paragraph, banned))
	fmt.Println()
}

func test03() {
	paragraph := "Bob"
	banned := []string{""}
	success := "bob"

	fmt.Println("test03 func paragraph:=", paragraph, " banned:=", banned)
	fmt.Println("success result:", success)
	fmt.Println("result:", mostCommonWord(paragraph, banned))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
