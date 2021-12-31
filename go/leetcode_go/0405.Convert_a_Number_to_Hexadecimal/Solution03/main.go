package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	helper := "0123456789abcdef"

	result := ""
	for num != 0 && len(result) < 8 {
		result = string(helper[num&15]) + result
		num >>= 4
	}
	return result
}

func test01() {
	num := 26
	successresultult := "1a"
	fmt.Println("test01 num:", num)
	fmt.Println("success resultult:", successresultult)
	fmt.Println("resultult: ", toHex(num))
	fmt.Println()
}

func test02() {
	num := -1
	successresultult := "ffffffff"
	fmt.Println("test02 num:", num)
	fmt.Println("success resultult:", successresultult)
	fmt.Println("resultult: ", toHex(num))
	fmt.Println()
}

func test03() {
	num := 684
	successresultult := "2ac"
	fmt.Println("test03 num:", num)
	fmt.Println("success resultult:", successresultult)
	fmt.Println("resultult: ", toHex(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
