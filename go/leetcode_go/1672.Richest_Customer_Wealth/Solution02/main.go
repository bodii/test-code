package main

import (
	"fmt"
)

func maximumWealth(accounts [][]int) int {
	res := 0

	for _, v := range accounts {
		res = max(res, sum(v))
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func sum(ints []int) int {
	count := 0

	for _, v := range ints {
		count += v
	}

	return count
}

func test01() {
	accounts := [][]int{{1, 2, 3}, {3, 2, 1}}
	succResult := 6
	fmt.Println("test01 accounts:=", accounts)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumWealth(accounts))
	fmt.Println()
}

func test02() {
	accounts := [][]int{{1, 5}, {7, 3}, {3, 5}}
	succResult := 10
	fmt.Println("test02 accounts:=", accounts)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumWealth(accounts))
	fmt.Println()
}

func test03() {
	accounts := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	succResult := 17
	fmt.Println("test03 accounts:=", accounts)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumWealth(accounts))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
