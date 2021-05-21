package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	Months := map[string]string{"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04",
		"May": "05", "Jun": "06", "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10",
		"Nov": "11", "Dec": "12"}

	dateArr := strings.Split(date, " ")
	Day := ""
	if len(dateArr[0]) == 3 {
		Day = "0" + dateArr[0][0:1]
	} else {
		Day = dateArr[0][0:2]
	}

	Month, ok := Months[dateArr[1]]
	if !ok {
		Month = ""
	}

	return fmt.Sprintf("%s-%s-%s", dateArr[2], Month, Day)
}

func test01() {
	date := "20th Oct 2052"
	succResult := "2052-10-20"
	fmt.Println("test01 date:=", date)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatDate(date))
	fmt.Println()
}

func test02() {
	date := "6th Jun 1933"
	succResult := "1933-06-06"
	fmt.Println("test02 date:=", date)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatDate(date))
	fmt.Println()
}

func test03() {
	date := "26th May 1960"
	succResult := "1960-05-26"
	fmt.Println("test03 date:=", date)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reformatDate(date))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
