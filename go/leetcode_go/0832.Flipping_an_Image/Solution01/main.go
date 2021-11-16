package main

import (
	"fmt"
)

func flipAndInvertImage(image [][]int) [][]int {
	lineLastIdx := len(image[0]) - 1
	lineMid := lineLastIdx / 2
	for k, line := range image {
		for i := 0; i <= lineMid; i++ {
			image[k][i], image[k][lineLastIdx-i] = line[lineLastIdx-i]^1, line[i]^1
		}
	}

	return image
}

func test01() {
	image := [][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}
	succResult := [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}
	fmt.Println("test01 image:=", image)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", flipAndInvertImage(image))
	fmt.Println()
}

func test02() {
	image := [][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}}
	succResult := [][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}
	fmt.Println("test02 image:=", image)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", flipAndInvertImage(image))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
