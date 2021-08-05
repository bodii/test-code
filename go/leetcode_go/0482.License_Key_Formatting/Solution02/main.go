package main

import (
	"fmt"
	"strings"
)

func licenseKeyFormatting(s string, k int) string {
	list := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '-' {
			continue
		}

		si := s[i]
		if si >= 'a' && si <= 'z' {
			si -= 'a' - 'A'
		}
		list = append(list, si)
	}

	listLen := len(list)
	first := listLen % k

	result := strings.Builder{}
	if first != 0 {
		result.Write(list[:first])
	}

	for i := first; i < listLen; i += k {
		if i != 0 {
			result.WriteByte('-')
		}
		result.Write(list[i : i+k])
	}

	return result.String()
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
