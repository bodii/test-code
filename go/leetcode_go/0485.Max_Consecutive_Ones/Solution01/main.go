package main

import (
	"fmt"
)

func findMaxConsecutiveOnes(nums []int) int {
	max, currentNums := 0, 0

	for _, v := range nums {
		if v == 0 {
			if max < currentNums {
				max = currentNums
			}
			currentNums = 0
		} else {
			currentNums++
		}
	}

	if max < currentNums {
		return currentNums
	}

	return max
}

func test01() {
	nums := []int{1, 1, 0, 1, 1, 1}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findMaxConsecutiveOnes(nums))
	fmt.Println()
}

func main() {
	test01()
}
