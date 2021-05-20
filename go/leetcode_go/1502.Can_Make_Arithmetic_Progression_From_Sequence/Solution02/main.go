package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	for i := 1; i < len(arr)-1; i++ {
		if arr[i]*2 != arr[i-1]+arr[i+1] {
			return false
		}
	}

	return true
}

func test01() {
	arr := []int{3, 5, 1}
	succResult := true
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canMakeArithmeticProgression(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 2, 4}
	succResult := false
	fmt.Println("test02 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canMakeArithmeticProgression(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
