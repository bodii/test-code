package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	for k, v := range numbers {
		l, r := k+1, len(numbers)-1
		for l <= r {
			mid := (r-l)/2 + l
			if v+numbers[mid] == target {
				return []int{k + 1, mid + 1}
			} else if v+numbers[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return []int{-1, -1}
}

func test01() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	succResult := []int{1, 2}
	fmt.Println("test01 numbers:=", numbers, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(numbers, target))
	fmt.Println()
}

func test02() {
	numbers := []int{2, 3, 4}
	target := 6
	succResult := []int{1, 3}
	fmt.Println("test02 numbers:=", numbers, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(numbers, target))
	fmt.Println()
}

func test03() {
	numbers := []int{-1, 0}
	target := -1
	succResult := []int{1, 2}
	fmt.Println("test03 numbers:=", numbers, " target:=", target)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoSum(numbers, target))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
