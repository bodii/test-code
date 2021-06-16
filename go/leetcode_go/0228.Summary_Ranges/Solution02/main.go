package main

import (
	"fmt"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	left, right, rangeIndex := 0, 0, len(nums)

	for right < rangeIndex {
		if right < rangeIndex-1 && nums[right]+1 == nums[right+1] {
			right++
		} else {
			if left == right {
				result = append(result, fmt.Sprintf("%d", nums[left]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[left], nums[right]))
			}
			right++
			left = right
		}
	}

	return result
}

func test01() {
	nums := []int{0, 1, 2, 4, 5, 7}
	succResult := []string{"0->2", "4->5", "7"}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", summaryRanges(nums))
	fmt.Println()
}

func test02() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	succResult := []string{"0", "2->4", "6", "8->9"}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", summaryRanges(nums))
	fmt.Println()
}

func test03() {
	nums := []int{}
	succResult := []string{}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", summaryRanges(nums))
	fmt.Println()
}

func test04() {
	nums := []int{-1}
	succResult := []string{"-1"}
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", summaryRanges(nums))
	fmt.Println()
}

func test05() {
	nums := []int{0}
	succResult := []string{"0"}
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", summaryRanges(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
