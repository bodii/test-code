package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	result := 0
	for i := 1; i < len(prices); i++ {
		result += max(0, prices[i]-prices[i-1])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	prices := []int{7, 1, 5, 3, 6, 4}
	succResult := 7
	fmt.Println("test01 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func test02() {
	prices := []int{1, 2, 3, 4, 5}
	succResult := 4
	fmt.Println("test02 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func test03() {
	prices := []int{7, 6, 4, 3, 1}
	succResult := 0
	fmt.Println("test03 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
