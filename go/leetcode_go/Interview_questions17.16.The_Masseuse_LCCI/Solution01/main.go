package main

import (
	"fmt"
)

func massage(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}

	dp1 /*不接*/, dp2 /*接*/ := 0, nums[0]
	for i := 1; i < numsLen; i++ { // 从第2个预约开始
		// 接这次，上次不接的时间加上这次接的时长
		// 不接这次，上次接和不接的最大值
		dp2, dp1 = dp1+nums[i], max(dp1, dp2)
		fmt.Println(dp1, dp2)
	}

	return max(dp1, dp2)
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func test01() {
	nums := []int{1, 2, 3, 1}
	succResult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", massage(nums))
	fmt.Println()
}

func test02() {
	nums := []int{2, 7, 9, 3, 1}
	succResult := 12
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", massage(nums))
	fmt.Println()
}

func test03() {
	nums := []int{2, 1, 4, 5, 3, 1, 1, 3}
	succResult := 12
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", massage(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
