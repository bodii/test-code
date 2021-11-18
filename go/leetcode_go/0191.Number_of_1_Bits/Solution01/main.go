package main

import (
	"fmt"
)

func hammingWeight(num uint32) int {
	count := 0

	for num > 0 {
		num &= num - 1
		count++
	}

	return count
}

func test01() {
	num := uint32(11)
	succResult := 3
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hammingWeight(num))
	fmt.Println()
}

func test02() {
	num := uint32(128)
	succResult := 1
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hammingWeight(num))
	fmt.Println()
}

func test03() {
	num := uint32(4294967293)
	succResult := 31
	fmt.Println("test03 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", hammingWeight(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
