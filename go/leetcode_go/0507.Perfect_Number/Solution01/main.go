package main

import "fmt"

// Time limit exceeded
func checkPerfectNumber(num int) bool {
	number := 0
	for i := 1; ; i++ {
		sum := (num / i)
		if i*sum == num {
			number += i
		}
		if sum <= 1 {
			return number == num
		}
	}
}

func test01() {
	num := 28
	fmt.Println("test01 result:", checkPerfectNumber(num))
}

func test02() {
	num := 6
	fmt.Println("test02 result:", checkPerfectNumber(num))
}

func test03() {
	num := 496
	fmt.Println("test03 result:", checkPerfectNumber(num))
}

func test04() {
	num := 8128
	fmt.Println("test04 result:", checkPerfectNumber(num))
}

func test05() {
	num := 2
	fmt.Println("test05 result:", checkPerfectNumber(num))
}

func test06() {
	num := 99999992
	fmt.Println("test06 result:", checkPerfectNumber(num))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
