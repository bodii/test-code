package main

import (
	"fmt"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	rr := len(image) - 1
	cr := len(image[0]) - 1
	irl, irr := sr, sr
	jcl, jcr := sc, sc

	for irl > 0 {
		if image[irl-1] == image[irl] {
			irl--
			image[irl] = newColor
		}
	}

	for irr < rr {
		if image[irr+1] == image[irr] {
			irr++
			image[irr] = newColor
		}
	}

	for jcl > 0 {
		if image[jcl-1] == image[jcl] {
			jcl--
			image[jcl] = newColor
		}
	}

	for jcr < cr {
		if image[jcr+1] == image[jcr] {
			jcr++
			image[jcr] = newColor
		}
	}

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

func main() {
	test01()
}
