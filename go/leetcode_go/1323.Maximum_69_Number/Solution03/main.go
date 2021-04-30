package main

import (
	"fmt"
)

func maximum69Number(num int) int {
	next := num
	nextP := 10000
	for i := 4; i >= 1; i-- {
		nextP /= 10
		if next/nextP == 6 {
			num += 3 * nextP
			return num
		}
		// fmt.Println(nextP, next, next/nextP, next%nextP, num)
		next %= nextP
	}
	return num
}

func test01() {
	num := 9669
	fmt.Println("test01 is 9669 result 9969:", maximum69Number(num))
}

func test02() {
	num := 9996
	fmt.Println("test02 is 9996 result 9999:", maximum69Number(num))
}

func test03() {
	num := 9999
	fmt.Println("test03 is 9999 result 9999:", maximum69Number(num))
}

func test04() {
	num := 9969
	fmt.Println("test04 is 9969 result 9999:", maximum69Number(num))
}

func test05() {
	num := 6969
	fmt.Println("test05 is 6969 result 9969:", maximum69Number(num))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
