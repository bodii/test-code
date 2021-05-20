package main

import (
	"fmt"
)

func average(salary []int) float64 {
	salaryLen := len(salary)
	min, max := salary[0], salary[0]
	count := 0

	for _, v := range salary {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
		count += v
	}

	return float64((count - max - min)) / float64((salaryLen - 2))
}

func test01() {
	salary := []int{4000, 3000, 1000, 2000}
	succResult := 2500.00000
	fmt.Println("test01 salary:=", salary)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", average(salary))
	fmt.Println()
}

func test02() {
	salary := []int{1000, 2000, 3000}
	succResult := 2000.00000
	fmt.Println("test02 salary:=", salary)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", average(salary))
	fmt.Println()
}

func test03() {
	salary := []int{6000, 5000, 4000, 3000, 2000, 1000}
	succResult := 3500.00000
	fmt.Println("test03 salary:=", salary)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", average(salary))
	fmt.Println()
}

func test04() {
	salary := []int{8000, 9000, 2000, 3000, 6000, 1000}
	succResult := 4750.00000
	fmt.Println("test04 salary:=", salary)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", average(salary))
	fmt.Println()
}

func test05() {
	salary := []int{48000, 59000, 99000, 13000, 78000, 45000, 31000, 17000, 39000, 37000, 93000, 77000, 33000, 28000, 4000, 54000, 67000, 6000, 1000, 11000}
	succResult := 41111.11111
	fmt.Println("test05 salary:=", salary)
	fmt.Println("success result:=", succResult)
	fmt.Println("result:=", average(salary))
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
	test04()
	test05()
}
