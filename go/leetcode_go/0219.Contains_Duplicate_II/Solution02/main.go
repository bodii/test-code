package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && j-i <= k {
				return true
			}
		}
	}

	return false
}

func test01() {
	nums := []int{1, 2, 3, 1}
	k := 3
	succResult := true
	fmt.Println("test01 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsNearbyDuplicate(nums, k))
	fmt.Println()
}

func test02() {
	nums := []int{1, 0, 1, 1}
	k := 1
	succResult := true
	fmt.Println("test02 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsNearbyDuplicate(nums, k))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	succResult := false
	fmt.Println("test02 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsNearbyDuplicate(nums, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
