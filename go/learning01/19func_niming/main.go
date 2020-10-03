package main

import (
	"fmt"
)

// 匿名函数
var f1 = func(x, y int) {
	fmt.Println(x, y)
}


// 函数体内不能声明带名字的函数
func main() {
	f2 := func(x, y int) {
		fmt.Println(x, y)
	}

	f1(1, 2)
	f2(1, 2)


	// 立即执行函数
	func(x, y int) {
		fmt.Println(x, y)
	}(1, 2)
}
