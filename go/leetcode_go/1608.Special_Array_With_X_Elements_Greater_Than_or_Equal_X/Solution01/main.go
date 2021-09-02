package main

import (
	"fmt"
	"sort"
)

func specialArray(nums []int) int {
	l := len(nums)
	sort.Ints(nums)

	for x := l; x > 0; x-- {
		sum := 0
		for j := len(nums) - 1; j >= 0; j-- {
			if nums[j] < x {
				break
			}
			sum++
			if sum > x {
				break
			}
		}

		if x == sum {
			return x
		}
	}

	return -1
}

func test01() {
	nums := []int{3, 5}
	succResult := 2
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", specialArray(nums))
	fmt.Println()
}

func test02() {
	nums := []int{0, 0}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", specialArray(nums))
	fmt.Println()
}

func test03() {
	nums := []int{0, 4, 3, 0, 4}
	succResult := 3
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", specialArray(nums))
	fmt.Println()
}

func test04() {
	nums := []int{3, 6, 7, 7, 0}
	succResult := -1
	fmt.Println("test04 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", specialArray(nums))
	fmt.Println()
}

func test05() {
	nums := []int{3, 9, 7, 8, 3, 8, 6, 6}
	succResult := 6
	fmt.Println("test05 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", specialArray(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
