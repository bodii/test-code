package main

import (
	"fmt"
)

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	num := 0
	for k, v := range startTime {
		if queryTime >= v && queryTime <= endTime[k] {
			num++
		}
	}

	return num
}

func test01() {
	startTime, endTime := []int{1, 2, 3}, []int{3, 2, 7}
	queryTime := 4
	succResult := 1
	fmt.Println("test01 startTime:=", startTime, " endTime:=", endTime)
	fmt.Println("queryTime:=", queryTime)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", busyStudent(startTime, endTime, queryTime))
	fmt.Println()
}

func test02() {
	startTime, endTime := []int{4}, []int{4}
	queryTime := 4
	succResult := 1
	fmt.Println("test02 startTime:=", startTime, " endTime:=", endTime)
	fmt.Println("queryTime:=", queryTime)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", busyStudent(startTime, endTime, queryTime))
	fmt.Println()
}

func test03() {
	startTime, endTime := []int{4}, []int{4}
	queryTime := 5
	succResult := 0
	fmt.Println("test03 startTime:=", startTime, " endTime:=", endTime)
	fmt.Println("queryTime:=", queryTime)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", busyStudent(startTime, endTime, queryTime))
	fmt.Println()
}

func test04() {
	startTime, endTime := []int{1, 1, 1, 1}, []int{1, 3, 2, 4}
	queryTime := 7
	succResult := 0
	fmt.Println("test04 startTime:=", startTime, " endTime:=", endTime)
	fmt.Println("queryTime:=", queryTime)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", busyStudent(startTime, endTime, queryTime))
	fmt.Println()
}

func test05() {
	startTime, endTime := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}
	queryTime := 5
	succResult := 5
	fmt.Println("test05 startTime:=", startTime, " endTime:=", endTime)
	fmt.Println("queryTime:=", queryTime)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", busyStudent(startTime, endTime, queryTime))
	fmt.Println()
}

func test06() {
	startTime, endTime := []int{16}, []int{60}
	queryTime := 58
	succResult := 1
	fmt.Println("test06 startTime:=", startTime, " endTime:=", endTime)
	fmt.Println("queryTime:=", queryTime)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", busyStudent(startTime, endTime, queryTime))
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
