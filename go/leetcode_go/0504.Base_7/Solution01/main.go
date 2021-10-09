package main

import (
	"fmt"
)

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	res := []byte{}
	isNegative := num < 0
	if isNegative {
		num = -num
	}

	for num > 0 {
		res = append([]byte{byte((num % 7) + '0')}, res...)
		num /= 7
	}

	if isNegative {
		res = append([]byte{'-'}, res...)
	}

	return string(res)
}

func test01() {
	num := 100
	successResult := "202"
	fmt.Println("test01 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", convertToBase7(num))
	fmt.Println()
}

func test02() {
	num := -7
	successResult := "-10"
	fmt.Println("test02 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", convertToBase7(num))
	fmt.Println()
}

func test03() {
	num := 0
	successResult := "0"
	fmt.Println("test03 num:", num)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", convertToBase7(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
