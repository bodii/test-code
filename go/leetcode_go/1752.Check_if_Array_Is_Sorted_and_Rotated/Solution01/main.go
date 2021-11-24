package main

import (
	"fmt"
)

func check(nums []int) bool {
	numsLen := len(nums)
	isSorted := nums[0] >= nums[numsLen-1]

	for i := 0; i < numsLen-1; i++ {
		if nums[i] > nums[i+1] {
			if !isSorted {
				return false
			}
			isSorted = false
		}
	}

	return true
}

func test01() {
	nums := []int{3, 4, 5, 1, 2}
	succResult := true
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", check(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 1, 3, 4}
	succResult := false
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", check(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 3}
	succResult := true
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", check(nums))
	fmt.Println()
}

func test04() {
	nums := []int{1, 1, 1}
	succResult := true
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", check(nums))
	fmt.Println()
}

func test05() {
	nums := []int{2, 1}
	succResult := true
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", check(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
