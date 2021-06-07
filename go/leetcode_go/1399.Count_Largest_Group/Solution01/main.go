package main

import (
	"fmt"
)

func countLargestGroup(n int) int {

	numsMap := make(map[int]int)
	for i := 1; i <= n; i++ {
		num, n := 0, i
		for n > 0 {
			num += n % 10
			n /= 10
		}

		numsMap[num]++
	}

	max, sum := 0, 0
	for _, v := range numsMap {
		if v > max {
			max = v
			sum = 1
		} else if v == max {
			sum++
		}
	}

	return sum
}

func test01() {
	n := 13
	success := 4

	fmt.Println("test01  n:=", n)
	fmt.Println("success result:", success)
	fmt.Println("result:", countLargestGroup(n))
	fmt.Println()
}

func test02() {
	n := 2
	success := 2

	fmt.Println("test02  n:=", n)
	fmt.Println("success result:", success)
	fmt.Println("result:", countLargestGroup(n))
	fmt.Println()
}

func test03() {
	n := 15
	success := 6

	fmt.Println("test03  n:=", n)
	fmt.Println("success result:", success)
	fmt.Println("result:", countLargestGroup(n))
	fmt.Println()
}

func test04() {
	n := 24
	success := 5

	fmt.Println("test04  n:=", n)
	fmt.Println("success result:", success)
	fmt.Println("result:", countLargestGroup(n))
	fmt.Println()
}

func test05() {
	n := 46
	success := 6

	fmt.Println("test05  n:=", n)
	fmt.Println("success result:", success)
	fmt.Println("result:", countLargestGroup(n))
	fmt.Println()
}

func test06() {
	n := 18
	success := 9

	fmt.Println("test06  n:=", n)
	fmt.Println("success result:", success)
	fmt.Println("result:", countLargestGroup(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
