package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	kingNums := 0
	sort.Ints(nums)
	for i := 0; i < 4; i++ {
		if nums[i] == 0 {
			kingNums++
		} else if nums[i] == nums[i+1] || nums[i+1]-nums[i]-1 > kingNums {
			return false
		} else if kingNums > 0 && nums[i+1]-nums[i]-1 == kingNums {
			kingNums = 0
		}
	}

	return true
}

func test01() {
	nums := []int{1, 2, 3, 4, 5}
	succResult := true
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isStraight(nums))
	fmt.Println()
}

func test02() {
	nums := []int{0, 0, 1, 2, 5}
	succResult := true
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isStraight(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
