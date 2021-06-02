package main

import (
	"fmt"
)

func createTargetArray(nums []int, index []int) []int {
	numsLen := len(nums)
	target := make([]int, numsLen)
	for k, v := range nums {
		if target[index[k]] > 0 {
			// tmp := append([]int{v}, target[index[k]:numsLen-1]...)
			// target = append(target[:index[k]], tmp...)
			target = append(target[:index[k]], append([]int{v}, target[index[k]:numsLen-1]...)...)
		} else {
			target[index[k]] = v
		}
	}

	return target
}

func test01() {
	nums, index := []int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}
	succResult := []int{0, 4, 1, 3, 2}
	fmt.Println("test01 nums:=", nums, " index:=", index)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", createTargetArray(nums, index))
	fmt.Println()
}

func test02() {
	nums, index := []int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}
	succResult := []int{0, 1, 2, 3, 4}
	fmt.Println("test02 nums:=", nums, " index:=", index)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", createTargetArray(nums, index))
	fmt.Println()
}

func test03() {
	nums, index := []int{1}, []int{0}
	succResult := []int{1}
	fmt.Println("test03 nums:=", nums, " index:=", index)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", createTargetArray(nums, index))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
