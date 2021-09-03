package main

import (
	"fmt"
)

func dayOfTheWeek(day int, month int, year int) string {
	weeks := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	return weeks[getDaysWeekIndex(year, month, day)]
}

/*
给定年月日计算星期几的公式有多种，此处列举一个计算方便的公式。
W = x + [x / 4] + [y / 4] - 2y + [26(m + 1) / 10] + d - 1
其中，y表示年份前两位，x表示年份后两位，m表示月，
1月和2月看成上一年的13月和14月，d表示日。
方括号表示向下取整。
对W计算除以7的余数即可知道是星期几，余数是0表示星期日，余数是1到6分别表示星期一到星期六。
需要注意的是W可能是负值，因此计算除以7的余数之后需要判断是大于等于0还是小于0，如果小于0则将余数加7。
by Storm
*/
// getDaysWeekIndex 获取天数对应的周索引
func getDaysWeekIndex(year, month, day int) int {
	if month < 3 {
		month += 12
		year--
	}

	y1, y2 := year/100, year%100

	index := y2 + y2/4 + y1/4 - 2*y1 + 26*(month+1)/10 + day - 1

	index %= 7
	if index < 0 {
		index += 7
	}

	return index
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
