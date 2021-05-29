package main

import (
	"fmt"
)

func threeConsecutiveOdds(arr []int) bool {
	for i := 2; i < len(arr); i++ {
		if arr[i-2]&1 == 1 && arr[i-1]&1 == 1 && arr[i]&1 == 1 {
			return true
		}
	}

	return false
}

func test01() {
	arr := []int{2, 6, 4, 1}
	succResult := false
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", threeConsecutiveOdds(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}
	succResult := true
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", threeConsecutiveOdds(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
