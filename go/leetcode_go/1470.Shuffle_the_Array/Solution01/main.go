package main

import (
	"fmt"
)

func shuffle(nums []int, n int) []int {
	res := make([]int, n*2)

	for i := 0; i < n; i++ {
		res[i*2] = nums[i]
		res[i*2+1] = nums[i+n]
	}

	return res
}

func test01() {
	nums := []int{2, 5, 1, 3, 4, 7}
	n := 3
	succResult := []int{2, 3, 5, 4, 1, 7}
	fmt.Println("test01 nums:=", nums, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shuffle(nums, n))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n := 4
	succResult := []int{1, 4, 2, 3, 3, 2, 4, 1}
	fmt.Println("test02 nums:=", nums, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shuffle(nums, n))
	fmt.Println()
}

func test03() {
	nums := []int{1, 1, 2, 2}
	n := 2
	succResult := []int{1, 2, 1, 2}
	fmt.Println("test03 nums:=", nums, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", shuffle(nums, n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
