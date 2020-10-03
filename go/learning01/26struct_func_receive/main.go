package main

import "fmt"

// 结构体
type person struct {
	name string
	age int
}

// 构造方法
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age: age,
	}
}

// 指定接收者方法： 传拷贝
/*
func (p person) birthday() {
	p.age ++
}
*/


// 指定接收者方法: 传内存地址
func (p *person) realBirthday() {
	p.age ++	
}

func (p *person) dream() {
	fmt.Println(p.name, "the dream is singer")
}


func main() {
	p := newPerson("jack", 18)
	fmt.Printf("%#v\n", p)

	/*
	p.birthday()
	fmt.Println(p.age)
	*/

	p.realBirthday()
	fmt.Println(p.age)

	p.dream()
}
