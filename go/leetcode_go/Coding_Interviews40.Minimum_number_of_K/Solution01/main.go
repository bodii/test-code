package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[0:k]
}

func test01() {
	arr := []int{1, 2, 5, 9, 5, 9, 5, 5, 5}
	k := 2
	succResult := []int{1, 2}
	fmt.Println("test01 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getLeastNumbers(arr, k))
	fmt.Println()
}

func test02() {
	arr := []int{0, 1, 2, 1}
	k := 1
	succResult := []int{0}
	fmt.Println("test02 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getLeastNumbers(arr, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
