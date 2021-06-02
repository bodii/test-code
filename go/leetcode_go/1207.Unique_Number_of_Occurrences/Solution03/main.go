package main

import (
	"fmt"
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	sort.Ints(arr)

	nums := make(map[int]int)
	cur, curNum := arr[0], 1
	for i := 1; i < len(arr); i++ {
		if arr[i] == cur {
			curNum++
			continue
		}

		if _, ok := nums[curNum]; ok {
			return false
		}
		nums[curNum] = 1

		cur, curNum = arr[i], 1
	}

	if _, ok := nums[curNum]; ok {
		return false
	}

	return true
}

func test01() {
	arr := []int{1, 2, 2, 1, 1, 3}
	success := true

	fmt.Println("test01  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", uniqueOccurrences(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 2}
	success := false

	fmt.Println("test02  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", uniqueOccurrences(arr))
	fmt.Println()
}

func test03() {
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	success := true

	fmt.Println("test03  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", uniqueOccurrences(arr))
	fmt.Println()
}

func test04() {
	arr := []int{26, 2, 16, 16, 5, 5, 26, 2, 5, 20, 20, 5, 2, 20, 2, 2, 20, 2, 16, 20, 16, 17, 16, 2, 16, 20, 26, 16}
	success := false

	fmt.Println("test04  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", uniqueOccurrences(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
