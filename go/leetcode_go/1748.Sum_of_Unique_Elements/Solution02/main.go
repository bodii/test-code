package main

import (
	"fmt"
)

func sumOfUnique(nums []int) int {
	result := 0                   // 要返回的结果
	possibles := make([]int, 101) // 生成可能出现的所有数数组(slice): nums[i]<=100

	for _, v := range nums {
		if possibles[v] == 0 { // 如果当前数不存在
			possibles[v] = 1 // 就把当前数放进可能出现的数数组中标记为:1
			result += v      // 并将当前数累加到要返回的数
		} else if possibles[v] == 1 { // 如果当前数存在过
			possibles[v] = -1 //  就把当前数指向的可能出现的数数组中标记为:-1
			result -= v       // 并将要返回的数减去当前数
		}
	}

	return result
}

func test01() {
	nums := []int{1, 2, 3, 2}
	succresultult := 4
	fmt.Println("test01 nums:=", nums)
	fmt.Println("success resultult:=", succresultult)
	fmt.Println("resultult:=", sumOfUnique(nums))
	fmt.Println()
}

func test02() {
	nums := []int{1, 1, 1, 1, 1}
	succresultult := 0
	fmt.Println("test02 nums:=", nums)
	fmt.Println("success resultult:=", succresultult)
	fmt.Println("resultult:=", sumOfUnique(nums))
	fmt.Println()
}

func test03() {
	nums := []int{1, 2, 3, 4, 5}
	succresultult := 15
	fmt.Println("test03 nums:=", nums)
	fmt.Println("success resultult:=", succresultult)
	fmt.Println("resultult:=", sumOfUnique(nums))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
