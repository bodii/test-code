package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	result := 0
	nums := []int{}
	numsIndx := -1
	for _, o := range ops {
		n := 0
		if o == "+" {
			n = nums[numsIndx] + nums[numsIndx-1]
			nums = append(nums, n)
			numsIndx++
		} else if o == "D" {
			n = nums[numsIndx] * 2
			nums = append(nums, n)
			numsIndx++
		} else if o == "C" {
			n = -nums[numsIndx]
			nums = nums[0:numsIndx]
			numsIndx--
		} else {
			n, _ = strconv.Atoi(o)
			nums = append(nums, n)
			numsIndx++
		}
		result += n
	}

	return result
}

func test01() {
	ops := []string{"5", "2", "C", "D", "+"}
	succResult := 30
	fmt.Println("test01 ops:=", ops)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", calPoints(ops))
	fmt.Println()
}

func test02() {
	ops := []string{"5", "-2", "4", "C", "D", "9", "+", "+"}
	succResult := 27
	fmt.Println("test02 ops:=", ops)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", calPoints(ops))
	fmt.Println()
}

func test03() {
	ops := []string{"1"}
	succResult := 1
	fmt.Println("test03 ops:=", ops)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", calPoints(ops))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
