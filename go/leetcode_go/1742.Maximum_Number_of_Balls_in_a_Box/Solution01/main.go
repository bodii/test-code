package main

import "fmt"

func countBalls(lowLimit int, highLimit int) int {
	box := make(map[int]int)

	for i := lowLimit; i <= highLimit; i++ {
		number := onesSum(i)
		box[number-1]++
	}

	max := box[0]
	for _, v := range box {
		if v > max {
			max = v
		}
	}

	return max
}

func onesSum(i int) int {
	sum := 0
	for i > 0 {
		sum += i % 10
		i /= 10
	}

	return sum
}

func test01() {
	lowLimit, highLimit := 1, 10
	succResult := 2
	fmt.Println("test01 lowLimit:=", lowLimit, " highLimit:=", highLimit)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBalls(lowLimit, highLimit))
	fmt.Println()
}

func test02() {
	lowLimit, highLimit := 5, 15
	succResult := 2
	fmt.Println("test02 lowLimit:=", lowLimit, " highLimit:=", highLimit)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBalls(lowLimit, highLimit))
	fmt.Println()
}

func test03() {
	lowLimit, highLimit := 19, 28
	succResult := 2
	fmt.Println("test03 lowLimit:=", lowLimit, " highLimit:=", highLimit)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBalls(lowLimit, highLimit))
	fmt.Println()
}

func test04() {
	lowLimit, highLimit := 11, 104
	succResult := 9
	fmt.Println("test04 lowLimit:=", lowLimit, " highLimit:=", highLimit)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", countBalls(lowLimit, highLimit))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
}
