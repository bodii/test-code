package main

import "fmt"

func example01() {
	data := []int{1, 2, 3}

	for i, v := range data {
		data[i] = v * 10
	}

	fmt.Println("data: ", data)
}

func example02() {
	data := []*struct{ num int }{{1}, {2}, {3}}
	for _, v := range data {
		v.num *= 10
	}

	fmt.Println(data[0], data[1], data[2])
}

func main() {
	example01()
	example02()
}
