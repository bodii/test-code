package main

import "fmt"

/*
type animal interface {
	move()
	eat(string)
}
*/
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}


type cat struct {
	name string
	feet int8
}

// 使用值接收者实现接口方法
//func (c cat) move() {
// 使用值接收者指针实现接口方法
func (c *cat) move() {
	fmt.Println("cat walking")
}

func (c *cat) eat(food string) {
	fmt.Println("cat eat", food)
}


func main() {
	c := cat {"Tom", 4}
	c2 := cat {"Garfield", 4}

	var a animal
	a = &c
	fmt.Println(a)
	a = &c2
	fmt.Println(a)
}
