package main

import (
	"fmt"
	"math"
)

func maximumProduct(nums []int) int {
	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32

	for _, v := range nums {
		if max3 < v {
			if max1 < v {
				max1, max2, max3 = v, max1, max2
			} else if max2 < v {
				max2, max3 = v, max2
			} else {
				max3 = v
			}
		}

		if min2 > v {
			if min1 > v {
				min1, min2 = v, min1
			} else {
				min2 = v
			}
		}
	}

	return max(max1*max2*max3, min1*min2*max1)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	nums := []int{1, 2, 3}
	succResult := 6
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3, 4}
	succResult := 24
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-1, -2, -3}
	succResult := -6
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func test04() {
	nums := []int{-100, -98, -1, 2, 3, 4}
	succResult := 39200
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumProduct(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
