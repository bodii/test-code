package main

import (
	"fmt"
)

func dayOfTheWeek(day int, month int, year int) string {
	weeks := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	days := 0
	for y := 1971; y < year; y++ {
		days += 365
		if isLeapYear(y) {
			days += 1
		}
	}

	for m := 1; m < month; m++ {
		days += getCurrentMonthTodays(m, year)
	}

	days += day + 4
	return weeks[days%7]
}

func getCurrentMonthTodays(month int, year int) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 2:
		if isLeapYear(year) {
			return 29
		} else {
			return 28
		}
	default:
		return 30
	}
}

func isLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	} else if year%4 == 0 && year%100 != 0 {
		return true
	}

	return false
}

func test01() {
	day, month, year := 31, 8, 2019
	success := "Saturday"

	fmt.Println("test01  day:=", day, " month:=", month, " year:=", year)
	fmt.Println("success result:", success)
	fmt.Println("result:", dayOfTheWeek(day, month, year))
	fmt.Println()
}

func test02() {
	day, month, year := 18, 7, 1999
	success := "Sunday"

	fmt.Println("test02  day:=", day, " month:=", month, " year:=", year)
	fmt.Println("success result:", success)
	fmt.Println("result:", dayOfTheWeek(day, month, year))
	fmt.Println()
}

func test03() {
	day, month, year := 15, 8, 1993
	success := "Sunday"

	fmt.Println("test03  day:=", day, " month:=", month, " year:=", year)
	fmt.Println("success result:", success)
	fmt.Println("result:", dayOfTheWeek(day, month, year))
	fmt.Println()
}

func test04() {
	day, month, year := 1, 1, 1971
	success := "Friday"

	fmt.Println("test04  day:=", day, " month:=", month, " year:=", year)
	fmt.Println("success result:", success)
	fmt.Println("result:", dayOfTheWeek(day, month, year))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
