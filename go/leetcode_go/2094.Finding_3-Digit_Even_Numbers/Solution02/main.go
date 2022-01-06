package main

import (
	"fmt"
)

func findEvenNumbers(digits []int) []int {
	digitList := make([]int, 10)
	for _, v := range digits {
		digitList[v]++
	}

	result := make([]int, 0)
	for i := 100; i < 1000; i += 2 {
		flag := false
		lines := make([]int, 10)
		for j := i; j > 0; j /= 10 {
			digit := j % 10
			if digitList[digit] == lines[digit] {
				flag = true
				break
			}
			lines[digit]++
		}

		if !flag {
			result = append(result, i)
		}
	}

	return result
}

func test01() {
	digits := []int{2, 1, 3, 0}
	succResult := []int{102, 120, 130, 132, 210, 230, 302, 310, 312, 320}
	fmt.Println("test01 digits:=", digits)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findEvenNumbers(digits))
	fmt.Println()
}

func test02() {
	digits := []int{2, 2, 8, 8, 2}
	succResult := []int{222, 228, 282, 288, 822, 828, 882}
	fmt.Println("test02 digits:=", digits)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findEvenNumbers(digits))
	fmt.Println()
}

func test03() {
	digits := []int{3, 7, 5}
	succResult := []int{}
	fmt.Println("test03 digits:=", digits)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findEvenNumbers(digits))
	fmt.Println()
}

func test04() {
	arr := []int{0, 2, 0, 0}
	succResult := []int{200}
	fmt.Println("test04 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findEvenNumbers(arr))
	fmt.Println()
}

func test05() {
	arr := []int{0, 0, 0}
	succResult := []int{}
	fmt.Println("test05 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findEvenNumbers(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
