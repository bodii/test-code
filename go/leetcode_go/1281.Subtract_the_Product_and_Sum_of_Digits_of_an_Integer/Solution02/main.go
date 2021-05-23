package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	product, sum := 1, 0
	for n > 0 {
		one := n % 10
		product *= one
		sum += one
		n /= 10
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
