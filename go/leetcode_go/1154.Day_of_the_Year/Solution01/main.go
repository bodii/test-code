package main

import (
	"fmt"
	"strconv"
	"strings"
)

func dayOfYear(date string) int {
	dateArr := strings.Split(date, "-")
	onlyMonthdays := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	year, _ := strconv.Atoi(dateArr[0])
	month, _ := strconv.Atoi(dateArr[1])
	days, _ := strconv.Atoi(dateArr[2])

	resultDays := 0
	for i := 0; i < month-1; i++ {
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			onlyMonthdays[1] = 29
		}
		resultDays += onlyMonthdays[i]

	}
	resultDays += days
	return resultDays
}

func test01() {
	date := "2019-01-09"
	fmt.Printf("test01  %v result [9]:  %v\n", date, dayOfYear(date))
}

func test02() {
	date := "2003-03-01"
	fmt.Printf("test02  %v result [60]:  %v\n", date, dayOfYear(date))
}

func test03() {
	date := "2004-03-01"
	fmt.Printf("test03  %v result [61]:  %v\n", date, dayOfYear(date))
}

func test04() {
	date := "2019-02-10"
	fmt.Printf("test04  %v result [41]:  %v\n", date, dayOfYear(date))
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
