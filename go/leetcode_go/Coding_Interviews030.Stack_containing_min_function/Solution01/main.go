package main

import (
	"fmt"
)

type element struct {
	value int
	min   int
}

type MinStack struct {
	data []element
	size int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]element, 0),
		size: 0,
	}
}

func (this *MinStack) Push(x int) {
	ele := element{x, x}
	if this.size > 0 && this.data[this.size-1].min < x {
		ele.min = this.data[this.size-1].min
	}
	this.data = append(this.data, ele)
	this.size++
}

func (this *MinStack) Pop() {
	this.data = this.data[0 : this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.data[this.size-1].value
}

func (this *MinStack) Min() int {
	return this.data[this.size-1].min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */

func test01() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("GetMin() success result: -3, result:", minStack.Min())
	minStack.Pop()
	fmt.Println("Top() success result: 0, result:", minStack.Top())
	fmt.Println("Min() success result: -2, result:", minStack.Min())
}

func main() {
	test01()
}
