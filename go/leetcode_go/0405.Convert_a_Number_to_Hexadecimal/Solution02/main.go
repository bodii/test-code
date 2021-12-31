package main

import (
	"fmt"
)

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	if num < 0 {
		num += 4294967296
	}
	result := []byte{}
	hash := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	for num != 0 {
		result = append([]byte{hash[num%16]}, result...)
		num /= 16
	}
	return string(result)
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
