package main

import (
	"fmt"
)

func destCity(paths [][]string) string {
	cityA := make(map[string]bool)
	cityB := make(map[string]bool)
	for _, v := range paths {
		cityA[v[0]] = true
		cityB[v[1]] = true
		if cityB[v[0]] {
			delete(cityB, v[0])
		}

		if cityA[v[1]] {
			delete(cityB, v[1])
		}
	}

	for k := range cityB {
		return k
	}

	return ""
}

func test01() {
	paths := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	succResult := "Sao Paulo"
	fmt.Println("test01 paths:=", paths)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", destCity(paths))
	fmt.Println()
}

func test02() {
	paths := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
	succResult := "A"
	fmt.Println("test02 paths:=", paths)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", destCity(paths))
	fmt.Println()
}

func test03() {
	paths := [][]string{{"A", "Z"}}
	succResult := "Z"
	fmt.Println("test03 paths:=", paths)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", destCity(paths))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
