package main

import (
	"container/list"
	"fmt"
)

type AnimalShelf struct {
	catList *list.List
	dogList *list.List
}

func Constructor() AnimalShelf {
	return AnimalShelf{
		catList: list.New(),
		dogList: list.New(),
	}
}

func (this *AnimalShelf) Enqueue(animal []int) {
	if animal[1] == 0 {
		this.catList.PushFront(animal[0])
	} else {
		this.dogList.PushFront(animal[0])
	}
}

func (this *AnimalShelf) DequeueAny() []int {
	if this.catList.Len() == 0 && this.dogList.Len() == 0 {
		return []int{-1, -1}
	} else if this.catList.Len() == 0 {
		return this.DequeueDog()
	} else if this.dogList.Len() == 0 {
		return this.DequeueCat()
	}

	if this.catList.Back().Value.(int) < this.dogList.Back().Value.(int) {
		return this.DequeueCat()
	}

	return this.DequeueDog()
}

func (this *AnimalShelf) DequeueDog() []int {
	if this.dogList.Len() == 0 {
		return []int{-1, -1}
	}

	dog := this.dogList.Back()
	this.dogList.Remove(dog)

	return []int{dog.Value.(int), 1}
}

func (this *AnimalShelf) DequeueCat() []int {
	if this.catList.Len() == 0 {
		return []int{-1, -1}
	}

	cat := this.catList.Back()
	this.catList.Remove(cat)

	return []int{cat.Value.(int), 0}
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
