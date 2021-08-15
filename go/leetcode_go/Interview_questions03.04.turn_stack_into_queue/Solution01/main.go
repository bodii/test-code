package main

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	stack1, stack2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.stack1.PushBack(x)
}

/* element of stack1 move to stack2 */
func (this *MyQueue) stack1ToStack2() {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.stack1ToStack2()

	if this.stack2.Len() > 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}

	return -1
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.stack1ToStack2()

	if this.stack2.Len() > 0 {
		e := this.stack2.Back()
		return e.Value.(int)
	}

	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.stack1.Len() == 0 && this.stack2.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func test01() {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	fmt.Println("success result: 1, result:", q.Peek())
	fmt.Println("success result: 1, result:", q.Pop())
	fmt.Println("success result: false, result:", q.Empty())
}

func main() {
	test01()
}
