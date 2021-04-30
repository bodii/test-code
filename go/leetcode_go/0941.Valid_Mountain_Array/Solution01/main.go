package main

import "fmt"

func validMountainArray(arr []int) bool {
	len := len(arr)
	if len < 3 {
		return false
	}

	if arr[0] > arr[1] {
		return false
	}

	isUp := true
	for i := 1; i < len-1; i++ {
		// fmt.Println(arr[i-1], arr[i], arr[i+1], isUp)
		if arr[i] > arr[i+1] {
			isUp = false
		}
		if arr[i-1] == arr[i] || arr[i+1] == arr[i] || (isUp && arr[i-1] > arr[i]) || (!isUp && arr[i+1] > arr[i]) {
			return false
		}
	}

	return !isUp
}

func test01() {
	arr := []int{2, 1}
	fmt.Printf("test01: %v, is Moutain false: %t\n", arr, validMountainArray(arr))
}

func test02() {
	arr := []int{3, 5, 5}
	fmt.Printf("test02: %v, is Moutain false: %t\n", arr, validMountainArray(arr))
}

func test03() {
	arr := []int{0, 3, 2, 1}
	fmt.Printf("test03: %v, is Moutain true: %t\n", arr, validMountainArray(arr))
}

func test04() {
	arr := []int{0, 2, 3, 3, 5, 2, 1, 0}
	fmt.Printf("test04: %v, is Moutain false: %t\n", arr, validMountainArray(arr))
}
func test05() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("test05: %v, is Moutain false: %t\n", arr, validMountainArray(arr))
}

func test06() {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Printf("test06: %v, is Moutain false: %t\n", arr, validMountainArray(arr))
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
	test06()
}
