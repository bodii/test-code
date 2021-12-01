package main

import (
	"fmt"
)

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return !(rec1[0] >= rec2[2] || rec1[2] <= rec2[0]) &&
		!(rec1[1] >= rec2[3] || rec1[3] <= rec2[1])
}

func test01() {
	rec1, rec2 := []int{0, 0, 2, 2}, []int{1, 1, 3, 3}
	succResult := true
	fmt.Println("test01:")
	fmt.Println("rec1:=", rec1, " rec2:=", rec2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isRectangleOverlap(rec1, rec2))
	fmt.Println()
}

func test02() {
	rec1, rec2 := []int{0, 0, 1, 1}, []int{1, 0, 2, 1}
	succResult := false
	fmt.Println("test02:")
	fmt.Println("rec1:=", rec1, " rec2:=", rec2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isRectangleOverlap(rec1, rec2))
	fmt.Println()
}

func test03() {
	rec1, rec2 := []int{0, 0, 1, 1}, []int{2, 2, 3, 3}
	succResult := false
	fmt.Println("test03:")
	fmt.Println("rec1:=", rec1, " rec2:=", rec2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isRectangleOverlap(rec1, rec2))
	fmt.Println()
}

func test04() {
	rec1, rec2 := []int{0, 0, 1, 1}, []int{0, 0, 1, 1}
	succResult := true
	fmt.Println("test04:")
	fmt.Println("rec1:=", rec1, " rec2:=", rec2)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", isRectangleOverlap(rec1, rec2))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
