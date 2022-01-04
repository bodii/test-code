package main

import (
	"fmt"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	val := image[sr][sc]
	rLen, cLen := len(image), len(image[0])
	var dfsFill func(i int, j int)
	dfsFill = func(i int, j int) {
		if i < 0 || i >= rLen || j < 0 || j >= cLen {
			return
		}

		if image[i][j] != val {
			return
		}

		if image[i][j] == -1 {
			return
		}

		image[i][j] = -1
		dfsFill(i-1, j)
		dfsFill(i+1, j)
		dfsFill(i, j-1)
		dfsFill(i, j+1)
		image[i][j] = newColor
	}

	dfsFill(sr, sc)
	return image
}

func test01() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	sr, sc := 1, 1
	newColor := 2
	succResult := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}
	fmt.Println("test01 image:=", image, " sr:=", sr, " sc:=", sc)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", floodFill(image, sr, sc, newColor))
	fmt.Println()
}

func test02() {
	image := [][]int{{1, 1, 1, 0}, {1, 0, 1, 0}, {1, 0, 0, 1}, {0, 1, 0, 1}, {1, 1, 1, 1}, {0, 1, 1, 0}}
	sr, sc := 2, 1
	newColor := 2
	succResult := [][]int{{1, 1, 1, 0}, {1, 2, 1, 0}, {1, 2, 2, 1}, {0, 1, 2, 1}, {1, 1, 1, 1}, {0, 1, 1, 0}}
	fmt.Println("test02 image:=", image, " sr:=", sr, " sc:=", sc)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", floodFill(image, sr, sc, newColor))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
