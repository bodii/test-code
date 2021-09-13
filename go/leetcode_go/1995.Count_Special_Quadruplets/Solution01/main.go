package main

import (
	"fmt"
)

func countQuadruplets(nums []int) int {
	result := 0
	numsLen := len(nums)
	for a := 0; a < numsLen-3; a++ {
		for b := a + 1; b < numsLen-2; b++ {
			for c := b + 1; c < numsLen-1; c++ {
				for d := c + 1; d < numsLen; d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						result++
					}
				}
			}
		}
	}

	return result
}

func test01() {
	nums := []int{1, 2, 3, 6}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countQuadruplets(nums))
	fmt.Println()
}

func test02() {
	nums := []int{3, 3, 6, 4, 5}
	succResult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countQuadruplets(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 1, 1, 3, 5}
	succResult := 4
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countQuadruplets(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
