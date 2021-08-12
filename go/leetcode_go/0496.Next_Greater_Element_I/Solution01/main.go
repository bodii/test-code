package main

import (
	"fmt"
)

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nums2Len := len(nums2)

	for k, v := range nums1 {
		j := 0
		// 找到相等的索引位置
		for ; j < nums2Len; j++ {
			if nums2[j] == v {
				break
			}
		}
		// 下一个索引位置
		j++
		// 找到大于当前数的
		for ; j < nums2Len; j++ {
			if nums2[j] > v {
				nums1[k] = nums2[j]
				break
			}
		}

		// 如果j=nums2Len 说明没有找到
		if j == nums2Len {
			nums1[k] = -1
		}
	}

	return nums1
}

func test01() {
	nums1, nums2 := []int{4, 1, 2}, []int{1, 3, 4, 2}
	succResult := []int{-1, 3, -1}
	fmt.Println("test01 nums1:=", nums1, " nums2:=", nums2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", nextGreaterElement(nums1, nums2))
	fmt.Println()
}

func test02() {
	nums1, nums2 := []int{2, 4}, []int{1, 2, 3, 4}
	succResult := []int{3, -1}
	fmt.Println("test02 nums1:=", nums1, " nums2:=", nums2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", nextGreaterElement(nums1, nums2))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
