package main

import (
	"fmt"
)

func camelMatch(queries []string, pattern string) []bool {
	result := make([]bool, len(queries))
	patternLen := len(pattern)

	for k, v := range queries {
		result[k] = match(v, pattern, patternLen)
	}

	return result
}

func match(word, pattern string, pLen int) bool {
	pInd := 0
	for k := range word {
		if pInd < pLen && word[k] == pattern[pInd] {
			pInd++
		} else if isUpper(word[k]) {
			return false
		}
	}

	return pInd == pLen
}

func isUpper(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func test01() {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FB"
	succResult := []bool{true, false, true, true, false}
	fmt.Println("test01 queries:=", queries, " pattern:=", pattern)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", camelMatch(queries, pattern))
	fmt.Println()
}

func test02() {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FoBa"
	succResult := []bool{true, false, true, false, false}
	fmt.Println("test02 queries:=", queries, " pattern:=", pattern)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", camelMatch(queries, pattern))
	fmt.Println()
}

func test03() {
	queries := []string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}
	pattern := "FoBaT"
	succResult := []bool{false, true, false, false, false}
	fmt.Println("test03 queries:=", queries, " pattern:=", pattern)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", camelMatch(queries, pattern))
	fmt.Println()
}

func test04() {
	queries := []string{"CompetitiveProgramming", "CounterPick", "ControlPanel"}
	pattern := "CooP"
	succResult := []bool{false, false, true}
	fmt.Println("test04 queries:=", queries, " pattern:=", pattern)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", camelMatch(queries, pattern))
	fmt.Println()
}

func test05() {
	queries := []string{"aksvbjLiknuTzqon", "ksvjLimflkpnTzqn", "mmkasvjLiknTxzqn", "ksvjLiurknTzzqbn", "ksvsjLctikgnTzqn", "knzsvzjLiknTszqn"}
	pattern := "ksvjLiknTzqn"
	succResult := []bool{true, true, true, true, true, true}
	fmt.Println("test05 queries:=", queries, " pattern:=", pattern)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", camelMatch(queries, pattern))
	fmt.Println()
}

func test06() {
	queries := []string{"uAxaqlzahfialcezsLfj", "cAqlzyahaslccezssLfj", "AqlezahjarflcezshLfj", "AqlzofahaplcejuzsLfj", "tAqlzahavslcezsLwzfj", "AqlzahalcerrzsLpfonj", "AqlzahalceaczdsosLfj", "eAqlzbxahalcezelsLfj"}
	pattern := "AqlzahalcezsLfj"
	succResult := []bool{true, true, true, true, true, true, true, true}
	fmt.Println("test06 queries:=", queries, " pattern:=", pattern)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", camelMatch(queries, pattern))
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
