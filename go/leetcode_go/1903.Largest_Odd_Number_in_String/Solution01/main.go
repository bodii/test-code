package main

import (
	"fmt"
)

func largestOddNumber(num string) string {
	for i := len(num) - 1; i >= 0; i-- {
		if (num[i]-'0')%2 != 0 {
			return num[0 : i+1]
		}
	}

	return ""
}

func test01() {
	num := "52"
	succResult := "5"
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestOddNumber(num))
	fmt.Println()
}

func test02() {
	num := "4206"
	succResult := ""
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestOddNumber(num))
	fmt.Println()
}

func test03() {
	num := "35427"
	succResult := "35427"
	fmt.Println("test03 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestOddNumber(num))
	fmt.Println()
}

func test04() {
	num := "239537672423884969653287101"
	succResult := "239537672423884969653287101"
	fmt.Println("test04 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", largestOddNumber(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
