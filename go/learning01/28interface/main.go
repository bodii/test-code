package main


import (
	"fmt"
)


type cat struct {
}

type dog struct {
}

type person struct {
}


// 定义一个接口类型
type barkingI interface {
	barking()
}


func (c cat) barking() {
	fmt.Println("Meow, Meow, Meow ~")
}

func (d dog) barking() {
	fmt.Println("Wow, Wow, Wow ~")
}

func (p person) barking() {
	fmt.Println("Ah, Ah, Ah ~")
}

func wake(b barkingI) {
	b.barking()
}


func main() {
	var (
		c cat
		d dog
		p person
	)

	wake(c)
	wake(d)
	wake(p)
}
