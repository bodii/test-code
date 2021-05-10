package main

import "fmt"

func lemonadeChange(bills []int) bool {
	len := len(bills)
	if len < 1 {
		return true
	}
	income := map[int]int{5: 0, 10: 0, 20: 0}
	if bills[0] != 5 {
		return false
	}
	income[5]++

	for i := 1; i < len; i++ {
		if bills[i] != 5 {
			if income[5] < 1 || (bills[i] == 20 && income[5] < 3 && income[10] < 1) {
				return false
			}
			if bills[i] == 10 {
				income[5]--
				income[10]++
			}
			if bills[i] == 20 {
				if income[10] > 0 {
					income[5]--
					income[10]--
				} else {
					income[5] -= 3
				}

				income[20]++
			}
		} else {
			income[5]++
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
