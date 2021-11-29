package main

import (
	"fmt"
)

func masterMind(solution string, guess string) []int {
	result := make([]int, 2)
	guessKeys, guessNums := make(map[rune][]int), make(map[rune]int)

	for k, v := range guess {
		if guessKeys[v] == nil {
			guessKeys[v] = make([]int, 4)
		}

		if solution[k] == byte(v) {
			guessKeys[v][k] = 1
			result[0]++
		} else {
			guessNums[v]++
		}
	}

	for k, v := range solution {
		if guessNums[v] > 0 && guessKeys[v][k] == 0 {
			result[1]++
			guessNums[v]--
		}
	}

	return result
}

func test01() {
	solution, guess := "RGBY", "GGRR"
	succResult := []int{1, 1}
	fmt.Println("test01 solution:=", solution, " guess:=", guess)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", masterMind(solution, guess))
	fmt.Println()
}

func test02() {
	solution, guess := "RGRB", "BBBY"
	succResult := []int{0, 1}
	fmt.Println("test02 solution:=", solution, " guess:=", guess)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", masterMind(solution, guess))
	fmt.Println()
}

func test03() {
	solution, guess := "BRBB", "RBGY"
	succResult := []int{0, 2}
	fmt.Println("test03 solution:=", solution, " guess:=", guess)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", masterMind(solution, guess))
	fmt.Println()
}

func test04() {
	solution, guess := "BGBG", "RGBR"
	succResult := []int{2, 0}
	fmt.Println("test04 solution:=", solution, " guess:=", guess)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", masterMind(solution, guess))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
