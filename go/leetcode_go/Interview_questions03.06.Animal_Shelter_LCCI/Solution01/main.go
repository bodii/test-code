package main

import (
	"container/list"
	"fmt"
)

type AnimalShelf struct {
	data *list.List
}

func Constructor() AnimalShelf {
	return AnimalShelf{
		data: list.New(),
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	this.data.PushFront(animal)
}

func (this *AnimalShelf) DequeueAny() []int {
	if this.data.Len() == 0 {
		return []int{-1, -1}
	}

	animal := this.data.Back()
	this.data.Remove(animal)

	return animal.Value.([]int)
}

func (this *AnimalShelf) DequeueDog() []int {
	if this.data.Len() == 0 {
		return []int{-1, -1}
	}

	for e := this.data.Back(); e != nil; e = e.Prev() {
		animal := e.Value.([]int)
		if animal[1] == 1 {
			this.data.Remove(e)
			return animal
		}
	}

	return []int{-1, -1}
}

func (this *AnimalShelf) DequeueCat() []int {
	if this.data.Len() == 0 {
		return []int{-1, -1}
	}

	for e := this.data.Back(); e != nil; e = e.Prev() {
		animal := e.Value.([]int)
		if animal[1] == 0 {
			this.data.Remove(e)
			return animal
		}
	}

	return []int{-1, -1}
}

/**
 * Your AnimalShelf object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Enqueue(animal);
 * param_2 := obj.DequeueAny();
 * param_3 := obj.DequeueDog();
 * param_4 := obj.DequeueCat();
 */

func test01() {
	fmt.Println("test01: ")
	obj := Constructor()
	obj.Enqueue([]int{0, 0})
	obj.Enqueue([]int{1, 0})

	fmt.Println("DequeueCat success result: [0, 0], result: ", obj.DequeueCat())
	fmt.Println("DequeueDog success result: [-1, -1], result: ", obj.DequeueDog())
	fmt.Println("DequeueAny success result: [1, 0], result: ", obj.DequeueAny())

	fmt.Println()
}

func test02() {
	fmt.Println("test02: ")
	obj := Constructor()
	obj.Enqueue([]int{0, 0})
	obj.Enqueue([]int{1, 0})
	obj.Enqueue([]int{2, 1})

	fmt.Println("DequeueDog success result: [2, 1], result: ", obj.DequeueDog())
	fmt.Println("DequeueCat success result: [0, 0], result: ", obj.DequeueCat())
	fmt.Println("DequeueAny success result: [1, 0], result: ", obj.DequeueAny())

	fmt.Println()
}

func main() {
	test01()
	test02()
}
