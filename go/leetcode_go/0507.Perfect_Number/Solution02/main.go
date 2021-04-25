package main

import "fmt"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}

	c := 1
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			c += i + num/i
		}
	}

	return c == num
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
