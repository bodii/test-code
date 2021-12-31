package main

import (
	"fmt"
	"math"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	result := make([]byte, 0)
	if num < 0 {
		num = math.MaxUint32 + num + 1
	}

	for num > 0 {
		r := num % 16
		if r < 10 {
			result = append([]byte{byte('0' + r)}, result...)
		} else {
			result = append([]byte{byte('a' + (r - 10))}, result...)
		}

		num /= 16
	}

	return string(result)
}

func test01() {
	num := 26
	successResult := "1a"
	fmt.Println("test01 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", toHex(num))
	fmt.Println()
}

func test02() {
	num := -1
	successResult := "ffffffff"
	fmt.Println("test02 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", toHex(num))
	fmt.Println()
}

func test03() {
	num := 684
	successResult := "2ac"
	fmt.Println("test03 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", toHex(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
