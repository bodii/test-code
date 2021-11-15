package main

import (
	"fmt"
)

func numTeams(rating []int) int {
	last := len(rating) - 1

	nums := 0
	for i := 0; i <= last-2; i++ {
		for j := i + 1; j <= last-1; j++ {
			for k := j + 1; k <= last; k++ {
				if (rating[i] < rating[j] && rating[j] < rating[k]) ||
					(rating[i] > rating[j] && rating[j] > rating[k]) {
					nums++
				}
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
