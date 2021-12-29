package main

import (
	"fmt"
)

func maxDistance(colors []int) int {
	firstInd, lastInd := 0, len(colors)-1

	maxDis := 0
	for i := lastInd; i > 0; i-- {
		if colors[0] != colors[i] {
			maxDis = max(maxDis, i)
		}

		if colors[firstInd] != colors[lastInd] {
			maxDis = max(maxDis, lastInd-firstInd)
		}
		firstInd++
	}

	return maxDis
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	prices := []int{1, 1, 1, 6, 1, 1, 1}
	succResult := 3
	fmt.Println("test01 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDistance(prices))
	fmt.Println()
}

func test02() {
	prices := []int{1, 8, 3, 8, 3}
	succResult := 4
	fmt.Println("test02 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDistance(prices))
	fmt.Println()
}

func test03() {
	prices := []int{0, 1}
	succResult := 1
	fmt.Println("test03 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDistance(prices))
	fmt.Println()
}

func test04() {
	prices := []int{4, 4, 4, 11, 4, 4, 11, 4, 4, 4, 4, 4}
	succResult := 8
	fmt.Println("test04 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxDistance(prices))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
