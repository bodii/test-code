package main

import "fmt"

func checkPossibility(nums []int) bool {
	stat := 0
	numsLastInd := len(nums) - 1
	for i := numsLastInd; i > 0; i-- {
		if nums[i] < nums[i-1] {
			stat++
			pre, next := i+1, i-1
			if pre <= numsLastInd {
				if nums[pre] > nums[next] {
					nums[i] = nums[pre]
				} else {
					nums[next] = nums[i]
				}
				i++
			}
		}

		if stat == 2 {
			return false
		}
	}

	return true
}

func test01() {
	nums := []int{3, 4, 2, 3}
	succResult := false
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkPossibility(nums))
	fmt.Println()
}

func test02() {
	nums := []int{5, 7, 1, 8}
	succResult := true
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkPossibility(nums))
	fmt.Println()
}

func test03() {
	nums := []int{4, 2, 3}
	succResult := true
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", checkPossibility(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
