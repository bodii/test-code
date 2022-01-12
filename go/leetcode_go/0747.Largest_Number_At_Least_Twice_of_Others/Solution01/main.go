package main

import (
	"fmt"
)

func dominantIndex(nums []int) int {
	first, secord := 0, 0
	firstInd := -1

	for k, v := range nums {
		first, secord = findTwoLargestNums(first, secord, v)
		if first == v {
			firstInd = k
		}
	}

	if first == 0 || first/2 < secord {
		return -1
	}

	return firstInd

}

func findTwoLargestNums(a, b, c int) (int, int) {
	if a > b && a > c {
		if b > c {
			return a, b
		} else {
			return a, c
		}
	} else if b > a && b > c {
		if a > c {
			return b, a
		} else {
			return b, c
		}
	}

	if a > b {
		return c, a
	}

	return c, b
}

func test01() {
	nums := []int{3, 6, 1, 0}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", dominantIndex(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 2, 3, 4}
	succResult := -1
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", dominantIndex(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1}
	succResult := 0
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", dominantIndex(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
