package main

import (
	"fmt"
)

func twoOutOfThree(nums1 []int, nums2 []int, nums3 []int) []int {
	m1, m2, m3 := getSliceToMap(nums1), getSliceToMap(nums2), getSliceToMap(nums3)

	s1 := make([]int, 101)
	for k := range m1 {
		s1[k]++
	}

	for k := range m2 {
		s1[k]++
	}

	for k := range m3 {
		s1[k]++
	}

	result := make([]int, 0)
	for k, v := range s1 {
		if v > 1 {
			result = append(result, k)
		}
	}

	return result
}

func getSliceToMap(nums []int) map[int]bool {
	m := make(map[int]bool)
	for _, n := range nums {
		m[n] = true
	}

	return m
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
