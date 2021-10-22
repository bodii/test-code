package main

import (
	"fmt"
)

func canChoose(groups [][]int, nums []int) bool {
	gLen, g2Len := len(groups), len(groups[0])
	d1, d2 := 0, 0

	for i := 0; i < len(nums); i++ {
		if g2Len <= d2 {
			d1++

			if d1 >= gLen {
				return true
			}
			g2Len, d2 = len(groups[d1]), 0
		}

		if nums[i] == groups[d1][d2] {
			d2++
		} else {
			i -= d2
			d2 = 0
		}
	}

	return d1+1 == gLen && d2 == g2Len
}

func test01() {
	groups := [][]int{{1, -1, -1}, {3, -2, 0}}
	nums := []int{1, -1, 0, 1, -1, -1, 3, -2, 0}
	succResult := true
	fmt.Println("test01 nums:=", nums)
	fmt.Println("groups:=", groups)

	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canChoose(groups, nums))
	fmt.Println()
}

func test02() {
	groups := [][]int{{10, -2}, {1, 2, 3, 4}}
	nums := []int{1, 2, 3, 4, 10, -2}
	succResult := false
	fmt.Println("test02 nums:=", nums)
	fmt.Println("groups:=", groups)

	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canChoose(groups, nums))
	fmt.Println()
}

func test03() {
	groups := [][]int{{1, 2, 3}, {3, 4}}
	nums := []int{7, 7, 1, 2, 3, 4, 7, 7}
	succResult := false
	fmt.Println("test03 nums:=", nums)
	fmt.Println("groups:=", groups)

	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canChoose(groups, nums))
	fmt.Println()
}

func test04() {
	groups := [][]int{{6636698, 4693069, -2178984, -2253405, -2732131, 8550889, -2324014, -2561263}, {-8973571, -9146179, 7704305, -1063430, -6569826}, {2791091}, {-9691134, 651171, 9835817, 4163881, 4944714, 8166788, -9025553, -9250995, 1599781}}
	nums := []int{8550889, -2178984, 6636698, 4693069, -2178984, -2253405, -2732131, 8550889, -2324014, -2561263, -2324014, 8550889, -8973571, -9146179, 7704305, -1063430, -6569826, 2791091, -9691134, 651171, 9835817, 4163881, 4944714, 8166788, -9025553, -9250995, 1599781}
	succResult := true
	fmt.Println("test04 nums:=", nums)
	fmt.Println("groups:=", groups)

	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canChoose(groups, nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
