package main

import (
	"fmt"
)

func addToArrayForm(num []int, k int) []int {
	remainder := 0
	numLen := len(num)
	for i := numLen - 1; i >= 0; i-- {
		remainder += num[i]
		if k > 0 {
			remainder += k % 10
			k /= 10
		}
		num[i] = remainder % 10

		if remainder >= 10 {
			remainder = 1
		} else {
			remainder = 0
		}
	}

	remainder += k
	for remainder > 0 {
		num = append([]int{remainder % 10}, num...)
		remainder /= 10
	}

	return num
}

func test01() {
	k := 34
	num := []int{1, 2, 0, 0}
	succResult := []int{1, 2, 3, 4}
	fmt.Println("test01 k:=", k, " num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addToArrayForm(num, k))
	fmt.Println()
}

func test02() {
	k := 181
	num := []int{2, 7, 4}
	succResult := []int{4, 5, 5}
	fmt.Println("test02 k:=", k, " num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addToArrayForm(num, k))
	fmt.Println()
}

func test03() {
	k := 806
	num := []int{2, 1, 5}
	succResult := []int{1, 0, 2, 1}
	fmt.Println("test03 k:=", k, " num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addToArrayForm(num, k))
	fmt.Println()
}

func test04() {
	k := 1
	num := []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}
	succResult := []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	fmt.Println("test04 k:=", k, " num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addToArrayForm(num, k))
	fmt.Println()
}

func test05() {
	k := 23
	num := []int{0}
	succResult := []int{2, 3}
	fmt.Println("test05 k:=", k, " num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addToArrayForm(num, k))
	fmt.Println()
}

func test06() {
	k := 809
	num := []int{6}
	succResult := []int{8, 1, 5}
	fmt.Println("test06 k:=", k, " num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", addToArrayForm(num, k))
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
