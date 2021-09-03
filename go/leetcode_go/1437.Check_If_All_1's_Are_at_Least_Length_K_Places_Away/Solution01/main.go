package main

import (
	"fmt"
)

func kLengthApart(nums []int, k int) bool {
	if k == 0 {
		return true
	}

	prev := -1
	for i, v := range nums {
		if v == 1 {
			if prev != -1 && i-prev-1 < k {
				return false
			}
			prev = i
		}
	}

	return true
}

func test01() {
	nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	k := 2
	succResult := true
	fmt.Println("test01 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kLengthApart(nums, k))
	fmt.Println()
}

func test02() {
	nums := []int{1, 0, 0, 1, 0, 1}
	k := 2
	succResult := false
	fmt.Println("test02 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kLengthApart(nums, k))
	fmt.Println()
}

func test03() {
	nums := []int{1, 1, 1, 1, 1}
	k := 0
	succResult := true
	fmt.Println("test03 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kLengthApart(nums, k))
	fmt.Println()
}

func test04() {
	nums := []int{0, 1, 0, 1}
	k := 1
	succResult := true
	fmt.Println("test04 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kLengthApart(nums, k))
	fmt.Println()
}

func test05() {
	nums := []int{1, 1, 1, 0}
	k := 3
	succResult := false
	fmt.Println("test05 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kLengthApart(nums, k))
	fmt.Println()
}

func test06() {
	nums := []int{1, 0, 0, 0, 1, 0, 0, 1, 0}
	k := 2
	succResult := true
	fmt.Println("test06 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kLengthApart(nums, k))
	fmt.Println()
}

func test07() {
	nums := []int{1, 0, 1}
	k := 2
	succResult := false
	fmt.Println("test07 nums:=", nums, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kLengthApart(nums, k))
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
}
