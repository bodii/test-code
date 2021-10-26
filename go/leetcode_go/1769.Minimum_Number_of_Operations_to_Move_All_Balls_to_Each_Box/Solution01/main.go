package main

import (
	"fmt"
)

func minOperations(boxes string) []int {
	result := make([]int, len(boxes))
	for k := range boxes {
		result[k] = findDistance(boxes, k)
	}

	return result
}

func findDistance(boxes string, k int) int {
	distance := 0
	for i := range boxes {
		if boxes[i] == '1' {
			distance += abs(i - k)
		}
	}

	return distance
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func test01() {
	boxes := "110"
	success := []int{1, 1, 3}

	fmt.Println("test01  boxes:=", boxes)
	fmt.Println("success result:", success)
	fmt.Println("result:", minOperations(boxes))
	fmt.Println()
}

func test02() {
	boxes := "001011"
	success := []int{11, 8, 5, 4, 3, 4}

	fmt.Println("test01  boxes:=", boxes)
	fmt.Println("success result:", success)
	fmt.Println("result:", minOperations(boxes))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
