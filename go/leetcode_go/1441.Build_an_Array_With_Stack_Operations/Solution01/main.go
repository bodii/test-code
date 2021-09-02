package main

import (
	"fmt"
)

func buildArray(target []int, n int) []string {
	result := make([]string, 0)

	i := 1
	for j := 0; j < len(target); j++ {
		for ; i < target[j]; i++ {
			result = append(result, "Push", "Pop")
		}
		result = append(result, "Push")
		i++
	}

	return result
}

func test01() {
	target := []int{1, 3}
	n := 3
	succResult := []string{"Push", "Push", "Pop", "Push"}
	fmt.Println("test01 target:=", target, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buildArray(target, n))
	fmt.Println()
}

func test02() {
	target := []int{1, 2, 3}
	n := 3
	succResult := []string{"Push", "Push", "Push"}
	fmt.Println("test02 target:=", target, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buildArray(target, n))
	fmt.Println()
}

func test03() {
	target := []int{1, 2}
	n := 4
	succResult := []string{"Push", "Push"}
	fmt.Println("test03 target:=", target, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buildArray(target, n))
	fmt.Println()
}

func test04() {
	target := []int{2, 3, 4}
	n := 4
	succResult := []string{"Push", "Pop", "Push", "Push", "Push"}
	fmt.Println("test04 target:=", target, " n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", buildArray(target, n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
