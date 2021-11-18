package main

import (
	"fmt"
	"sort"
)

func sortByBits(arr []int) []int {
	bits := make([]int, 10001)
	for _, a := range arr {
		if bits[a] == 0 {
			bits[a] = countBit1Num(a)
		}
	}

	sort.SliceStable(arr, func(i, j int) bool {
		if bits[arr[i]] == bits[arr[j]] {
			return arr[i] < arr[j]
		}

		return bits[arr[i]] < bits[arr[j]]
	})

	return arr
}

func countBit1Num(a int) int {
	count := 0

	for a > 0 {
		a &= a - 1
		count++
	}

	return count
}

func test01() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	succResult := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortByBits(arr))
	fmt.Println()
}

func test02() {
	arr := []int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1}
	succResult := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	fmt.Println("test02 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortByBits(arr))
	fmt.Println()
}

func test03() {
	arr := []int{10000, 10000}
	succResult := []int{10000, 10000}
	fmt.Println("test03 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortByBits(arr))
	fmt.Println()
}

func test04() {
	arr := []int{2, 3, 5, 7, 11, 13, 17, 19}
	succResult := []int{2, 3, 5, 17, 7, 11, 13, 19}
	fmt.Println("test04 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortByBits(arr))
	fmt.Println()
}

func test05() {
	arr := []int{10, 100, 1000, 10000}
	succResult := []int{10, 100, 10000, 1000}
	fmt.Println("test05 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", sortByBits(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
