package main

import "fmt"

func fraction(cont []int) []int {
	index := len(cont) - 1
	n := 1
	m := cont[index]

	for index--; index >= 0; index-- {
		fmt.Println(n, m, cont[index])

		n = cont[index]*m + n
		n, m = m, n
	}

	return []int{m, n}
}

func test01() {
	cont := []int{3, 2, 0, 2}
	succResult := []int{13, 4}

	fmt.Println("test01 cont:=", cont)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fraction(cont))
	fmt.Println()
}

func test02() {
	cont := []int{0, 0, 3}
	succResult := []int{3, 1}

	fmt.Println("test02 cont:=", cont)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", fraction(cont))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
