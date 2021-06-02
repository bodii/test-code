package main

import "fmt"

func selfDividingNumbers(left int, right int) []int {
	result := make([]int, 0)
	for i := left; i <= right; i++ {
		if isDividingNumber(i) {
			result = append(result, i)
		}
	}

	return result
}

func isDividingNumber(n int) bool {
	c := n
	for c > 0 {
		if c%10 == 0 || n%(c%10) != 0 {
			return false
		}

		c /= 10
	}

	return true
}

func test01() {
	left, right := 1, 22
	success := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}

	fmt.Println("test01  left:=", left, " right:=", right)
	fmt.Println("success result:", success)
	fmt.Println("result:", selfDividingNumbers(left, right))
	fmt.Println()
}

func test02() {
	left, right := 47, 85
	success := []int{48, 55, 66, 77}

	fmt.Println("test02  left:=", left, " right:=", right)
	fmt.Println("success result:", success)
	fmt.Println("result:", selfDividingNumbers(left, right))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
