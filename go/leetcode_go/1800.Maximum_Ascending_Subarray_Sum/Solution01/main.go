package main

import (
	"fmt"
)

func maxAscendingSum(nums []int) int {
	numsLen, maxSum := len(nums), nums[0]
	for i := 0; i < numsLen-1; i++ {
		currSum := nums[i]
		if nums[i] >= nums[i+1] {
			if currSum > maxSum {
				maxSum = currSum
			}
			continue
		}

		currSum += nums[i+1]
		for j := i + 2; j < numsLen; j++ {
			if nums[j-1] >= nums[j] {
				break
			}
			currSum += nums[j]
			i = j
		}

		if currSum > maxSum {
			maxSum = currSum
		}

	}

	return maxSum
}

func test01() {
	nums := []int{10, 20, 30, 5, 10, 50}
	succResult := 65
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxAscendingSum(nums))
	fmt.Println()
}

func test02() {
	nums := []int{10, 20, 30, 40, 50}
	succResult := 150
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxAscendingSum(nums))
	fmt.Println()
}

func test03() {
	nums := []int{12, 17, 15, 13, 10, 11, 12}
	succResult := 33
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxAscendingSum(nums))
	fmt.Println()
}

func test04() {
	nums := []int{100, 10, 1}
	succResult := 100
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxAscendingSum(nums))
	fmt.Println()
}

func test05() {
	nums := []int{3, 6, 10, 1, 8, 9, 9, 8, 9}
	succResult := 19
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxAscendingSum(nums))
	fmt.Println()
}

func test06() {
	nums := []int{5, 5, 6, 6, 6, 9, 1, 2}
	succResult := 15
	fmt.Println("test06 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maxAscendingSum(nums))
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
