package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	scoreLen := len(score)
	scoreSort := make([]int, scoreLen)
	for k := range score {
		scoreSort[k] = k
	}

	sort.Slice(scoreSort, func(i, j int) bool {
		return score[scoreSort[i]] > score[scoreSort[j]]
	})

	result := make([]string, len(score))
	for k := range score {
		switch k {
		case 0:
			result[scoreSort[k]] = "Gold Medal"
		case 1:
			result[scoreSort[k]] = "Silver Medal"
		case 2:
			result[scoreSort[k]] = "Bronze Medal"
		default:
			result[scoreSort[k]] = strconv.Itoa(k + 1)
		}
	}

	return result
}

func test01() {
	score := []int{5, 4, 3, 2, 1}
	succResult := []string{"Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"}
	fmt.Println("test01 score:=", score)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRelativeRanks(score))
	fmt.Println()
}

func test02() {
	score := []int{10, 3, 8, 9, 4}
	succResult := []string{"Gold Medal", "5", "Bronze Medal", "Silver Medal", "4"}
	fmt.Println("test02 score:=", score)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", findRelativeRanks(score))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
