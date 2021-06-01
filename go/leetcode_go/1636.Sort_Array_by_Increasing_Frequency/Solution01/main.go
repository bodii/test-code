package main

import (
	"fmt"
	"sort"
)

func frequencySort(nums []int) []int {
	numsMap := make(map[int]int)
	numsKeyMap := make(map[int][]int)

	for _, v := range nums {
		numsMap[v]++
	}

	for k, v := range numsMap {
		numsKeyMap[v] = append(numsKeyMap[v], k)
	}

	numsSort := make([]int, 0)
	for k, v := range numsKeyMap {
		numsSort = append(numsSort, k)
		sort.Sort(sort.Reverse(sort.IntSlice(v)))
		numsKeyMap[k] = v
	}

	sort.Ints(numsSort)

	result := make([]int, 0)
	for _, v := range numsSort {
		for _, v2 := range numsKeyMap[v] {
			for i := 0; i < v; i++ {
				result = append(result, v2)

			}
		}
	}

	return result
}

func test01() {
	nums := []int{1, 1, 2, 2, 2, 3}
	succResult := []int{3, 1, 1, 2, 2, 2}
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", frequencySort(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 3, 1, 3, 2}
	succResult := []int{1, 3, 3, 2, 2}
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", frequencySort(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	succResult := []int{5, -1, 4, 4, -6, -6, 1, 1, 1}
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", frequencySort(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
