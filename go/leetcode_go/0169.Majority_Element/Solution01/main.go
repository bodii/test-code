package main

import "fmt"

func majorityElement(nums []int) int {
	numsLen := len(nums)

	half := numsLen / 2
	box := make(map[int]int)
	for i := 0; i < numsLen; i++ {
		box[nums[i]]++
		if box[nums[i]] > half {
			return nums[i]
		}
	}

	return -1
}

func test01() {
	nums := []int{3, 2, 3}
	succResult := 3
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", majorityElement(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	succResult := 2
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", majorityElement(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
