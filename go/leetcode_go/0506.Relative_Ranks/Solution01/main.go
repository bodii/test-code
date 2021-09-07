package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	scoreSort := make([]int, len(score))
	copy(scoreSort, score)
	sort.Sort(sort.Reverse(sort.IntSlice(scoreSort)))

	scoreSortMap := make(map[int]string)
	for k, v := range scoreSort {
		if k == 0 {
			scoreSortMap[v] = "Gold Medal"
		} else if k == 1 {
			scoreSortMap[v] = "Silver Medal"
		} else if k == 2 {
			scoreSortMap[v] = "Bronze Medal"
		} else {
			scoreSortMap[v] = strconv.Itoa(k + 1)
		}
	}

	result := make([]string, len(score))
	for k, v := range score {
		result[k] = scoreSortMap[v]
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
