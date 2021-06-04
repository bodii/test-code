package main

import (
	"fmt"
)

func sortArrayByParity(nums []int) []int {
	numsLen := len(nums)
	left, right := 0, numsLen-1
	for left < right {
		if nums[left]%2 != 0 {
			for right > left {
				if nums[right]%2 == 0 {
					nums[left], nums[right] = nums[right], nums[left]
					break
				} else {
					right--
				}
			}
		}
		left++
	}

	return nums
}

func test01() {
	nums := []int{3, 1, 2, 4}
	success := []int{2, 4, 3, 1}

	fmt.Println("test01  nums:=", nums)
	fmt.Println("success result:", success)
	fmt.Println("result:", sortArrayByParity(nums))
	fmt.Println()
}

func main() {
	test01()
}
