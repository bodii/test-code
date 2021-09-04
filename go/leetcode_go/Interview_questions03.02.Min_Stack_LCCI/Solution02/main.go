package main

import (
	"container/list"
	"fmt"
	"math"
)

type MinStack struct {
	data *list.List
}

type Elem struct {
	val int
	min int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{data: list.New()}
}

func (this *MinStack) Push(x int) {
	elem := Elem{val: x, min: x}
	if this.data.Len() > 0 {
		last := this.data.Back().Value.(Elem)
		if last.min < x {
			elem.min = last.min
		}
	}
	this.data.PushBack(elem)
}

func (this *MinStack) Pop() {
	if this.data.Len() > 0 {
		this.data.Remove(this.data.Back())
	}
}

func (this *MinStack) Top() int {
	if this.data.Len() > 0 {
		return this.data.Back().Value.(Elem).val
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.data.Len() > 0 {
		return this.data.Back().Value.(Elem).min
	}
	return math.MaxInt32
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func test01() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("GetMin() success result: -3, result:", minStack.GetMin())
	minStack.Pop()
	fmt.Println("Top() success result: 0, result:", minStack.Top())
	fmt.Println("GetMin() success result: -2, result:", minStack.GetMin())
	fmt.Println()
}

func test02() {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-1)
	fmt.Println("GetMin() success result: -2, result:", minStack.GetMin())
	fmt.Println("Top() success result: -1, result:", minStack.Top())
	minStack.Pop()
	fmt.Println("GetMin() success result: -2, result:", minStack.GetMin())
	fmt.Println()
}

func test03() {
	// ["MinStack","push","push","push",
	// "top","pop","getMin","pop","getMin","pop","push","top","getMin","push","top","getMin","pop","getMin"]
	// [[],[2147483646],[2147483646],[2147483647],[],[],[],[],[],[],[2147483647],[],[],[-2147483648],[],[],[],[]]

	minStack := Constructor()
	minStack.Push(2147483646)
	minStack.Push(2147483646)
	minStack.Push(2147483647)
	fmt.Println("Top() success result: 2147483647, result:", minStack.Top())
	minStack.Pop()
	fmt.Println("GetMin() success result: 2147483646, result:", minStack.GetMin())
	minStack.Pop()
	fmt.Println("GetMin() success result: 2147483646, result:", minStack.GetMin())
	minStack.Pop()
	minStack.Push(2147483647)
	fmt.Println("Top() success result: 2147483647, result:", minStack.Top())
	fmt.Println("GetMin() success result: 2147483647, result:", minStack.GetMin())
	minStack.Push(-2147483648)
	fmt.Println("Top() success result: -2147483648, result:", minStack.Top())
	fmt.Println("GetMin() success result: -2147483648, result:", minStack.GetMin())
	minStack.Pop()
	fmt.Println("GetMin() success result: 2147483647, result:", minStack.GetMin())
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
