package main

import "fmt"

func canFormArray(arr []int, pieces [][]int) bool {
	arrLen := len(arr)

	for k1 := 0; k1 < arrLen; k1++ {
		isIn := false
		for _, v2 := range pieces {
			if arr[k1] == v2[0] {
				v2Len := len(v2)
				if arrLen-k1 < v2Len {
					return false
				}

				for k3, v3 := range v2 {
					if arr[k1+k3] != v3 {
						return false
					}
				}
				k1 += v2Len - 1
				isIn = true
			}
		}
		if !isIn {
			return false
		}
	}

	return true
}

func test01() {
	arr := []int{85}
	pieces := [][]int{{85}}
	successResult := true
	fmt.Println("test01 arr:=", arr)
	fmt.Println("pieces:=", pieces)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", canFormArray(arr, pieces))
	fmt.Println()
}

func test02() {
	arr := []int{15, 88}
	pieces := [][]int{{88}, {15}}
	successResult := true
	fmt.Println("test02 arr:=", arr)
	fmt.Println("pieces:=", pieces)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", canFormArray(arr, pieces))
	fmt.Println()
}

func test03() {
	arr := []int{49, 18, 16}
	pieces := [][]int{{16, 18, 46}}
	successResult := false
	fmt.Println("test03 arr:=", arr)
	fmt.Println("pieces:=", pieces)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", canFormArray(arr, pieces))
	fmt.Println()
}

func test04() {
	arr := []int{91, 4, 64, 78}
	pieces := [][]int{{78}, {4, 64}, {91}}
	successResult := true
	fmt.Println("test04 arr:=", arr)
	fmt.Println("pieces:=", pieces)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", canFormArray(arr, pieces))
	fmt.Println()
}
func test05() {
	arr := []int{1, 3, 5, 7}
	pieces := [][]int{{2, 4, 6, 8}}
	successResult := false
	fmt.Println("test05 arr:=", arr)
	fmt.Println("pieces:=", pieces)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", canFormArray(arr, pieces))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
