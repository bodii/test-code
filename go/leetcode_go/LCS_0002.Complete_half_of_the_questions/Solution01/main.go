package main

import (
	"fmt"
	"sort"
)

func halfQuestions(questions []int) int {
	questionHash := make([]int, 1001)
	half := len(questions) / 2

	for _, v := range questions {
		questionHash[v]++
	}

	sort.Sort(sort.Reverse(sort.IntSlice(questionHash)))

	i := 0
	for ; i < 1001; i++ {
		if half <= 0 || questionHash[i] == 0 {
			break
		}
		half -= questionHash[i]
	}

	return i
}

func test01() {
	questions := []int{2, 1, 6, 2}
	succResult := 1
	fmt.Println("test01 questions:=", questions)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", halfQuestions(questions))
	fmt.Println()
}

func test02() {
	questions := []int{1, 5, 1, 3, 4, 5, 2, 5, 3, 3, 8, 6}
	succResult := 2
	fmt.Println("test02 questions:=", questions)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", halfQuestions(questions))
	fmt.Println()
}

func test03() {
	questions := []int{8, 11, 6, 10}
	succResult := 2
	fmt.Println("test03 questions:=", questions)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", halfQuestions(questions))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
