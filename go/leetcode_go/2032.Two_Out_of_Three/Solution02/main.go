package main

import (
	"fmt"
)

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	s1, s2, s3 := make([]bool, 101), make([]bool, 101), make([]bool, 101)
	for _, v := range nums1 {
		s1[v] = true
	}

	for _, v := range nums2 {
		s2[v] = true
	}

	for _, v := range nums3 {
		s3[v] = true
	}

	result := make([]int, 0)
	for k := range s1 {
		if (s1[k] && s2[k]) || (s1[k] && s3[k]) || (s2[k] && s3[k]) {
			result = append(result, k)
		}
	}

	return result
}

func test01() {
	nums1 := []int{1, 1, 3, 2}
	nums2 := []int{2, 3}
	nums3 := []int{3}
	succResult := []int{3, 2}
	fmt.Println("test01:")
	fmt.Println("nums1:=", nums1)
	fmt.Println("nums2:=", nums2)
	fmt.Println("nums3:=", nums3)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoOutOfThree(nums1, nums2, nums3))
	fmt.Println()
}

func test02() {
	nums1 := []int{3, 1}
	nums2 := []int{2, 3}
	nums3 := []int{1, 2}
	succResult := []int{2, 3, 1}
	fmt.Println("test02:")
	fmt.Println("nums1:=", nums1)
	fmt.Println("nums2:=", nums2)
	fmt.Println("nums3:=", nums3)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoOutOfThree(nums1, nums2, nums3))
	fmt.Println()
}

func test03() {
	nums1 := []int{1, 2, 2}
	nums2 := []int{4, 3, 3}
	nums3 := []int{5}
	succResult := []int{}
	fmt.Println("test03:")
	fmt.Println("nums1:=", nums1)
	fmt.Println("nums2:=", nums2)
	fmt.Println("nums3:=", nums3)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", twoOutOfThree(nums1, nums2, nums3))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
