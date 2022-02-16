package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minimumSum(num int) int {
	numToByte := []byte(strconv.Itoa(num))
	sort.Slice(numToByte, func(i, j int) bool {
		return numToByte[i] < numToByte[j]
	})

	return int((numToByte[0]+numToByte[1])&15)*10 + int((numToByte[2]+numToByte[3])&15)
}

func test01() {
	num := 2932
	succResult := 52
	fmt.Println("test01 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumSum(num))
	fmt.Println()
}

func test02() {
	num := 4009
	succResult := 13
	fmt.Println("test02 num:=", num)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", minimumSum(num))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
