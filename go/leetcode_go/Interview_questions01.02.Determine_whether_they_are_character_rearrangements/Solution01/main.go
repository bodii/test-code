package main

import (
	"fmt"
	"sort"
)

func CheckPermutation(s1 string, s2 string) bool {
	s1Slice, s2Slice := []byte(s1), []byte(s2)
	sort.SliceStable(s1Slice, func(i, j int) bool {
		return s1Slice[i] < s1Slice[j]
	})

	sort.SliceStable(s2Slice, func(i, j int) bool {
		return s2Slice[i] < s2Slice[j]
	})

	return string(s1Slice) == string(s2Slice)
}

func test01() {
	s1, s2 := "abc", "bca"
	successResult := true
	fmt.Println("test01 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", CheckPermutation(s1, s2))
	fmt.Println()
}

func test02() {
	s1, s2 := "abc", "bad"
	successResult := false
	fmt.Println("test02 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", CheckPermutation(s1, s2))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
