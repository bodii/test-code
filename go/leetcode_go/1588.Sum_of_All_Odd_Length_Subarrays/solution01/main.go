package main

import "fmt"

func sumOddLengthSubarrays(arr []int) int {
	len := len(arr)
	result := 0

	for i := 0; i < len; i++ {
		left, right := (i + 1), (len - i)
		leftOdd, rightOdd := left/2, right/2
		leftEven, rightEven := (left+1)/2, (right+1)/2

		result += arr[i] * (leftEven*rightEven + leftOdd*rightOdd)
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
