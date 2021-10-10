package main

import (
	"fmt"
	"math"
	"time"
)

func daysBetweenDates(date1 string, date2 string) int {
	layoutISO := "2006-01-02"
	t1, _ := time.Parse(layoutISO, date1)
	t2, _ := time.Parse(layoutISO, date2)
	return int(math.Abs(t1.Sub(t2).Hours())) / 24
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
