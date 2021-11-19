package main

import (
	"fmt"
)

func reverseBits(num uint32) uint32 {
	result := uint32(0)

	for i := 31; i >= 0 && num > 0; i-- {
		result |= num & 1 << i
		num >>= 1
	}

	return result
}

func test01() {
	num := uint32(43261596)
	succResult := uint32(964176192)
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reverseBits(num))
	fmt.Println()
}

func test02() {
	num := uint32(4294967293)
	succResult := 3221225471
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reverseBits(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
