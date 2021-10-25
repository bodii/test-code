package main

import (
	"fmt"
)

func reorderedPowerOf2(n int) bool {
	counts := count(n)
	for i := 0; i < 31; i++ {
		if sliceEquals(counts, count(1<<i)) {
			return true
		}
	}

	return false
}

func sliceEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if v != b[k] {
			return false
		}
	}

	return true
}

func count(n int) []int {
	result := make([]int, 10)
	for n > 0 {
		result[n%10]++
		n /= 10
	}

	return result
}

func test01() {
	n := 1
	succResult := true
	fmt.Println("test01 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reorderedPowerOf2(n))
	fmt.Println()
}

func test02() {
	n := 10
	succResult := false
	fmt.Println("test02 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reorderedPowerOf2(n))
	fmt.Println()
}

func test03() {
	n := 16
	succResult := true
	fmt.Println("test03 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reorderedPowerOf2(n))
	fmt.Println()
}

func test04() {
	n := 24
	succResult := false
	fmt.Println("test04 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reorderedPowerOf2(n))
	fmt.Println()
}

func test05() {
	n := 46
	succResult := true
	fmt.Println("test05 n:=", n)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", reorderedPowerOf2(n))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
