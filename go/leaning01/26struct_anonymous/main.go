package main

import (
	"fmt"
)

type person struct {
	name, gender string
}

func f (p person) {
	p.gender = "woman"
}

func f2(p *person) {
	//(*p).gender = "woman"
	p.gender = "woman"
}


type x struct {
	a int8
	b int8
	c int8
	d string
}

func main() {
	var s struct {
		name string	
		age int8
	}

	s.name = "hehe"
	s.age = 18

	fmt.Printf("%#v %T\n", s, s)


	var p person
	p.name = "lilei"
	p.gender = "man"

	f(p)
	fmt.Printf("%#v\n", p)

	f2(&p)
	fmt.Printf("%#v\n", p)

	// 指针
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	p2.name = "aa"
	p2.gender = "man"
	fmt.Println(p2)
	fmt.Println(&p2)


	// 结构体指针
	// 声明变量并初始化
	// key-value 初始化
	var p3 = person {
		name: "hello",
		gender: "man",
	}
	fmt.Printf("%#v\n", p3)

	p4 := person {
		"jack",
		"man",
	}
	fmt.Printf("%#v\n", p4)

	m := x {
		a: int8(10),
		b: int8(20),
		c: int8(30),
		d: string("c"),
	}

	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))
	fmt.Printf("%p\n", &(m.d))

}
