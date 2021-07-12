package main

import (
	"fmt"
)

func balancedStringSplit(s string) int {
	count, sum := 0, 0
	for _, v := range s {
		if v == 'R' {
			sum++
		} else {
			sum--
		}

		if sum == 0 {
			count++
		}
	}

	return count
}

func test01() {
	s := "RLRRLLRLRL"
	succResult := 4
	fmt.Println("test01 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", balancedStringSplit(s))
	fmt.Println()
}

func test02() {
	s := "RLLLLRRRLR"
	succResult := 3
	fmt.Println("test02 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", balancedStringSplit(s))
	fmt.Println()
}

func test03() {
	s := "LLLLRRRR"
	succResult := 1
	fmt.Println("test03 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", balancedStringSplit(s))
	fmt.Println()
}

func test04() {
	s := "RLRRRLLRLL"
	succResult := 2
	fmt.Println("test04 s:=", s)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", balancedStringSplit(s))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
