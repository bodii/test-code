package main

import (
	"fmt"
	"sort"
)

func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)

	diff := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
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
