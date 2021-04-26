package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	n := make([]int, 0)
	sort.Ints(nums) // 排序，从小到大
	n = append(n, nums[0])
	// 去重
	for j := 1; j < len(nums); j++ {
		if nums[j-1] != nums[j] {
			n = append(n, nums[j])
		}
	}
	ln := len(n)
	if ln <= 2 { // 如果不存在第三大的数，则返回最大数
		return n[ln-1]
	}
	return n[ln-3]
}

func test01() {
	nums := []int{3, 2, 1}
	fmt.Println("test01 result:", thirdMax(nums))
}

func test02() {
	nums := []int{1, 2}
	fmt.Println("test02 result:", thirdMax(nums))
}

func test03() {
	nums := []int{2, 2, 3, 1}
	fmt.Println("test03 result:", thirdMax(nums))
}

func test04() {
	nums := []int{1, 2, -2147483648}
	fmt.Println("test04 result:", thirdMax(nums))
}

func test05() {
	nums := []int{-2147483648, 1, 1}
	fmt.Println("test05 result:", thirdMax(nums))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
