package main

import (
	"fmt"
	"sort"
	"strings"
)

func CheckPermutation(s1 string, s2 string) bool {
	s1s, s2s := strings.Split(s1, ""), strings.Split(s2, "")
	sort.Strings(s1s)
	sort.Strings(s2s)

	return strings.Join(s1s, "") == strings.Join(s2s, "")
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
