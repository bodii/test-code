package main

import (
	"fmt"
)

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	ind := -1
	for i := 0; i < len(s1); i++ {
		if ind == -1 && s1[i] != s2[i] {
			ind = i
		} else if ind > -1 && s1[i] != s2[i] {
			if s1[i] == s2[ind] && s1[ind] == s2[i] {
				ind = -2
			} else {
				return false
			}
		} else if ind < -1 && s1[i] != s2[i] {
			return false
		}
	}

	return ind == -2
}

func test01() {
	s1, s2 := "bank", "kanb"
	successResult := true
	fmt.Println("test01 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test02() {
	s1, s2 := "attack", "defend"
	successResult := false
	fmt.Println("test02 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test03() {
	s1, s2 := "kelb", "kelb"
	successResult := true
	fmt.Println("test03 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test04() {
	s1, s2 := "abcd", "dcba"
	successResult := false
	fmt.Println("test04 s1:=", s1, " s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test05() {
	s1, s2 := "ysmpagrkzsmmzmsssutzgpxrmoylkgemgfcperptsxjcsgojwourhxlhqkxumonfgrczmjvbhwvhpnocz",
		"ysmpagrqzsmmzmsssutzgpxrmoylkgemgfcperptsxjcsgojwourhxlhkkxumonfgrczmjvbhwvhpnocz"

	successResult := true
	fmt.Println("test05:")
	fmt.Println("s1:=", s1)
	fmt.Println("s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test06() {
	s1, s2 := "ab", "ac"

	successResult := false
	fmt.Println("test06:")
	fmt.Println("s1:=", s1)
	fmt.Println("s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func test07() {
	s1, s2 := "qgqeg", "gqgeq"

	successResult := false
	fmt.Println("test07:")
	fmt.Println("s1:=", s1)
	fmt.Println("s2:=", s2)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", areAlmostEqual(s1, s2))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
}
