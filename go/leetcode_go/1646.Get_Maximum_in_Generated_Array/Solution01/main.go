package main

import (
	"fmt"
)

func getMaximumGenerated(n int) int {
	if n < 2 {
		return n
	}

	nums := make([]int, n+1)
	nums[0], nums[1] = 0, 1
	max := 1
	for i := 1; i <= n/2; i++ {
		nums[i*2] = nums[i]
		if nums[i] > max {
			max = nums[i]
		}

		if i*2+1 <= n {
			nums[i*2+1] = nums[i] + nums[i+1]
			if nums[i]+nums[i+1] > max {
				max = nums[i] + nums[i+1]
			}
		}

	}

	return max
}

func test01() {
	n := 7
	successResult := 3
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", getMaximumGenerated(n))
	fmt.Println()
}

func test02() {
	n := 2
	successResult := 1
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", getMaximumGenerated(n))
	fmt.Println()
}

func test03() {
	n := 3
	successResult := 2
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", getMaximumGenerated(n))
	fmt.Println()
}

func test04() {
	n := 9
	successResult := 4
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", getMaximumGenerated(n))
	fmt.Println()
}

func test05() {
	n := 100
	successResult := 21
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", getMaximumGenerated(n))
	fmt.Println()
}

func test06() {
	n := 0
	successResult := 0
	fmt.Println("test06 n:=", n)
	fmt.Println("success result:=", successResult)
	fmt.Println("result:=", getMaximumGenerated(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
