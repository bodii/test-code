package main

import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	len := len(arr)
	if len == 3 {
		return 1
	}

	left, right := 0, len-1
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func test01() {
	arr := []int{0, 1, 0}
	success := 1

	fmt.Println("test01 arr: ", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func test02() {
	arr := []int{0, 2, 1, 0}
	success := 1

	fmt.Println("test02 arr: ", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func test03() {
	arr := []int{0, 10, 5, 2}
	success := 1

	fmt.Println("test03 arr: ", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func test04() {
	arr := []int{3, 4, 5, 1}
	success := 2

	fmt.Println("test04 arr: ", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}
func test05() {
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	success := 2

	fmt.Println("test05 arr: ", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
