package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums1 = append(nums1, nums2[:n]...)
	sort.Ints(nums1)
}

func test01() {
	nums1, m := []int{1, 2, 3, 0, 0, 0}, 3
	nums2, n := []int{2, 5, 6}, 3
	successResult := []int{1, 2, 2, 3, 5, 6}
	fmt.Println("test01 nums1:", nums1, " m:", m)
	fmt.Println("nums2:", nums2, " n:", n)
	fmt.Println("success result:", successResult)
	merge(nums1, m, nums2, n)
	fmt.Println("result: ", nums1)
	fmt.Println()
}

func test02() {
	nums1, m := []int{1}, 1
	nums2, n := []int{}, 0
	successResult := []int{1}
	fmt.Println("test02 nums1:", nums1, " m:", m)
	fmt.Println("nums2:", nums2, " n:", n)
	fmt.Println("success result:", successResult)
	merge(nums1, m, nums2, n)
	fmt.Println("result: ", nums1)
	fmt.Println()
}

func test03() {
	nums1, m := []int{0}, 0
	nums2, n := []int{1}, 1
	successResult := []int{1}
	fmt.Println("test03 nums1:", nums1, " m:", m)
	fmt.Println("nums2:", nums2, " n:", n)
	fmt.Println("success result:", successResult)
	merge(nums1, m, nums2, n)
	fmt.Println("result: ", nums1)
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
