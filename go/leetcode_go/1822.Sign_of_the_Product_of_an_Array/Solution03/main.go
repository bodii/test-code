package main

import (
	"fmt"
)

func arraySign(nums []int) int {
	// 负数的个数， 如果负数的个数为奇数最终返回-1,因为除0以外，
	// 奇数个负数，乘以任何个大于0的数都得负数
	negatives := 0
	i, j := 0, len(nums)-1

	for i <= j {
		if nums[i] == 0 { // 如果有一个数是0，则返回0，任何数与0相乘都得0
			return 0
		} else if nums[i] < 0 {
			negatives++
		}
		i++

		if i >= j {
			continue
		}

		if nums[j] == 0 { // 如果有一个数是0，则返回0，任何数与0相乘都得0
			return 0
		} else if nums[j] < 0 {
			negatives++
		}
		j--
	}

	if negatives%2 == 0 {
		return 1
	}

	return -1
}

func test01() {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	succResult := 1
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arraySign(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 5, 0, 2, -3}
	succResult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arraySign(nums))
	fmt.Println()
}

func test03() {
	nums := []int{-1, 1, -1, 1, -1}
	succResult := -1
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", arraySign(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
