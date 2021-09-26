package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int
type KthLargest struct {
	h *IntHeap
	k int
}

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func (h *IntHeap) Top() int {
	if (*h).Len() == 0 {
		return 0
	}

	return (*h)[0]
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for _, v := range nums {
		heap.Push(h, v)
	}

	for h.Len() > k {
		heap.Pop(h)
	}

	return KthLargest{
		h: h,
		k: k,
	}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.h, val)
	fmt.Println(this.h)
	for this.h.Len() > this.k {
		heap.Pop(this.h)
	}

	return this.h.Top()
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
