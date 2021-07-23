package main

import (
	"fmt"
)

func maximumPopulation(logs [][]int) int {
	yearInds := make([]int, 100)

	for _, v := range logs {
		yearInds[v[0]-1950]++
		yearInds[v[1]-1950]--
	}

	yearInd, max, current := 0, 0, 0
	for k, v := range yearInds {
		current += v
		if current > max {
			max = current
			yearInd = k
		}
	}

	return yearInd + 1950
}

func test01() {
	logs := [][]int{{1993, 1999}, {2000, 2010}}
	succResult := 1993
	fmt.Println("test01 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumPopulation(logs))
	fmt.Println()
}

func test02() {
	logs := [][]int{{1950, 1961}, {1960, 1971}, {1970, 1981}}
	succResult := 1960
	fmt.Println("test02 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", maximumPopulation(logs))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
