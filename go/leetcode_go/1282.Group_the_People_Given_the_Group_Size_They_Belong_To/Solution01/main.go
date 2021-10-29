package main

import (
	"fmt"
)

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)

	for k, v := range groupSizes {
		groups[v] = append(groups[v], k)
	}

	result := [][]int{}
	for k, v := range groups {
		len, l := len(v), 0
		for l < len {
			result = append(result, v[l:k+l])
			l += k
		}
	}

	return result
}

func test01() {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	successResult := [][]int{{5}, {0, 1, 2}, {3, 4, 6}}
	fmt.Println("test01 groupSizes:", groupSizes)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", groupThePeople(groupSizes))
	fmt.Println()
}

func test02() {
	groupSizes := []int{2, 1, 3, 3, 3, 2}
	successResult := [][]int{{1}, {0, 5}, {2, 3, 4}}
	fmt.Println("test02 groupSizes:", groupSizes)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", groupThePeople(groupSizes))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
