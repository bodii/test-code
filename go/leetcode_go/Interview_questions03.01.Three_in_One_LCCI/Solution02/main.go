package main

import (
	"fmt"
)

type TripleInOne struct {
	stacks    [3][]int
	stackSize int
	tops      [3]int
}

func Constructor(stackSize int) TripleInOne {
	return TripleInOne{
		stacks:    [3][]int{},
		stackSize: stackSize,
		tops:      [3]int{},
	}
}

func (this *TripleInOne) Push(stackNum int, value int) {
	if this.tops[stackNum] >= this.stackSize {
		return
	}

	if this.tops[stackNum] >= len(this.stacks[stackNum]) {
		this.stacks[stackNum] = append(this.stacks[stackNum], value)
	} else {
		this.stacks[stackNum][this.tops[stackNum]] = value
	}
	this.tops[stackNum]++
}

func (this *TripleInOne) Pop(stackNum int) int {
	if this.IsEmpty(stackNum) {
		return -1
	}

	this.tops[stackNum]--
	return this.stacks[stackNum][this.tops[stackNum]]
}

func (this *TripleInOne) Peek(stackNum int) int {
	if this.IsEmpty(stackNum) {
		return -1
	}

	return this.stacks[stackNum][this.tops[stackNum]-1]
}

func (this *TripleInOne) IsEmpty(stackNum int) bool {
	return this.tops[stackNum] <= 0
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
