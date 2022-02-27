package main

import (
	"fmt"
)

func lexicalOrder(n int) []int {
	result := make([]int, 0)

	var dfs func(start int)
	dfs = func(start int) {
		if start > n {
			return
		}

		for i := 0; i < 10 && start+i <= n; i++ {
			if start == 1 && i == 9 {
				break
			}
			result = append(result, start+i)
			dfs((start + i) * 10)
		}
	}

	dfs(1)

	return result
}

func test01() {
	n := 13
	succResult := []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lexicalOrder(n))
	fmt.Println()
}

func test02() {
	n := 2
	succResult := []int{1, 2}

	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lexicalOrder(n))
	fmt.Println()
}

func test03() {
	n := 34
	succResult := []int{1, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
		2, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 3, 30, 31, 32,
		33, 34, 4, 5, 6, 7, 8, 9}

	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lexicalOrder(n))
	fmt.Println()
}

func test04() {
	n := 101
	succResult := []int{1, 10, 100, 101, 11, 12, 13, 14, 15, 16,
		17, 18, 19, 2, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 3,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 4, 40, 41, 42, 43,
		44, 45, 46, 47, 48, 49, 5, 50, 51, 52, 53, 54, 55, 56, 57,
		58, 59, 6, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 7, 70,
		71, 72, 73, 74, 75, 76, 77, 78, 79, 8, 80, 81, 82, 83, 84,
		85, 86, 87, 88, 89, 9, 90, 91, 92, 93, 94, 95, 96, 97, 98,
		99}

	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", lexicalOrder(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
