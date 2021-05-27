package main

import "fmt"

type MyStack struct {
	data []int
	top  int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		data: make([]int, 0),
		top:  -1,
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data = append(this.data, x)
	this.top++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	v := this.data[this.top]
	this.data = this.data[:this.top]
	this.top--
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.data[this.top]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.top < 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func test01() {
	fmt.Println("test01:")
	myStack := Constructor()
	myStack.Push(1)
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Empty())
	fmt.Println()
}

func test02() {
	fmt.Println("test02:")
	myStack := Constructor()
	myStack.Push(1)
	fmt.Println(myStack.Pop())
	myStack.Push(2)
	fmt.Println(myStack.Top())
	fmt.Println()
}

func main() {
	test01()
	test02()
}
