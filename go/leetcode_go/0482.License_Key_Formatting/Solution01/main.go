package main

import (
	"fmt"
)

func licenseKeyFormatting(s string, k int) string {
	list := []rune{}
	for _, v := range s {
		if v != '-' {
			if v >= 'a' && v <= 'z' {
				v += 'A' - 'a'
			}
			list = append(list, v)
		}
	}

	listLen := len(list)
	first := listLen % k

	result := ""
	if first != 0 {
		result = string(list[:first])
	}

	for i := first; i < listLen; i += k {
		if i != 0 {
			result += "-"
		}
		result += string(list[i : i+k])
	}

	return result
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
