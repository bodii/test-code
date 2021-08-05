package main

import (
	"fmt"
)

func licenseKeyFormatting(s string, k int) string {
	list := []byte{}
	t := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != '-' {
			if t == k {
				list = append([]byte{'-'}, list...)
				t = 0
			}
			si := s[i]
			if s[i] >= 'a' && s[i] <= 'z' {
				si -= 32
			}
			list = append([]byte{si}, list...)
			t++
		}
	}

	return string(list)
}

func test01() {
	s := "5F3Z-2e-9-w"
	k := 4
	succResult := "5F3Z-2E9W"
	fmt.Println("test01 s:=", s, " k:=", 4)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", licenseKeyFormatting(s, k))
	fmt.Println()
}

func test02() {
	s := "2-5g-3-J"
	k := 2
	succResult := "2-5G-3J"
	fmt.Println("test02 s:=", s, " k:=", 4)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", licenseKeyFormatting(s, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
