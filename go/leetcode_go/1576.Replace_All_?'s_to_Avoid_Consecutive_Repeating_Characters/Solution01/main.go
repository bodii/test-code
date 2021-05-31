package main

import (
	"fmt"
)

func modifyString(s string) string {
	sArr := []byte(s)
	sLen := len(s)

	for i := 0; i < sLen; i++ {
		if s[i] == '?' {
			left, right := byte('-'), byte('-')
			if i != 0 {
				left = sArr[i-1]
			}

			if i != sLen-1 {
				right = sArr[i+1]
			}

			a := byte('a')
			for a == left || a == right {
				a++
			}
			sArr[i] = a
		}
	}

	return string(sArr)
}

func test01() {
	s := "?zs"
	successResult := "azs"
	fmt.Println("test01 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", modifyString(s))
	fmt.Println()
}

func test02() {
	s := "ubv?w"
	successResult := "ubvaw"
	fmt.Println("test02 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", modifyString(s))
	fmt.Println()
}

func test03() {
	s := "j?qg??b"
	successResult := "jaqgacb"
	fmt.Println("test03 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", modifyString(s))
	fmt.Println()
}

func test04() {
	s := "??yw?ipkj?"
	successResult := "acywaipkja"
	fmt.Println("test04 s:", s)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", modifyString(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
