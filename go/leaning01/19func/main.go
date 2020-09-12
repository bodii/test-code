package main

import (
	"fmt"
)


func sum(x int, y int) (ret int) {
	return x + y
}

func f1(x int, y int) {
	fmt.Println(x, y)
}

func f2() {
	fmt.Println("f2 function")
}

func f3() int {
	return 3
}

func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

func f5() (int, string) {
	return 1, "val"
}


func f6(x, y int) int {
	return x + y
}

// 可变长参数
// 可变长参数必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
	f1(1, 2)
	f2()

	r2 := f3()
	fmt.Println(r2)

	r3 := f4(1, 2)
	fmt.Println(r3)

	r4, v := f5()
	fmt.Println(r4, v)

	r6 := f6(1, 2)
	fmt.Println(r6)

	f7("string",  1, 3, 4)
	f7("string02")
}
