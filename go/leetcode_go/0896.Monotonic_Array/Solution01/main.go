package main

import (
	"fmt"
)

func isMonotonic(nums []int) bool {
	lastIndex := len(nums) - 1
	isUp := nums[lastIndex] > nums[0]

	for i := 0; i < lastIndex; i++ {
		if isUp {
			if nums[i] > nums[i+1] {
				return false
			}
		} else {
			if nums[i] < nums[i+1] {
				return false
			}
		}
	}

	return true
}

func test01() {
	prices := []int{1, 2, 2, 3}
	succResult := true
	fmt.Println("test01 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isMonotonic(prices))
	fmt.Println()
}

func test02() {
	prices := []int{6, 5, 4, 4}
	succResult := true
	fmt.Println("test02 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isMonotonic(prices))
	fmt.Println()
}

func test03() {
	prices := []int{1, 3, 2}
	succResult := false
	fmt.Println("test03 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isMonotonic(prices))
	fmt.Println()
}

func test04() {
	prices := []int{1, 2, 4, 5}
	succResult := true
	fmt.Println("test04 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isMonotonic(prices))
	fmt.Println()
}

func test05() {
	prices := []int{1, 1, 1}
	succResult := true
	fmt.Println("test05 prices:=", prices)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isMonotonic(prices))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
