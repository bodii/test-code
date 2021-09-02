package main

import (
	"fmt"
)

func distributeCandies(candies int, num_people int) []int {
	peoples := make([]int, num_people)

	for i := 1; candies > 0; i++ {
		if candies >= i {
			peoples[(i-1)%num_people] += i
		} else {
			peoples[(i-1)%num_people] += candies
		}
		candies -= i
	}

	return peoples
}

func test01() {
	candies, num_people := 7, 4
	succResult := []int{1, 2, 3, 1}
	fmt.Println("test01 candies:=", candies, " num_people:=", num_people)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", distributeCandies(candies, num_people))
	fmt.Println()
}

func test02() {
	candies, num_people := 10, 3
	succResult := []int{5, 2, 3}
	fmt.Println("test01 candies:=", candies, " num_people:=", num_people)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", distributeCandies(candies, num_people))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
