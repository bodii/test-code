package main

import (
	"fmt"
)

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}
	sLen := len(s1)
	s2Bytes := []byte(s2)

	swap := false
	for i := 0; i < sLen; i++ {
		if s1[i] != s2Bytes[i] && !swap {
			for j := i + 1; j < sLen; j++ {
				if s1[i] == s2Bytes[j] && s1[j] != s2Bytes[j] {
					s2Bytes[i], s2Bytes[j] = s2Bytes[j], s2Bytes[i]
					swap = true
				}
			}
			if !swap {
				return false
			}
		} else if s1[i] != s2Bytes[i] && swap {
			return false
		}
	}

	return true
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

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
