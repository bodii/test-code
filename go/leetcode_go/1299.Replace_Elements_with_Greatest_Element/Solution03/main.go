package main

import (
	"fmt"
)

func replaceElements(arr []int) []int {
	arrLen := len(arr)
	lastIndex := arrLen - 1
	lMax, rMax := 0, arr[lastIndex]
	arr[lastIndex] = -1
	if lastIndex == 0 {
		return arr
	}

	for i := arrLen - 2; i >= 0; i-- {
		lMax, arr[i] = arr[i], rMax
		if lMax > rMax {
			rMax = lMax
		}
	}

	return arr
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

func test03() {
	arr := []int{56903, 18666, 60499, 57517, 26961}
	succResult := []int{60499, 60499, 57517, 26961, -1}
	fmt.Println("test03 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", replaceElements(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
