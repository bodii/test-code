package main

import  (
	"fmt"
)

type carI interface {
	run(exercise string)
}

type car struct {
	name string
	types string
}


func (c car) run(exercise string) {
	fmt.Printf("%s in %s speeding.\n", c.name, exercise)
}

func drive(c carI, exercise string) {
	c.run(exercise)
}


func main() {
	train := car {
		name: "green train",
		types: "train",
	}

	bicycle := car {
		name: "mountain bike",
		types: "bicycle",
	}

	drive(train, "the mountains")
	drive(bicycle, "the highway")
}
