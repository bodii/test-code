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

	five, ten := 0, 0
	for _, v := range bills {
		if v == 5 {
			five++
		} else if v == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
		} else if v == 20 {
			if ten > 0 && five > 0 {
				ten--
				five--
			} else if five > 2 {
				five -= 3
			} else {
				return false
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
