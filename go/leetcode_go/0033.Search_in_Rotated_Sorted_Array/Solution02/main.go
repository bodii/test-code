package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[l] {
			if nums[mid] > target && nums[l] <= target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}

func test01() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	succResult := 4
	fmt.Println("test01:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test02() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 3
	succResult := -1
	fmt.Println("test02:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test03() {
	nums := []int{1}
	target := 0
	succResult := -1
	fmt.Println("test03:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test04() {
	nums := []int{1}
	target := 1
	succResult := 0
	fmt.Println("test04:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test05() {
	nums := []int{3, 5, 1}
	target := 5
	succResult := 1
	fmt.Println("test05:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test06() {
	nums := []int{1, 3}
	target := 3
	succResult := 1
	fmt.Println("test06:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test07() {
	nums := []int{1, 3}
	target := 1
	succResult := 0
	fmt.Println("test07:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test08() {
	nums := []int{3, 5, 1}
	target := 3
	succResult := 0
	fmt.Println("test08:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
	fmt.Println()
}

func test09() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 5
	succResult := 1
	fmt.Println("test09:")
	fmt.Println("nums:", nums, " target:=", target)
	fmt.Println("success result: ", succResult)
	fmt.Println("result:", search(nums, target))
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
	test09()
}
