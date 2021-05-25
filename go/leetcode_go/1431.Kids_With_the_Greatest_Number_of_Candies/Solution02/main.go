package main

import (
	"fmt"
)

func kidsWithCandies(candies []int, extraCandies int) []bool {
	candiesLen := len(candies)

	max := func(nums []int, len int) int {
		max := nums[0]
		for i := 1; i < len; i++ {
			if nums[i] > max {
				max = nums[i]
			}
		}

		return max
	}(candies, candiesLen)

	result := make([]bool, candiesLen)
	for k, v := range candies {
		result[k] = v+extraCandies >= max
	}

	return result
}

func test01() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	succResult := []bool{true, true, true, false, true}
	fmt.Println("test01 candies:=", candies, " extraCandies:=", extraCandies)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kidsWithCandies(candies, extraCandies))
	fmt.Println()
}

func test02() {
	candies := []int{4, 2, 1, 1, 2}
	extraCandies := 1
	succResult := []bool{true, false, false, false, false}
	fmt.Println("test02 candies:=", candies, " extraCandies:=", extraCandies)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kidsWithCandies(candies, extraCandies))
	fmt.Println()
}

func test03() {
	candies := []int{12, 1, 12}
	extraCandies := 10
	succResult := []bool{true, false, true}
	fmt.Println("test03 candies:=", candies, " extraCandies:=", extraCandies)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kidsWithCandies(candies, extraCandies))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
