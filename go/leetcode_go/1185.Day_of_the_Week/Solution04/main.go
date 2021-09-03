package main

import (
	"fmt"
)

func dayOfTheWeek(day int, month int, year int) string {
	weeks := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	return weeks[getDaysWeekIndex(year, month, day)]
}

/*
** 思路：
基姆拉尔森计算公式
W= (d+2m+3(m+1)/5+y+y/4-y/100+y/400+1)%7 //C++计算公式
在公式中d表示日期中的日数，m表示月份数，y表示年数。
w表示星期，w的取值范围是0~6，分别代表星期一到星期日。
注意：在公式中有个与其他公式不同的地方：
把一月和二月看成是上一年的十三月和十四月，例：如果是2004-1-10则换算成：2003-13-10来代入公式计算。

作者：CodingTwoandahalfyearsTrainee
链接：https://leetcode-cn.com/problems/day-of-the-week/solution/1185-yi-zhou-zhong-de-di-ji-tian-zui-jia-lk9q/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// getDaysWeekIndex 获取天数对应的周索引
func getDaysWeekIndex(year, month, day int) int {
	if month <= 2 {
		month += 12
		year -= 1
	}

	return (day + 2*month + 3*(month+1)/5 + year + year/4 - year/100 + year/400 + 1) % 7
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
