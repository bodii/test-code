package main

import (
	"fmt"
)

func replaceElements(arr []int) []int {
	arrLen := len(arr)
	lastIndex := arrLen - 1
	result := make([]int, arrLen)
	result[lastIndex] = -1
	if lastIndex == 0 {
		return result
	}

	for i := 0; i < lastIndex; i++ {
		result[i] = findMax(arr[i+1:])
	}

	return result
}

func findMax(subArr []int) int {
	max := subArr[0]
	for i := 1; i < len(subArr); i++ {
		if subArr[i] > max {
			max = subArr[i]
		}
	}

	return max
}

func test01() {
	arr := []int{17, 18, 5, 4, 6, 1}
	succResult := []int{18, 6, 6, 6, 1, -1}
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceElements(arr))
	fmt.Println()
}

func test02() {
	arr := []int{400}
	succResult := []int{-1}
	fmt.Println("test02 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceElements(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
