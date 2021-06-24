package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	lastIndex := len(prices) - 1
	max, min := 0, math.MaxInt64
	for i := 0; i <= lastIndex; i++ {
		if min > prices[i] {
			min = prices[i]
		} else if prices[i]-min > max {
			max = prices[i] - min
		}
	}

	return max
}

func test01() {
	prices := []int{7, 1, 5, 3, 6, 4}
	succResult := 5
	fmt.Println("test01 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func test02() {
	prices := []int{7, 6, 4, 3, 1}
	succResult := 0
	fmt.Println("test02 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func test03() {
	prices := []int{2, 4, 1}
	succResult := 2
	fmt.Println("test03 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func test04() {
	prices := []int{1, 2}
	succResult := 1
	fmt.Println("test04 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func test05() {
	prices := []int{2, 1, 2, 1, 0, 1, 2}
	succResult := 2
	fmt.Println("test05 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
	fmt.Println()
}

func test06() {
	prices := []int{3, 2, 6, 5, 0, 3}
	succResult := 4
	fmt.Println("test06 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxProfit(prices))
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
