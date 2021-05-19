package main

import (
	"fmt"
	"sort"
)

func canBeEqual(target []int, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)

	for k, v := range target {
		if v != arr[k] {
			return false
		}
	}

	return true
}

func test01() {
	target, arr := []int{1, 2, 3, 4}, []int{2, 4, 1, 3}
	succResult := true
	fmt.Println("test01 target:=", target)
	fmt.Println("test01 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeEqual(target, arr))
	fmt.Println()
}

func test02() {
	target, arr := []int{7}, []int{7}
	succResult := true
	fmt.Println("test02 target:=", target)
	fmt.Println("test02 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeEqual(target, arr))
	fmt.Println()
}

func test03() {
	target, arr := []int{1, 12}, []int{12, 1}
	succResult := true
	fmt.Println("test03 target:=", target)
	fmt.Println("test03 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeEqual(target, arr))
	fmt.Println()
}

func test04() {
	target, arr := []int{3, 7, 9}, []int{3, 7, 11}
	succResult := false
	fmt.Println("test04 target:=", target)
	fmt.Println("test04 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeEqual(target, arr))
	fmt.Println()
}

func test05() {
	target, arr := []int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}
	succResult := true
	fmt.Println("test05 target:=", target)
	fmt.Println("test05 arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", canBeEqual(target, arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
