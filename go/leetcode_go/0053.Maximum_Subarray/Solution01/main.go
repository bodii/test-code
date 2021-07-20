package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	numsLen := len(nums)
	if numsLen == 1 {
		return nums[0]
	}

	max, sum := nums[0], 0
	for i := 1; i < numsLen; i++ {
		if sum > 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}

		if max < sum {
			max = sum
		}
	}

	return max
}

func test01() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	succResult := 6
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxSubArray(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1}
	succResult := 1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxSubArray(nums))
	fmt.Println()
}

func test03() {
	nums := []int{0}
	succResult := 0
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxSubArray(nums))
	fmt.Println()
}

func test04() {
	nums := []int{-1}
	succResult := -1
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxSubArray(nums))
	fmt.Println()
}

func test05() {
	nums := []int{-100000}
	succResult := -100000
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxSubArray(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
