package main

import (
	"fmt"
)

func exchangeBits(num int) int {
	return (num&0x55555555)<<1 | (num&0xaaaaaaaa)>>1
}

func test01() {
	num := 2
	succResult := 1
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", exchangeBits(num))
	fmt.Println()
}

func test02() {
	num := 3
	succResult := 3
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", exchangeBits(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
