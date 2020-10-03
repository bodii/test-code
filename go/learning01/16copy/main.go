package main

import (
	"fmt"
)

func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	var a3 = make([]int, 3)
	copy(a3, a1)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	// 从切片中删除元素 
	a4 := []int{30, 31, 32, 33, 34, 35, 36, 37}
	a4 = append(a4[:2], a4[3:]...)
	fmt.Println(a4)

	var a = make([]string, 5, 10)
	for i:=0;i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i))
	}
	fmt.Println(a)
}
