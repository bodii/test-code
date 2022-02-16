package main

import (
	"fmt"
	"sort"
)

func minimumSum(num int) int {
	numList := make([]int, 4)
	for i := 0; num > 0; i++ {
		numList[i] = num % 10
		num /= 10
	}

	sort.Ints(numList)

	return (numList[0]+numList[1])*10 + numList[2] + numList[3]
}

func test01() {
	num := 2932
	succResult := 52
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumSum(num))
	fmt.Println()
}

func test02() {
	num := 4009
	succResult := 13
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumSum(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
