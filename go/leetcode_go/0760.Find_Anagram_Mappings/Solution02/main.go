package main

import (
	"fmt"
)

func anagramMappings(nums1 []int, nums2 []int) []int {
	keys := make(map[int]int)
	l, r := 0, len(nums1)-1
	for l <= r {
		keys[nums2[l]] = l
		keys[nums2[r]] = r
		l++
		r--
	}

	l, r = 0, len(nums1)-1
	for l <= r {
		nums2[l] = keys[nums1[l]]
		nums2[r] = keys[nums1[r]]
		l++
		r--
	}

	return nums2
}

func test01() {
	nums1, nums2 := []int{12, 28, 46, 32, 50}, []int{50, 12, 32, 46, 28}
	succResult := []int{1, 4, 3, 2, 0}
	fmt.Println("test01 nums1:=", nums1, " nums2:=", nums2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", anagramMappings(nums1, nums2))
	fmt.Println()
}

func test02() {
	nums1, nums2 := []int{84, 46}, []int{84, 46}
	succResult := []int{0, 1}
	fmt.Println("test02 nums1:=", nums1, " nums2:=", nums2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", anagramMappings(nums1, nums2))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
