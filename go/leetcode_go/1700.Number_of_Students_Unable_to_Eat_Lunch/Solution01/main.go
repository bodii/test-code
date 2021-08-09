package main

import (
	"fmt"
)

func countStudents(students []int, sandwiches []int) int {
	for len(students) > 0 {
		fmt.Println(students, sandwiches)

		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			if !hasNext(students, sandwiches[0]) {
				break
			}
			students = append(students[1:], students[0])
		}
	}

	return len(students)
}

func hasNext(students []int, sandwichesFirst int) bool {
	for _, s := range students {
		if s == sandwichesFirst {
			return true
		}
	}

	return false
}

func test01() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}
	succResult := 0
	fmt.Println("test01 students:=", students, " sandwiches:=", sandwiches)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countStudents(students, sandwiches))
	fmt.Println()
}

func test02() {
	students := []int{1, 1, 1, 0, 0, 1}
	sandwiches := []int{1, 0, 0, 0, 1, 1}
	succResult := 3
	fmt.Println("test02 students:=", students, " sandwiches:=", sandwiches)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countStudents(students, sandwiches))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
