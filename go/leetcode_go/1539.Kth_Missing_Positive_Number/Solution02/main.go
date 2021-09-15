package main

import (
	"fmt"
)

func findKthPositive(arr []int, k int) int {
	diffNum, dLen := 0, 0
	arrLen := len(arr)
	for i, j := 1, 0; ; i++ {
		if dLen == k {
			return diffNum
		}

		if j >= arrLen || i != arr[j] {
			dLen++
			diffNum = i
		} else if j < arrLen {
			j++
		}
	}
}

func test01() {
	arr := []int{2, 3, 4, 7, 11}
	k := 5
	succResult := 9
	fmt.Println("test01 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findKthPositive(arr, k))
	fmt.Println()
}

func test02() {
	arr := []int{1, 2, 3, 4}
	k := 2
	succResult := 6
	fmt.Println("test02 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findKthPositive(arr, k))
	fmt.Println()
}

func test03() {
	arr := []int{1, 2, 4}
	k := 2
	succResult := 5
	fmt.Println("test03 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findKthPositive(arr, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
