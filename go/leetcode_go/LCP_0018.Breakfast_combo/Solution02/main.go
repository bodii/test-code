package main

import (
	"fmt"
	"sort"
)

func breakfastNumber(staple []int, drinks []int, x int) int {
	nums := 0
	max := 1000000007

	sort.Ints(drinks)

	for _, v := range staple {
		if x-v <= 0 {
			continue
		}
		c := sort.SearchInts(drinks, x-v+1)
		nums += c
	}

	return nums % max
}

func test01() {
	staple, drinks := []int{10, 20, 5}, []int{5, 5, 2}
	x := 15
	succResult := 6
	fmt.Println("test01 staple:=", staple, " drinks:=", drinks, " x:=", x)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", breakfastNumber(staple, drinks, x))
	fmt.Println()
}

func test02() {
	staple, drinks := []int{2, 1, 1}, []int{8, 9, 5, 1}
	x := 9
	succResult := 8
	fmt.Println("test02 staple:=", staple, " drinks:=", drinks, " x:=", x)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", breakfastNumber(staple, drinks, x))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
