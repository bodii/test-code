package main

import (
	"fmt"
)

func findContinuousSequence(target int) [][]int {
	result := [][]int{}
	l, r := 1, 2
	for l < r {
		sum := (l + r) * (r - l + 1) / 2
		if sum == target {
			var list []int
			for i := l; i <= r; i++ {
				list = append(list, i)
			}
			result = append(result, list)
			l++
		} else if sum < target {
			r++
		} else {
			l++
		}
	}
	return result
}

func test01() {
	target := 9
	succResult := [][]int{{2, 3, 4}, {4, 5}}
	fmt.Println("test01 target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findContinuousSequence(target))
	fmt.Println()
}

func test02() {
	target := 15
	succResult := [][]int{{1, 2, 3, 4, 5}, {4, 5, 6}, {7, 8}}
	fmt.Println("test02 target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findContinuousSequence(target))
	fmt.Println()
}

func test03() {
	target := 100
	succResult := [][]int{{9, 10, 11, 12, 13, 14, 15, 16}, {18, 19, 20, 21, 22}}
	fmt.Println("test03 target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findContinuousSequence(target))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
