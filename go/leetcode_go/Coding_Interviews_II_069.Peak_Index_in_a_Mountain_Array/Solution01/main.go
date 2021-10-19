package main

import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	arrLen := len(arr)
	if arrLen == 3 {
		return 1
	}

	mid := arrLen / 2
	for i := 1; i <= mid; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			return i
		}

		if arr[mid+i-2] < arr[mid+i-1] && arr[mid+i-1] > arr[mid+i] {
			return mid + i - 1
		}
	}

	return mid
}

func test01() {
	arr := []int{0, 1, 0}
	success := 1

	fmt.Println("test01  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 3, 5, 4, 2}
	success := 2

	fmt.Println("test02  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func test03() {
	arr := []int{0, 10, 5, 2}
	success := 1

	fmt.Println("test03  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func test04() {
	arr := []int{3, 4, 5, 1}
	success := 2

	fmt.Println("test04  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", peakIndexInMountainArray(arr))
	fmt.Println()
}

func test05() {
	arr := []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19}
	success := 2

	fmt.Println("test05  arr:=", arr)
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
