package main

import (
	"fmt"
)

func numberOfSteps(num int) int {
	count := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num--
		}
		count++
	}

	return count
}

func test01() {
	num := 14
	succResult := 6
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numberOfSteps(num))
	fmt.Println()
}

func test02() {
	num := 8
	succResult := 4
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numberOfSteps(num))
	fmt.Println()
}

func test03() {
	num := 123
	succResult := 12
	fmt.Println("test03 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numberOfSteps(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
