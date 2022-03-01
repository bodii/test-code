package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Ints(cost)

	sum := 0
	for i := len(cost) - 1; i >= 0; {
		if i-1 >= 0 {
			sum += cost[i] + cost[i-1]
			if i-2 >= 0 {
				i -= 3
			} else {
				i -= 2
			}
		} else {
			sum += cost[i]
			i--
		}
	}

	return sum
}

func test01() {
	cost := []int{1, 2, 3}
	succResult := 5
	fmt.Println("test01 cost:=", cost)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumCost(cost))
	fmt.Println()
}

func test02() {
	cost := []int{6, 5, 7, 9, 2, 2}
	succResult := 23
	fmt.Println("test02 cost:=", cost)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumCost(cost))
	fmt.Println()
}

func test03() {
	cost := []int{5, 5}
	succResult := 10
	fmt.Println("test03 cost:=", cost)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumCost(cost))
	fmt.Println()
}

func test04() {
	cost := []int{1}
	succResult := 1
	fmt.Println("test04 cost:=", cost)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumCost(cost))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
