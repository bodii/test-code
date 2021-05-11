package main

import "fmt"

func game(guess []int, answer []int) int {
	correctNums := 0
	for k, v := range guess {
		if isEqual(v, answer[k]) {
			correctNums++
		}
	}

	return correctNums
}

func isEqual(a, b int) bool {
	return a == b
}

func test01() {
	guess := []int{1, 2, 3}
	answer := []int{1, 2, 3}
	succResult := 3

	fmt.Println("test01 guess:=", guess, " answer:=", answer)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", game(guess, answer))
	fmt.Println()
}

func test02() {
	guess := []int{2, 2, 3}
	answer := []int{3, 2, 1}
	succResult := 1

	fmt.Println("test02 guess:=", guess, " answer:=", answer)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", game(guess, answer))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
