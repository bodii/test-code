package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() > 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}

	if this.stack2.Len() != 0 {
		e := this.stack2.Back()
		this.stack2.Remove(e)
		return e.Value.(int)
	}

	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func test01() {
	fmt.Println("test01:")
	cq := Constructor()
	cq.AppendTail(3)
	fmt.Println("result success: 3, result:", cq.DeleteHead())
	fmt.Println("result success: -1, result:", cq.DeleteHead())
}

func test02() {
	fmt.Println("test02:")
	cq := Constructor()
	fmt.Println("result success: -1, result:", cq.DeleteHead())
	cq.AppendTail(5)
	cq.AppendTail(2)
	fmt.Println("result success: 5, result:", cq.DeleteHead())
	fmt.Println("result success: 2, result:", cq.DeleteHead())
}

func test03() {
	fmt.Println("test03:")
	cq := Constructor()
	fmt.Println("result success: -1, result:", cq.DeleteHead())
	cq.AppendTail(5)
	cq.AppendTail(2)
	fmt.Println("result success: 5, result:", cq.DeleteHead())
	fmt.Println("result success: 2, result:", cq.DeleteHead())
	fmt.Println("result success: -1, result:", cq.DeleteHead())
	cq.AppendTail(3)
	cq.AppendTail(4)
	fmt.Println("result success: 3, result:", cq.DeleteHead())
	cq.AppendTail(5)
	fmt.Println("result success: 4, result:", cq.DeleteHead())

}

func main() {
	test01()
	test02()
	test03()
}
