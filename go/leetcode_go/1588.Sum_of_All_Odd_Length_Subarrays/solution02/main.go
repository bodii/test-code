package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	result := 0
	len := len(arr)

	for i := 1; i <= len; i += 2 {
		// j + i 为每一行取的数量
		for j := 0; j+i <= len; j++ {
			// k 为每一行子数组的下标
			for k := j; k < j+i; k++ {
				result += arr[k]
			}
		}
	}

	return result
}

func test1() {
	arr := []int{1, 4, 2, 5, 3}
	fmt.Println("result: ", sumOddLengthSubarrays(arr))
}

func test2() {
	arr := []int{1, 2}
	fmt.Println("result: ", sumOddLengthSubarrays(arr))
}

func test3() {
	arr := []int{10, 11, 12}
	fmt.Println("result: ", sumOddLengthSubarrays(arr))
}

func main() {
	// test1()

	test2()

	// test3()
}
