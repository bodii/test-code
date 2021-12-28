package main

import "fmt"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	}

	for i := num / 2; i >= 2; i++ {
		if i*i == num {
			return true
		}
	}

	return false
}

func test01() {
	num := 16
	succResult := true
	fmt.Println("test01:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isPerfectSquare(num))
	fmt.Println()
}

func test02() {
	num := 14
	succResult := false
	fmt.Println("test02:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isPerfectSquare(num))
	fmt.Println()
}

func test03() {
	num := 5
	succResult := false
	fmt.Println("test03:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isPerfectSquare(num))
	fmt.Println()
}

func test04() {
	num := 104976
	succResult := true
	fmt.Println("test04:")
	fmt.Println("num:=", num)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", isPerfectSquare(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
