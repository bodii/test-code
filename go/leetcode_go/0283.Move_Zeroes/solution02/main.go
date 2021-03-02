package main

import "fmt"

func moveZeroes(nums []int) {
	l := len(nums)
	j := 0
	for i := 0; i < l; i++ {
		if 0 != nums[i] {
			nums[j] = nums[i]
			j++
		}
	}

	for i := j; i < l; i++ {
		nums[i] = 0
	}

	fmt.Println(nums)
}

func main() {
	nums := []int{0, 1, 0, 3, 12}

	moveZeroes(nums)
}
