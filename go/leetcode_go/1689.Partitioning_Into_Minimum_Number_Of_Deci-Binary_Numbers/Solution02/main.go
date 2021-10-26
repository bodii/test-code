package main

import (
	"fmt"
	"sort"
)

func minPartitions(n string) int {
	nBytes := []byte(n)
	sort.SliceStable(nBytes, func(i, j int) bool {
		return nBytes[i] > nBytes[j]
	})

	return int(nBytes[0] - '0')
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
