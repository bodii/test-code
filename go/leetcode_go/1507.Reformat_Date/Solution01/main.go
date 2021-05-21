package main

import (
	"fmt"
	"strings"
)

func reformatDate(date string) string {
	Months := map[string]string{"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04",
		"May": "05", "Jun": "06", "Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10",
		"Nov": "11", "Dec": "12"}
	Days := map[string]string{
		"1st":  "01",
		"2nd":  "02",
		"3rd":  "03",
		"4th":  "04",
		"5th":  "05",
		"6th":  "06",
		"7th":  "07",
		"8th":  "08",
		"9th":  "09",
		"10th": "10",
		"11th": "11",
		"12th": "12",
		"13th": "13",
		"14th": "14",
		"15th": "15",
		"16th": "16",
		"17th": "17",
		"18th": "18",
		"19th": "19",
		"20th": "20",
		"21st": "21",
		"22nd": "22",
		"23rd": "23",
		"24th": "24",
		"25th": "25",
		"26th": "26",
		"27th": "27",
		"28th": "28",
		"29th": "29",
		"30th": "30",
		"31st": "31",
	}
	dateArr := strings.Split(date, " ")
	DayStr, MonthStr, Year := dateArr[0], dateArr[1], dateArr[2]
	Day, ok := Days[DayStr]
	if !ok {
		Day = ""
	}

	Month, ok := Months[MonthStr]
	if !ok {
		Month = ""
	}

	return fmt.Sprintf("%s-%s-%s", Year, Month, Day)

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
