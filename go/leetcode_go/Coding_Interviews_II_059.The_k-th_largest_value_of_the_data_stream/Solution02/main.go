package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type numsHeap struct {
	sort.IntSlice
}
type KthLargest struct {
	list *numsHeap
	k    int
}

func (n *numsHeap) Push(v interface{}) { n.IntSlice = append(n.IntSlice, v.(int)) }
func (n *numsHeap) push(v int)         { heap.Push(n, v) }
func (n *numsHeap) Pop() interface{} {
	old := n.IntSlice
	v := old[len(old)-1]
	n.IntSlice = old[:len(old)-1]
	return v
}
func (n *numsHeap) pop() int { return heap.Pop(n).(int) }

func (n numsHeap) Top() int {
	if n.Len() == 0 {
		return 0
	}
	return n.IntSlice[0]
}

func Constructor(k int, nums []int) KthLargest {
	h := &numsHeap{}
	kthLargest := KthLargest{
		list: h,
		k:    k,
	}

	for _, v := range nums {
		kthLargest.Add(v)
	}

	return kthLargest
}

func (this *KthLargest) Add(val int) int {
	this.list.push(val)
	fmt.Println(this.list)
	for this.list.Len() > this.k {
		this.list.pop()
	}
	fmt.Println(this.list)

	return this.list.Top()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func test01() {
	fmt.Println("test01 :")
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	param_1 := obj.Add(3)
	fmt.Println("param_1 success result:=4, result:=", param_1)
	param_1 = obj.Add(5)
	fmt.Println("param_1 success result:=5, result:=", param_1)
	param_1 = obj.Add(10)
	fmt.Println("param_1 success result:=5, result:=", param_1)
	param_1 = obj.Add(9)
	fmt.Println("param_1 success result:=8, result:=", param_1)
	param_1 = obj.Add(4)
	fmt.Println("param_1 success result:=8, result:=", param_1)
	fmt.Println()
}

func main() {
	test01()
}
