package main

import "fmt"

type myInt int32

type person struct {
	name string
	age int
}

func newPerson(name string, age int) person{
	return person{
		name: name,
		age: age,
	}
}

func newPerson2(name string, age int) *person {
	return &person {
		name: name,
		age: age,
	}
}

func main() {
	var x int32
	x = 10
	fmt.Println(x)

	var x1 int32 = 10
	fmt.Println(x1)

	var x2 = int32(10)
	fmt.Println(x2)

	x3 := 10
	fmt.Println(x3)

	var x4 myInt
	x4 = 10
	fmt.Println(x4)

	var x5 myInt = 10
	fmt.Println(x5)

	var x6 = myInt(10)
	fmt.Println(x6)

	var p person
	p.name = "jack"
	p.age = 18
	fmt.Printf("%#v\n", p)


	var p2 = person{
		name: "make",
		age : 19,
	}
	fmt.Printf("%#v\n", p2)
}
