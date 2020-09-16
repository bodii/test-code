package main

import (
	"fmt"
)

// 自定义类型和类型别名

// type后面跟的是类型
type myInt int

// 类型别名
type yourInt = int

func main() {
	var n myInt
	n = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n)

	var m yourInt
	m = 100
	fmt.Println(m)
	fmt.Printf("%T\n", m)


	var c rune
	c = 'h'
	fmt.Printf("%T %v\n", c, c)
}
