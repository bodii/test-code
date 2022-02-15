package main

import "fmt"

func checkIfExist(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]*2 == arr[j] || arr[i] == arr[j]*2 {
				return true
			}
		}
	}

	return false
}

func test01() {
	arr := []int{10, 2, 5, 3}
	succResult := true

	fmt.Println("test01:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func test02() {
	arr := []int{7, 1, 14, 11}
	succResult := true

	fmt.Println("test02:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func test03() {
	arr := []int{3, 1, 7, 11}
	succResult := false

	fmt.Println("test03:")
	fmt.Println("arr:=", arr)
	fmt.Println("success result:=", succResult)
	fmt.Println("result: ", checkIfExist(arr))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
