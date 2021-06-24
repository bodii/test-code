package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	lastIndex := len(prices) - 1
	max, min := -1, -1
	result := 0
	for i := 0; i < lastIndex; i++ {
		if min == -1 && prices[i] < prices[i+1] {
			min = prices[i]
			max = prices[i+1]
		} else if min != -1 && prices[i] < prices[i+1] {
			max = prices[i+1]
		} else if min != -1 && max != -1 && prices[i] > prices[i+1] {
			result += max - min
			min = -1
			max = -1
		}
	}
	result += max - min
	return result
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
