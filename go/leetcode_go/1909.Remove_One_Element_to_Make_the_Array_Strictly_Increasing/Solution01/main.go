package main

import (
	"fmt"
)

func canBeIncreasing(nums []int) bool {
	notDecrease := 0

	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] <= nums[i-1] {
			notDecrease++
			if notDecrease < 2 && i-2 >= 0 && i+1 < len(nums) &&
				nums[i] <= nums[i-2] && nums[i+1] <= nums[i-1] {
				notDecrease++
			}

			if notDecrease > 1 {
				return false
			}
		}
	}

	return true
}

func test01() {
	nums := []int{1, 2, 10, 5, 7}
	succResult := true
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 3, 1, 2}
	succResult := false
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 1, 1}
	succResult := false
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test04() {
	nums := []int{1, 2, 3}
	succResult := true
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test05() {
	nums := []int{105, 924, 32, 968}
	succResult := true
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test06() {
	nums := []int{1, 3, 2}
	succResult := true
	fmt.Println("test06 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test07() {
	nums := []int{4, 3, 1}
	succResult := false
	fmt.Println("test07 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test08() {
	nums := []int{2, 3, 1}
	succResult := true
	fmt.Println("test08 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test09() {
	nums := []int{2, 3, 1, 2}
	succResult := false
	fmt.Println("test09 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test10() {
	nums := []int{2, 3, 1, 4}
	succResult := true
	fmt.Println("test10 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test11() {
	nums := []int{541, 783, 433, 744}
	succResult := false
	fmt.Println("test11 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func test12() {
	nums := []int{226, 529, 646, 712, 817, 853, 1000, 229, 665}
	succResult := false
	fmt.Println("test12 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeIncreasing(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
	test09()
	test10()
	test11()
	test12()
}
