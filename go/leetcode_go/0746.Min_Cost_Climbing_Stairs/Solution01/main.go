package main

import (
	"fmt"
)

func minCostClimbingStairs(cost []int) int {
	pre, cur := 0, 0
	for i := 2; i <= len(cost); i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}

	return cur
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
