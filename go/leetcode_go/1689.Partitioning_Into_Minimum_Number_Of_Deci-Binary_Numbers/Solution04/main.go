package main

import (
	"fmt"
)

func minPartitions(n string) int {
	max := 0
	for _, v := range n {
		if int(v-'0') > max {
			max = int(v - '0')
		}
	}
	return max
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
