package main

import (
	"fmt"
)

func checkIfExist(arr []int) bool {
	doubles := make(map[int]int)
	for k, v := range arr {
		doubles[v*2] = k
	}

	for k, v := range arr {
		if val, ok := doubles[v]; ok && val != k {
			return true
		}
	}

	return false
}

func test01() {
	arr := []int{10, 2, 5, 3}
	succResult := true

	fmt.Println("test01:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func test02() {
	arr := []int{7, 1, 14, 11}
	succResult := true

	fmt.Println("test02:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func test03() {
	arr := []int{3, 1, 7, 11}
	succResult := false

	fmt.Println("test03:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func test04() {
	arr := []int{0, 0}
	succResult := true

	fmt.Println("test04:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func test05() {
	arr := []int{2, 3, 7, 0, 5, 8, 9}
	succResult := false

	fmt.Println("test04:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
