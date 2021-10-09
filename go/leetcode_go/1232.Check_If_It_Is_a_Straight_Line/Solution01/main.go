package main

import "fmt"

// 二阶矩阵的行列式
func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates) - 1
	for i := 0; i < l; i++ {
		if (coordinates[i][0]-coordinates[0][0])*(coordinates[i][1]-coordinates[l][1]) !=
			(coordinates[i][1]-coordinates[0][1])*(coordinates[i][0]-coordinates[l][0]) {
			return false
		}
	}

	return true
}

func test01() {
	coordinates := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}
	successResult := true
	fmt.Println("test01 coordinates:=", coordinates)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", checkStraightLine(coordinates))
	fmt.Println()
}

func test02() {
	coordinates := [][]int{{1, 1}, {2, 2}, {3, 4}, {4, 5}, {5, 6}, {7, 7}}
	successResult := false
	fmt.Println("test02 coordinates:=", coordinates)
	fmt.Println("success result:", successResult)
	fmt.Println("result: ", checkStraightLine(coordinates))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
