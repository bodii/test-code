package main

import "fmt"

func duplicateZeros(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		if arr[i] == 0 {
			i++
			if i < len {
				insert(&arr, i, 0, len)
			}
		}
	}
	//fmt.Println(arr)
}

func insert(arr *[]int, index, value, len int) {
	tmp := append([]int{}, (*arr)[index:len-1]...)
	*arr = append(append((*arr)[:index], value), tmp...)
}

func test01() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	fmt.Printf("test01 %v is [1,0,0,2,3,0,0,4] result:", arr)
	duplicateZeros(arr)
}

func test02() {
	arr := []int{1, 2, 3}
	fmt.Printf("test02 %v is [1,2,3] result:", arr)
	duplicateZeros(arr)
}

func test03() {
	arr := []int{0, 0, 0, 0, 0, 0, 0}
	fmt.Printf("test03 %v is [0,0,0,0,0,0,0] result:", arr)
	duplicateZeros(arr)
}

func main() {
	test01()
	test02()
	test03()
}
