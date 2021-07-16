package main

import (
	"fmt"
)

type Node struct {
	val int
	min int
}

type MinStack struct {
	nodes []Node
	size  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	n := Node{val: val, min: val}
	if this.size > 0 && this.nodes[this.size-1].min < val {
		n.min = this.nodes[this.size-1].min
	}

	this.nodes = append(this.nodes, n)
	this.size++
}

func (this *MinStack) Pop() {
	if this.size > 0 {
		this.size--
		this.nodes = this.nodes[:this.size]
	}
}

func (this *MinStack) Top() int {
	return this.nodes[this.size-1].val
}

func (this *MinStack) GetMin() int {
	return this.nodes[this.size-1].min
}

func test01() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
}

func main() {
	test01()
}
