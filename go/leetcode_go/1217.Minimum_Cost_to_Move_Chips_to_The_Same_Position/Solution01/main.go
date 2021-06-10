package main

import (
	"fmt"
)

func minCostToMoveChips(position []int) int {
	odd, even := 0, 0

	for _, v := range position {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if odd > even {
		return even
	}

	return odd
}

func test01() {
	position := []int{1, 2, 3}
	success := 1

	fmt.Println("test01  position:=", position)
	fmt.Println("success result:", success)
	fmt.Println("result:", minCostToMoveChips(position))
	fmt.Println()
}

func test02() {
	position := []int{2, 2, 2, 3, 3}
	success := 2

	fmt.Println("test02  position:=", position)
	fmt.Println("success result:", success)
	fmt.Println("result:", minCostToMoveChips(position))
	fmt.Println()
}

func test03() {
	position := []int{1, 1000000000}
	success := 1

	fmt.Println("test03  position:=", position)
	fmt.Println("success result:", success)
	fmt.Println("result:", minCostToMoveChips(position))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
