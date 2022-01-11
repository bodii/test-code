package main

import (
	"fmt"
)

func maximumTime(time string) string {
	time2 := []byte(time)
	if time2[0] == '?' {
		if time2[1] == '?' || time2[1]-'0' <= 3 {
			time2[0] = '2'
		} else {
			time2[0] = '1'
		}
	}

	if time2[1] == '?' {
		if time2[0] == '2' {
			time2[1] = '3'
		} else {
			time2[1] = '9'
		}
	}

	if time2[3] == '?' {
		time2[3] = '5'
	}

	if time2[4] == '?' {
		time2[4] = '9'
	}

	return string(time2)
}

func test01() {
	time := "2?:?0"
	successResult := "23:50"
	fmt.Println("test01 time:", time)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", maximumTime(time))
	fmt.Println()
}

func test02() {
	time := "0?:3?"
	successResult := "09:39"
	fmt.Println("test02 time:", time)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", maximumTime(time))
	fmt.Println()
}

func test03() {
	time := "1?:22"
	successResult := "19:22"
	fmt.Println("test03 time:", time)
	fmt.Println("success result:", successResult)
	fmt.Println("result:", maximumTime(time))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
