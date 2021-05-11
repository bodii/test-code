package main

import "fmt"

func minCount(coins []int) int {
	nums := 0
	for _, v := range coins {
		if v%2 == 0 {
			nums += v / 2
		} else {
			nums += v/2 + 1
		}
	}

	return nums
}

func test01() {
	coins := []int{4, 2, 1}
	succResult := 4

	fmt.Println("test01 coins:=", coins)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minCount(coins))
	fmt.Println()
}

func test02() {
	coins := []int{2, 3, 10}
	succResult := 8

	fmt.Println("test02 coins:=", coins)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minCount(coins))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
