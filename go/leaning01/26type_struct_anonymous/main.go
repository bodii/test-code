package main

import (
	"fmt"
)

// 匿名字段
// 字段比较少也比较简单的场景
// 不常用
type person struct {
	string
	int
}


func main() {
	p1 := person{
		"jack",
		28,
	}

	fmt.Printf("%#v\n", p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
