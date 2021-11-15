package main

import (
	"fmt"
)

func numTeams(rating []int) int {
	nums := 0

	ratingLen := len(rating)
	for i := 1; i < ratingLen-1; i++ {
		lMin, lMax := 0, 0
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				nums += lMin
				lMin++
			} else {
				nums += lMax
				lMax++
			}
		}
	}

	return nums
}

func test01() {
	rating := []int{2, 5, 3, 4, 1}
	succResult := 3
	fmt.Println("test01 rating:=", rating)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numTeams(rating))
	fmt.Println()
}

func test02() {
	rating := []int{2, 1, 3}
	succResult := 0
	fmt.Println("test02 rating:=", rating)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numTeams(rating))
	fmt.Println()
}

func test03() {
	rating := []int{1, 2, 3, 4}
	succResult := 4
	fmt.Println("test03 rating:=", rating)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", numTeams(rating))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
