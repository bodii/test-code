package main

import (
	"container/list"
	"fmt"
)

type element struct {
	value int
	min   int
}

type MinStack struct {
	data *list.List
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: list.New(),
	}
}

func (this *MinStack) Push(x int) {
	min := x
	if this.data.Len() > 0 && x > this.data.Back().Value.(element).min {
		min = this.data.Back().Value.(element).min
	}

	this.data.PushBack(element{x, min})
}

func (this *MinStack) Pop() {
	this.data.Remove(this.data.Back())
}

func (this *MinStack) Top() int {
	return this.data.Back().Value.(element).value
}

func (this *MinStack) Min() int {
	return this.data.Back().Value.(element).min
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
