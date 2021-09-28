package main

import (
	"fmt"
)

func maximumDifference(nums []int) int {
	diffMax := -1
	for i := len(nums) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[i] <= nums[j] {
				continue
			}

			diff := nums[i] - nums[j]
			if diff > diffMax {
				diffMax = diff
			}
		}
	}

	return diffMax
}

func test01() {
	nums := []int{7, 1, 5, 4}
	succResult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func test02() {
	nums := []int{9, 4, 3, 2}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 5, 2, 10}
	succResult := 9
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func test04() {
	nums := []int{0, 0}
	succResult := 0
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumDifference(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
