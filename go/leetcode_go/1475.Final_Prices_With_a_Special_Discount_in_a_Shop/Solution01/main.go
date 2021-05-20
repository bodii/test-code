package main

import (
	"fmt"
)

func finalPrices(prices []int) []int {
	pricesLen := len(prices)
	for j := 0; j < pricesLen-1; j++ {
		for i := j + 1; i < pricesLen; i++ {
			if prices[j] >= prices[i] {
				prices[j] -= prices[i]
				break
			}
		}
	}

	return prices
}

func test01() {
	prices := []int{8, 4, 6, 2, 3}
	succResult := []int{4, 2, 4, 2, 3}
	fmt.Println("test01 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", finalPrices(prices))
	fmt.Println()
}

func test02() {
	prices := []int{1, 2, 3, 4, 5}
	succResult := []int{1, 2, 3, 4, 5}
	fmt.Println("test02 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", finalPrices(prices))
	fmt.Println()
}

func test03() {
	prices := []int{10, 1, 1, 6}
	succResult := []int{9, 0, 1, 6}
	fmt.Println("test03 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", finalPrices(prices))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
