package main

import (
	"fmt"
)

func containsPattern(arr []int, m int, k int) bool {
	lastIdn := len(arr)

	for i := 0; i < lastIdn; {
		r, flag := 1, true
		for l := i + m; l < lastIdn; l += m {
			for j := 0; j < m; j++ {
				if l+j >= lastIdn || arr[i+j] != arr[l+j] {
					flag = false
					break
				}
			}
			if flag {
				r++
			} else {
				break
			}
		}

		if r >= k {
			return true
		}

		if r > 1 {
			i += r * m
		} else {
			i++
		}
	}

	return false
}

func test01() {
	arr := []int{1, 2, 4, 4, 4, 4}
	m, k := 1, 3
	succResult := true
	fmt.Println("test01 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func test02() {
	arr := []int{1, 2, 1, 2, 1, 1, 1, 3}
	m, k := 2, 2
	succResult := true
	fmt.Println("test02 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func test03() {
	arr := []int{1, 2, 1, 2, 1, 3}
	m, k := 2, 3
	succResult := false
	fmt.Println("test03 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func test04() {
	arr := []int{1, 2, 3, 1, 2}
	m, k := 2, 2
	succResult := false
	fmt.Println("test04 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func test05() {
	arr := []int{2, 2, 2, 2}
	m, k := 2, 3
	succResult := false
	fmt.Println("test05 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func test06() {
	arr := []int{3, 2, 2, 1, 2, 2, 1, 1, 1, 2, 3, 2, 2}
	m, k := 3, 2
	succResult := true
	fmt.Println("test06 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func test07() {
	arr := []int{2, 2}
	m, k := 1, 2
	succResult := true
	fmt.Println("test07 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func test08() {
	arr := []int{3, 2, 2, 1, 2, 2, 1, 1, 1, 2, 3, 2, 2}
	m, k := 3, 2
	succResult := true
	fmt.Println("test08 arr:=", arr, " m:=", m, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", containsPattern(arr, m, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
	test07()
	test08()
}
