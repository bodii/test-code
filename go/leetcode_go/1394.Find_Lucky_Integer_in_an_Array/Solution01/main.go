package main

import (
	"fmt"
)

func findLucky(arr []int) int {
	numsMap := make(map[int]int)

	for _, v := range arr {
		numsMap[v]++
	}

	maxLuck := -1
	for l, n := range numsMap {
		if l == n && n > maxLuck {
			maxLuck = n
		}
	}

	return maxLuck
}

func test01() {
	arr := []int{2, 2, 3, 4}
	success := 2

	fmt.Println("test01  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", findLucky(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1, 2, 2, 3, 3, 3}
	success := 3

	fmt.Println("test02  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", findLucky(arr))
	fmt.Println()
}

func test03() {
	arr := []int{2, 2, 2, 3, 3}
	success := -1

	fmt.Println("test03  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", findLucky(arr))
	fmt.Println()
}

func test04() {
	arr := []int{5}
	success := -1

	fmt.Println("test04  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", findLucky(arr))
	fmt.Println()
}

func test05() {
	arr := []int{7, 7, 7, 7, 7, 7, 7}
	success := 7

	fmt.Println("test05  arr:=", arr)
	fmt.Println("success result:", success)
	fmt.Println("result:", findLucky(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
