package main

import (
	"fmt"
	"strconv"
)

func getLucky(s string, k int) int {
	intStr := strToIntStr(s)

	for i := 0; i < k; i++ {
		result := 0

		for _, v := range intStr {
			result += int(v - '1' + 1)
		}

		if i+1 == k {
			return result
		}
		intStr = strconv.Itoa(result)
	}

	return 0
}

func strToIntStr(s string) string {
	sBytes := []byte(s)

	res := ""
	for _, v := range sBytes {
		res += strconv.Itoa(int(v - 'a' + 1))
	}

	return res
}

func test01() {
	s := "iiii"
	k := 1
	succResult := 36
	fmt.Println("test01 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getLucky(s, k))
	fmt.Println()
}

func test02() {
	s := "leetcode"
	k := 2
	succResult := 6
	fmt.Println("test02 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getLucky(s, k))
	fmt.Println()
}

func test03() {
	s := "zbax"
	k := 2
	succResult := 8
	fmt.Println("test03 s:=", s, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getLucky(s, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
