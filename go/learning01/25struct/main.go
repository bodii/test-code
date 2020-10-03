package main

import (
	"fmt"
)

// 结构体
type person struct {
	name string
	age	 int
	gender string
	hobby []string
}

func main() {
	var p1 person
	p1.name = "Tom"
	p1.age = 18
	p1.gender = "man"
	p1.hobby = []string{"listen to music", "to song", "reading"}
	
	fmt.Printf("%#v\n", p1)

	// access variable
	fmt.Println("name:", p1.name)
}
