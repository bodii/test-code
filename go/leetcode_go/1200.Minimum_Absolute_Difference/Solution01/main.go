package main

import (
	"fmt"
	"sort"
)

func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)

	result := make([][]int, 0)
	minDiff := arr[1] - arr[0]
	result = append(result, []int{arr[0], arr[1]})
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] == minDiff {
			result = append(result, []int{arr[i-1], arr[i]})
		}

		if arr[i]-arr[i-1] < minDiff {
			minDiff = arr[i] - arr[i-1]
			result = append([][]int{}, []int{arr[i-1], arr[i]})
		}
	}

	return result
}

func test01() {
	arr := []int{4, 2, 1, 3}
	successResult := [][]int{{1, 2}, {2, 3}, {3, 4}}
	fmt.Println("test01 arr:", arr)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", minimumAbsDifference(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 3, 6, 10, 15}
	successResult := [][]int{{1, 3}}
	fmt.Println("test02 arr:", arr)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", minimumAbsDifference(arr))
	fmt.Println()
}

func test03() {
	arr := []int{3, 8, -10, 23, 19, -4, -14, 27}
	successResult := [][]int{{-14, -10}, {19, 23}, {23, 27}}
	fmt.Println("test03 arr:", arr)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", minimumAbsDifference(arr))
	fmt.Println()
}

func test04() {
	arr := []int{40, 11, 26, 27, -20}
	successResult := [][]int{{26, 27}}
	fmt.Println("test04 arr:", arr)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", minimumAbsDifference(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
