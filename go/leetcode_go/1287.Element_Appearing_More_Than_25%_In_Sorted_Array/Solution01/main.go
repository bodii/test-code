package main

import (
	"fmt"
)

func findSpecialInteger(arr []int) int {
	max := len(arr) / 4

	arrMap := make(map[int]int)
	for _, v := range arr {
		arrMap[v]++
	}

	for k, v := range arrMap {
		if v > max {
			return k
		}
	}

	return -1
}

func test01() {
	arr := []int{1, 2, 2, 6, 6, 6, 6, 7, 10}
	success := 6

	fmt.Println("test01 arr: ", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", findSpecialInteger(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 1}
	success := 1

	fmt.Println("test02 arr: ", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", findSpecialInteger(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
