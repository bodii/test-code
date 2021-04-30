package main

import (
	"fmt"
	"time"
)

func repeatedNTimes(A []int) int {
	len := len(A)

	for i := 1; i < 4; i++ {
		for j := 0; j < len-i; j++ {
			// fmt.Println(A[j], A[j+i])
			if A[j] == A[i+j] {
				return A[j]
			}
		}
	}

	return -1
}

func test01() {
	arr := []int{1, 2, 3, 3}
	fmt.Printf("test01 %v is 3 result: %d\n", arr, repeatedNTimes(arr))
}

func test02() {
	arr := []int{2, 1, 2, 5, 3, 2}
	fmt.Printf("test02 %v is 2 result: %d\n", arr, repeatedNTimes(arr))
}

func test03() {
	arr := []int{5, 1, 5, 2, 5, 3, 5, 4}
	fmt.Printf("test03 %v is 5 result: %d\n", arr, repeatedNTimes(arr))
}

func main() {
	start := time.Now()

	test01()
	test02()
	test03()

	fmt.Println("start:", time.Since(start))
}
