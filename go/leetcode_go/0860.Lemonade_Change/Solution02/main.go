package main

import "fmt"

func lemonadeChange(bills []int) bool {
	len := len(bills)
	if len < 1 {
		return true
	}
	if bills[0] != 5 {
		return false
	}
	five, ten, twenty := 0, 0, 0

	for _, v := range bills {
		if v == 5 {
			five++
		} else {
			if five < 1 || (v == 20 && ten < 1 && five < 3) {
				return false
			}

			if v == 10 {
				five--
				ten++
			} else if ten > 0 {
				five--
				ten--
				twenty++
			} else {
				five -= 3
				twenty++
			}
		}
	}

	return true
}

func test01() {
	bills := []int{5, 5, 5, 10, 20}
	success := true

	fmt.Println("test01 bills:=", bills)
	fmt.Println("success:=", success)
	fmt.Println("result:=", lemonadeChange(bills))
	fmt.Println()
}

func test02() {
	bills := []int{5, 5, 10}
	success := true

	fmt.Println("test02 bills:=", bills)
	fmt.Println("success:=", success)
	fmt.Println("result:=", lemonadeChange(bills))
	fmt.Println()
}

func test03() {
	bills := []int{10, 10}
	success := false

	fmt.Println("test03 bills:=", bills)
	fmt.Println("success:=", success)
	fmt.Println("result:=", lemonadeChange(bills))
	fmt.Println()
}

func test04() {
	bills := []int{5, 5, 10, 10, 20}
	success := false

	fmt.Println("test04 bills:=", bills)
	fmt.Println("success:=", success)
	fmt.Println("result:=", lemonadeChange(bills))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
