package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	data []int
	size int
	min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([]int, 0),
		size: 0,
		min:  math.MinInt32,
	}
}

func (this *MinStack) Push(x int) {
	if this.size == 0 || this.min > x {
		this.min = x
	}

	this.data = append(this.data, x)
	this.size++
}

func (this *MinStack) Pop() {
	if this.size > 0 {
		this.size--
		n := this.data[this.size]
		this.data = this.data[0:this.size]
		if this.min == n && this.size > 0 {
			this.min = findMin(this.data)
		} else if this.size == 0 {
			this.min = math.MinInt32
		}
	}
}

func findMin(d []int) int {
	min := d[0]

	for _, v := range d {
		if min > v {
			min = v
		}
	}

	return min
}

func (this *MinStack) Top() int {
	if this.size > 0 {
		return this.data[this.size-1]
	}
	return -1
}

func (this *MinStack) GetMin() int {
	if this.size == 0 {
		return math.MinInt32
	}
	return this.min
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
}

func main() {
	test01()
	test02()
	test03()
}
