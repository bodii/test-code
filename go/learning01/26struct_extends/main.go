package main

import (
	"fmt"
)

type animal struct {
	name string
}


func (a animal) move() {
	fmt.Printf("%s move\n", a.name)
}

type dog struct {
	feet uint8
	animal
}

func (d dog)barking() {
	fmt.Printf("%s barking\n", d.name)
}

func main() {

	d := dog{
		feet: 4,
		animal: animal {
			name: "black dog",	
		},
	}

	d.barking()
	d.move()
}
