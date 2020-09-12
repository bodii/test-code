package main

import (
	"fmt"
)

func main() {
	// 1. &:取地址
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)

	// 2. *:根据地址取值
	m := *p
	fmt.Println(m)

	var a = new(int)
	*a = 100
	fmt.Println(*a, a)
}
