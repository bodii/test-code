package main

import (
	"fmt"
)

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	num := 0
	for _, v1 := range arr1 {
		dis := true
		for _, v2 := range arr2 {
			if abs(v1-v2) <= d {
				dis = false
				break
			}
		}

		if dis {
			num++
		}
	}

	return num
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func test01() {
	arr1, arr2 := []int{4, 5, 8}, []int{10, 9, 1, 8}
	d := 2
	successResult := 2
	fmt.Println("test01 arr1:", arr1, " arr2:", arr2)
	fmt.Println("d:=", d)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", findTheDistanceValue(arr1, arr2, d))
	fmt.Println()
}

func test02() {
	arr1, arr2 := []int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}
	d := 3
	successResult := 2
	fmt.Println("test02 arr1:", arr1, " arr2:", arr2)
	fmt.Println("d:=", d)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", findTheDistanceValue(arr1, arr2, d))
	fmt.Println()
}

func test03() {
	arr1, arr2 := []int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}
	d := 6
	successResult := 1
	fmt.Println("test03 arr1:", arr1, " arr2:", arr2)
	fmt.Println("d:=", d)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", findTheDistanceValue(arr1, arr2, d))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
