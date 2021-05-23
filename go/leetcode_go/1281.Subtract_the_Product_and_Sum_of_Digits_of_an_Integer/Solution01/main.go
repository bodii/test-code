package main

import (
	"fmt"
	"strconv"
)

func subtractProductAndSum(n int) int {
	nStr := strconv.Itoa(n)
	byteToInt := map[rune]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6,
		'7': 7, '8': 8, '9': 9}
	product, sum := 1, 0

	for _, v := range nStr {
		if vInt, ok := byteToInt[v]; ok {
			product *= vInt
			sum += vInt
		}
	}

	return product - sum
}

func test01() {
	n := 234
	succResult := 15
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", subtractProductAndSum(n))
	fmt.Println()
}

func test02() {
	n := 4421
	succResult := 21
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", subtractProductAndSum(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
