package main

import (
	"fmt"
)

func decrypt(code []int, k int) []int {
	len := len(code)
	result := make([]int, len)

	if k > 0 {
		for i := 0; i < len; i++ {
			for j := 0; j < k; j++ {
				result[i] += code[(i+1+j)%len]
			}
		}
	} else if k < 0 {
		for i := 0; i < len; i++ {
			for j := 0; j < -k; j++ {
				result[i] += code[(len+k+i+j)%len]
			}
		}
	}

	return result
}

func test01() {
	code := []int{5, 7, 1, 4}
	k := 3
	succResult := []int{12, 10, 16, 13}
	fmt.Println("test01 code:=", code, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decrypt(code, k))
	fmt.Println()
}

func test02() {
	code := []int{1, 2, 3, 4}
	k := 0
	succResult := []int{0, 0, 0, 0}
	fmt.Println("test02 code:=", code, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decrypt(code, k))
	fmt.Println()
}

func test03() {
	code := []int{2, 4, 9, 3}
	k := -2
	succResult := []int{12, 5, 6, 13}
	fmt.Println("test03 code:=", code, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decrypt(code, k))
	fmt.Println()
}

func test04() {
	code := []int{10, 5, 7, 7, 3, 2, 10, 3, 6, 9, 1, 6}
	k := -4
	succResult := []int{22, 26, 22, 28, 29, 22, 19, 22, 18, 21, 28, 19}
	fmt.Println("test04 code:=", code, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", decrypt(code, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
