package main

import (
	"fmt"
)

type TripleInOne struct {
	stacks    [][]int
	stackSize int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stacks:    make([][]int, 3),
		stackSize: stackSize,
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if len(this.stacks[stackNum]) < this.stackSize {
		this.stacks[stackNum] = append(this.stacks[stackNum], value)
	}
}

func (this *TripleInOne) Pop(stackNum int) int {
	val := this.Peek(stackNum)
	if val != -1 {
		this.stacks[stackNum] = this.stacks[stackNum][0 : len(this.stacks[stackNum])-1]
		return val
	}

	return val
}

func (this *TripleInOne) Peek(stackNum int) int {
	if !this.IsEmpty(stackNum) {
		return this.stacks[stackNum][len(this.stacks[stackNum])-1]
	}

	return -1
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return len(this.stacks[stackNum]) <= 0
}

/**
 * Your TripleInOne object will be instantiated and called as such:
 * obj := Constructor(stackSize);
 * obj.Push(stackNum,value);
 * param_2 := obj.Pop(stackNum);
 * param_3 := obj.Peek(stackNum);
 * param_4 := obj.IsEmpty(stackNum);
 */

func test01() {
	obj := Constructor(1)
	obj.Push(0, 1)
	obj.Push(0, 2)
	fmt.Println("success result: 1, result:", obj.Pop(0))
	fmt.Println("success result: -1, result:", obj.Pop(0))
	fmt.Println("success result: -1, result:", obj.Pop(0))
	fmt.Println("success result: true, result:", obj.IsEmpty(0))
	fmt.Println()
}

func test02() {
	obj := Constructor(2)
	obj.Push(0, 1)
	obj.Push(0, 2)
	obj.Push(0, 3)
	fmt.Println("success result: 2, result:", obj.Pop(0))
	fmt.Println("success result: 1, result:", obj.Pop(0))
	fmt.Println("success result: -1, result:", obj.Pop(0))
	fmt.Println("success result: -1, result:", obj.Peek(0))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
