package main

import (
	"fmt"
)

func minOperations(boxes string) []int {
	total, l, r := 0, 0, 0
	if boxes[0] == '1' {
		l = 1
	}

	for i := 1; i < len(boxes); i++ {
		if boxes[i] == '1' {
			r++
			total += i
		}
	}

	result := make([]int, len(boxes))
	result[0] = total

	for i := 1; i < len(boxes); i++ {
		total += l - r
		if boxes[i] == '1' {
			l++
			r--
		}
		result[i] = total
	}

	return result
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
