package main

import (
	"fmt"
)

func minPartitions(n string) int {
	max := n[0]
	left, right := 1, len(n)-1
	for left <= right {
		if n[left] > max {
			max = n[left]
		}

		if n[right] > max {
			max = n[right]
		}
		left++
		right--
	}

	return int(max - '0')
}

func test01() {
	n := "32"
	succResult := 3
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minPartitions(n))
	fmt.Println()
}

func test02() {
	n := "82734"
	succResult := 8
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minPartitions(n))
	fmt.Println()
}

func test03() {
	n := "27346209830709182346"
	succResult := 9
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minPartitions(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
