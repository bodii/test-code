package main

import (
	"fmt"
)

func masterMind(solution string, guess string) []int {
	result := make([]int, 2)

	pseudos := make(map[byte]int)
	for i := 0; i < 4; i++ {
		if solution[i] == guess[i] {
			result[0]++
			continue
		}

		if pseudos[solution[i]] < 0 {
			result[1]++
		}
		pseudos[solution[i]]++

		if pseudos[guess[i]] > 0 {
			result[1]++
		}
		pseudos[guess[i]]--
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
