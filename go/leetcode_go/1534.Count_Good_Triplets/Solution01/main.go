package main

import (
	"fmt"
)

func countGoodTriplets(arr []int, a int, b int, c int) int {
	arrLen := len(arr)

	result := 0

	abs := func(i int) int {
		if i < 0 {
			return -i
		}
		return i
	}

	for i := 0; i < arrLen-2; i++ {
		for j := i + 1; j < arrLen-1; j++ {
			for k := j + 1; k < arrLen; k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					// fmt.Println(arr[i], arr[j], arr[k])
					result++
				}
			}
		}
	}

	return result
}

func test01() {
	arr := []int{3, 0, 1, 1, 9, 7}
	a, b, c := 7, 2, 3
	succResult := 4
	fmt.Println("test01 arr:=", arr)
	fmt.Println("a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countGoodTriplets(arr, a, b, c))
	fmt.Println()
}

func test02() {
	arr := []int{1, 1, 2, 2, 3}
	a, b, c := 0, 0, 1
	succResult := 0
	fmt.Println("test02 arr:=", arr)
	fmt.Println("a:=", a, " b:=", b, " c:=", c)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countGoodTriplets(arr, a, b, c))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
