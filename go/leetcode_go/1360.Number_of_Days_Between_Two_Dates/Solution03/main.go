package main

import (
	"fmt"
	"strconv"
	"strings"
)

func daysBetweenDates(date1 string, date2 string) int {
	diff := getDays(date1) - getDays(date2)

	if diff < 0 {
		return -diff
	}

	return diff
}

func getDays(date string) int {
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	dates := strings.Split(date, "-")
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])

	days := 0
	for y := 1971; y < year; y++ {
		days += 365 + isLeapDay(y)
	}

	for m := 0; m < month-1; m++ {
		if m == 1 {
			days += isLeapDay(year) + 28
		} else {
			days += months[m]
		}
	}

	return days + day
}

func isLeapDay(a int) int {
	if (a%100 != 0 && a%4 == 0) || a%400 == 0 {
		return 1
	}

	return 0
}

func test01() {
	date1, date2 := "2019-06-29", "2019-06-30"
	succResult := 1
	fmt.Println("test01 date1:=", date1, " date2:=", date2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", daysBetweenDates(date1, date2))
	fmt.Println()
}

func test02() {
	date1, date2 := "2020-01-15", "2019-12-31"
	succResult := 15
	fmt.Println("test02 date1:=", date1, " date2:=", date2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", daysBetweenDates(date1, date2))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
