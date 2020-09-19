package main

import (
	"fmt"
)

// 地址类型的结构体
type address struct {
	province string
	city string
}

// 结构体嵌套
type person struct {
	name string
	age int
	addr address
}

// 匿名嵌套结构体
type company struct {
	name string
	address
}


func main() {
	// 嵌套结构体的使用
	p := person {
		name: "jack",
		age: 18,
		addr: address{
			province: "New York Region",
			city: "Oneida",
		},
	}
	fmt.Printf("%#v\n", p)
	fmt.Printf("%s the addr: %s province %s city\n", p.name, p.addr.province, p.addr.city)

	// 匿名嵌套结构体的使用
	c := company {
		name: "youke",
		address: address {
			province: "beijing",
			city: "beijing",
		},
	}
	fmt.Printf("%#v\n", c)
	fmt.Printf("company: %s city: %s\n", c.name, c.city)
}
