package main

import (
	"fmt"
)

func minOperations(logs []string) int {
	step := 0

	for _, p := range logs {
		if p == "../" {
			if step > 0 {
				step -= 1
			}
		} else if p != "./" {
			step += 1
		}
	}

	return step
}

func test01() {
	logs := []string{"d1/", "d2/", "../", "d21/", "./"}
	succResult := 2
	fmt.Println("test01 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(logs))
	fmt.Println()
}

func test02() {
	logs := []string{"d1/", "d2/", "./", "d3/", "../", "d31/"}
	succResult := 3
	fmt.Println("test02 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(logs))
	fmt.Println()
}

func test03() {
	logs := []string{"d1/", "../", "../", "../"}
	succResult := 0
	fmt.Println("test03 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(logs))
	fmt.Println()
}

func test04() {
	logs := []string{"./", "../", "./"}
	succResult := 0
	fmt.Println("test04 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(logs))
	fmt.Println()
}

func test05() {
	logs := []string{"./", "wz4/", "../", "mj2/", "../", "../", "ik0/", "il7/"}
	succResult := 2
	fmt.Println("test05 logs:=", logs)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minOperations(logs))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
