package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	result := make([]int, rowIndex+1)
	result[0] = 1
	for i := 1; i <= rowIndex; i++ {
		result[i] = result[i-1] * (rowIndex + 1 - i) / i
	}

	return result
}

func test01() {
	rowIndex := 3
	succResult := []int{1, 3, 3, 1}
	fmt.Println("test01 rowIndex:=", rowIndex)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getRow(rowIndex))
	fmt.Println()
}

func test02() {
	rowIndex := 0
	succResult := []int{1}
	fmt.Println("test02 rowIndex:=", rowIndex)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getRow(rowIndex))
	fmt.Println()
}

func test03() {
	rowIndex := 1
	succResult := []int{1, 1}
	fmt.Println("test03 rowIndex:=", rowIndex)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getRow(rowIndex))
	fmt.Println()
}

func test04() {
	rowIndex := 5
	succResult := []int{1, 1}
	fmt.Println("test04 rowIndex:=", rowIndex)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getRow(rowIndex))
	fmt.Println()
}

func test05() {
	rowIndex := 25
	succResult := []int{1, 1}
	fmt.Println("test05 rowIndex:=", rowIndex)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", getRow(rowIndex))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
