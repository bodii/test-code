package main

import "fmt"

func arraysIntersection(arr1 []int, arr2 []int, arr3 []int) []int {
	result := make([]int, 0)
	for _, v := range arr1 {
		num := 1
		for _, v2 := range arr2 {
			if v2 > v {
				break
			}

			if v == v2 {
				num++
			}
		}

		for _, v3 := range arr3 {
			if v3 > v {
				break
			}

			if v == v3 {
				num++
			}
		}

		if num == 3 {
			result = append(result, v)
		}
	}

	return result
}

func test01() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{1, 2, 5, 7, 9}
	arr3 := []int{1, 3, 4, 5, 8}

	succResult := []int{1, 5}

	fmt.Println("testt01:")
	fmt.Println("arr1:=", arr1, " arr2:=", arr2, " arr3:=", arr3)
	fmt.Println("success result:", succResult)
	fmt.Println("result:", arraysIntersection(arr1, arr2, arr3))
	fmt.Println()
}

func main() {
	test01()
}
