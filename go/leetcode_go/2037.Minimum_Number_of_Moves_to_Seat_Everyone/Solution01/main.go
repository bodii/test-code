package main

import (
	"fmt"
	"sort"
)

func minMovesToSeat(seats []int, students []int) int {
	sort.Ints(seats)
	sort.Ints(students)

	result := 0
	for k, v := range students {
		result += abs(v - seats[k])
	}

	return result
}

func abs(i int) int {
	if i >= 0 {
		return i
	}

	return -i
}

func test01() {
	seats, students := []int{3, 1, 5}, []int{2, 7, 4}
	succResult := 4
	fmt.Println("test01 seats:=", seats, " students:=", students)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMovesToSeat(seats, students))
	fmt.Println()
}

func test02() {
	seats, students := []int{4, 1, 5, 9}, []int{1, 3, 2, 6}
	succResult := 7
	fmt.Println("test02 seats:=", seats, " students:=", students)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMovesToSeat(seats, students))
	fmt.Println()
}

func test03() {
	seats, students := []int{2, 2, 6, 6}, []int{1, 3, 2, 6}
	succResult := 4
	fmt.Println("test03 seats:=", seats, " students:=", students)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minMovesToSeat(seats, students))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
