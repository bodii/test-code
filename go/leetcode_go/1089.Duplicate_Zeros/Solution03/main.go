package main

import "fmt"

func duplicateZeros(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		if arr[i] == 0 {
			i++
			if i < len {
				insertAndMove(&arr, i, len)
			}
		}
	}
	// fmt.Println(arr)
}
func insertAndMove(arr *[]int, index, len int) {
	tmp := make([]int, len-1-index)
	copy(tmp, (*arr)[index:len-1])
	(*arr)[index] = 0
	for k, v := range tmp {
		(*arr)[index+k+1] = v
	}
}

func test01() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Printf("test01 %v is [1,0,0,2,3,0,0,4] result:", arr)
	duplicateZeros(arr)
	fmt.Println(arr)
}

func test02() {
	arr := []int{1, 2, 3}
	fmt.Printf("test02 %v is [1,2,3] result:", arr)
	duplicateZeros(arr)
	fmt.Println(arr)
}

func test03() {
	arr := []int{0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("test03 %v is [0,0,0,0,0,0,0] result:", arr)
	duplicateZeros(arr)
	fmt.Println(arr)
}

func main() {
	test01()
	test02()
	test03()
}
