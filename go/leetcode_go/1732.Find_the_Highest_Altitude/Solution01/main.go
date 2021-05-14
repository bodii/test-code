package main

import (
	"fmt"
)

func largestAltitude(gain []int) int {
	maxNum := 0
	nums := 0

	for _, v := range gain {
		nums += v
		if nums > maxNum {
			maxNum = nums
		}
	}

	return maxNum
}

func test01() {
	gain := []int{-5, 1, 5, 0, -7}
	succResult := 1
	fmt.Println("test01 gain:=", gain)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestAltitude(gain))
	fmt.Println()
}

func test02() {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	succResult := 0
	fmt.Println("test01 gain:=", gain)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestAltitude(gain))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
