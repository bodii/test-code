package main

import (
	"fmt"
)

func kthDistinct(arr []string, k int) string {
	strs := make(map[string]int)

	for _, v := range arr {
		strs[v]++
	}

	num := 0
	for _, v := range arr {
		if strs[v] == 1 {
			num++
			if num == k {
				return v
			}
		}
	}

	return ""
}

func test01() {
	arr := []string{"d", "b", "c", "b", "c", "a"}
	k := 2
	succResult := "a"
	fmt.Println("test01 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kthDistinct(arr, k))
	fmt.Println()
}

func test02() {
	arr := []string{"aaa", "aa", "a"}
	k := 1
	succResult := "aaa"
	fmt.Println("test02 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kthDistinct(arr, k))
	fmt.Println()
}

func test03() {
	arr := []string{"a", "b", "a"}
	k := 3
	succResult := ""
	fmt.Println("test03 arr:=", arr, " k:=", k)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", kthDistinct(arr, k))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
