package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	t1, t2 := dateToTime(date1), dateToTime(date2)

	diffDays := int(t1.Sub(t2).Hours() / 24)

	if diffDays < 0 {
		return -diffDays
	}

	return diffDays
}

func dateToTime(date string) time.Time {
	dates := strings.Split(date, "-")
	year, _ := strconv.Atoi(dates[0])
	month, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
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
