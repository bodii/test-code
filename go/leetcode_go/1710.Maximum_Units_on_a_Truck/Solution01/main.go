package main

import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.SliceStable(boxTypes, func(i, j int) bool {
		return boxTypes[i][1] > boxTypes[j][1]
	})

	result, nums := 0, 0
	for _, v := range boxTypes {
		if nums+v[0] < truckSize {
			result += v[0] * v[1]
			nums += v[0]
		} else {
			result += (truckSize - nums) * v[1]
			break
		}
	}

	return result
}

func test01() {
	boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	truckSize := 4
	succResult := 8
	fmt.Println("test01 boxTypes:=", boxTypes, " truckSize:=", truckSize)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumUnits(boxTypes, truckSize))
	fmt.Println()
}

func test02() {
	boxTypes := [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}
	truckSize := 10
	succResult := 91
	fmt.Println("test02 boxTypes:=", boxTypes, " truckSize:=", truckSize)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumUnits(boxTypes, truckSize))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
