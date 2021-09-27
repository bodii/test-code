package main

import (
	"fmt"
)

func findSpecialInteger(arr []int) int {
	arrLen := len(arr)
	if arrLen < 2 {
		return arr[0]
	}

	max := arrLen / 4
	for i := 0; i < arrLen-max; i++ {
		if arr[i] == arr[i+max] {
			return arr[i]
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
