package main

import (
	"fmt"
)

const (
	m0 = 0x5555555555555555
	m1 = 0x3333333333333333
	m2 = 0x0f0f0f0f0f0f0f0f
	m3 = 0x00ff00ff00ff00ff
	m4 = 0x0000ffff0000ffff
)

func reverseBits(num uint32) uint32 {
	const m = 1<<32 - 1
	num = num>>1&(m0&m) | num&(m0&m)<<1
	num = num>>2&(m1&m) | num&(m1&m)<<2
	num = num>>4&(m2&m) | num&(m2&m)<<4
	num = num>>8&(m3&m) | num&(m3&m)<<8
	return num>>16 | num<<16
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
