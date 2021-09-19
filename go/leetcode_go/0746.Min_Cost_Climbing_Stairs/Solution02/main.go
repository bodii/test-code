package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	prev, curr := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		prev, curr = curr, min(prev, curr)+cost[i]
	}

	return min(prev, curr)
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func test01() {
	cost := []int{10, 15, 20}
	succResult := 15
	fmt.Println("test01 cost:=", cost)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minCostClimbingStairs(cost))
	fmt.Println()
}

func test02() {
	cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	succResult := 6
	fmt.Println("test02 cost:=", cost)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minCostClimbingStairs(cost))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
