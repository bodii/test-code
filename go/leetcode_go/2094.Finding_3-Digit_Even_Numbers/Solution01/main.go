package main

import (
	"fmt"
	"sort"
)

func findEvenNumbers(digits []int) []int {
	sort.Ints(digits)

	digitMaps := make(map[int]bool)
	result := make([]int, 0)

	for k1, v1 := range digits {
		if v1 == 0 {
			continue
		}

		for k2, v2 := range digits {
			if k1 == k2 {
				continue
			}
			for k3, v3 := range digits {
				if k1 == k3 || k2 == k3 || v3&1 == 1 {
					continue
				}
				val := v1*100 + v2*10 + v3
				if !digitMaps[val] {
					result = append(result, val)
					digitMaps[val] = true
				}
			}
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
