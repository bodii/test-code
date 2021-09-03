package main

import (
	"fmt"
	"time"
)

func dayOfTheWeek(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local).Weekday().String()
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
