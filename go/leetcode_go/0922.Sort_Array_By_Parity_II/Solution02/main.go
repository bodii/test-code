package main

import (
	"fmt"
)

func sortArrayByParityII(nums []int) []int {
	numsLen := len(nums)
	for i := 0; i < numsLen-1; i++ {
		if i%2 == 0 && nums[i]%2 == 1 {
			for k := i; k < numsLen; k++ {
				if nums[k]%2 == 0 {
					nums[i], nums[k] = nums[k], nums[i]
					break
				}
			}
		} else if i%2 == 1 && nums[i]%2 == 0 {
			for k := i; k < numsLen; k++ {
				if nums[k]%2 == 1 {
					nums[i], nums[k] = nums[k], nums[i]
					break
				}
			}
		}
	}

	return nums
}

func test01() {
	nums := []int{4, 2, 5, 7}
	success := []int{4, 5, 2, 7}

	fmt.Println("test01  nums:=", nums)
	fmt.Println("success result:", success)
	fmt.Println("result:", sortArrayByParityII(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 3}
	success := []int{2, 3}

	fmt.Println("test02  nums:=", nums)
	fmt.Println("success result:", success)
	fmt.Println("result:", sortArrayByParityII(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
