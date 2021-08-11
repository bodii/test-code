package main

import (
	"fmt"
	"math"
	"sort"
)

func trimMean(arr []int) float64 {
	sort.Ints(arr)

	arrLen := len(arr)
	min := int(math.Ceil(float64(arrLen) * 0.05))

	sum, sumLen := 0, 0
	for i := min; i < arrLen-min; i++ {
		sum += arr[i]
		sumLen++
	}

	return float64(sum) / float64(sumLen)
}

func test01() {
	arr := []int{1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3}
	succResult := 2.00000
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trimMean(arr))
	fmt.Println()
}

func test02() {
	arr := []int{6, 2, 7, 5, 1, 2, 0, 3, 10, 2, 5, 0, 5, 5, 0, 8, 7, 6, 8, 0}
	succResult := 4.00000
	fmt.Println("test02 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trimMean(arr))
	fmt.Println()
}

func test03() {
	arr := []int{6, 0, 7, 0, 7, 5, 7, 8, 3, 4, 0, 7, 8, 1, 6, 8, 1, 1, 2, 4, 8, 1, 9, 5, 4, 3, 8, 5, 10, 8, 6, 6, 1, 0, 6, 10, 8, 2, 3, 4}
	succResult := 4.77778
	fmt.Println("test03 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trimMean(arr))
	fmt.Println()
}

func test04() {
	arr := []int{9, 7, 8, 7, 7, 8, 4, 4, 6, 8, 8, 7, 6, 8, 8, 9, 2, 6, 0, 0, 1, 10, 8, 6, 3, 3, 5, 1, 10, 9, 0, 7, 10, 0, 10, 4, 1, 10, 6, 9, 3, 6, 0, 0, 2, 7, 0, 6, 7, 2, 9, 7, 7, 3, 0, 1, 6, 1, 10, 3}
	succResult := 5.27778
	fmt.Println("test04 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trimMean(arr))
	fmt.Println()
}

func test05() {
	arr := []int{4, 8, 4, 10, 0, 7, 1, 3, 7, 8, 8, 3, 4, 1, 6, 2, 1, 1, 8, 0, 9, 8, 0, 3, 9, 10, 3, 10, 1, 10, 7, 3, 2, 1, 4, 9, 10, 7, 6, 4, 0, 8, 5, 1, 2, 1, 6, 2, 5, 0, 7, 10, 9, 10, 3, 7, 10, 5, 8, 5, 7, 6, 7, 6, 10, 9, 5, 10, 5, 5, 7, 2, 10, 7, 7, 8, 2, 0, 1, 1}
	succResult := 5.29167
	fmt.Println("test05 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", trimMean(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
