package main

import "fmt"

func isSameAfterReversals(num int) bool {
	if num == 0 {
		return true
	}

	return num%10 != 0
}

func test01() {
	num := 526
	succResult := true
	fmt.Println("test01:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isSameAfterReversals(num))
	fmt.Println()
}

func test02() {
	num := 1800
	succResult := false
	fmt.Println("test02:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isSameAfterReversals(num))
	fmt.Println()
}

func test03() {
	num := 0
	succResult := true
	fmt.Println("test03:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isSameAfterReversals(num))
	fmt.Println()
}

func test04() {
	num := 10
	succResult := false
	fmt.Println("test04:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isSameAfterReversals(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
