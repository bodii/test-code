package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	list []int
	len  int
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	numsLen := len(nums)
	if numsLen > k {
		nums = nums[0:k]
		numsLen = k
	}

	return KthLargest{
		list: nums,
		len:  numsLen,
		k:    k,
	}
}

func (this *KthLargest) Add(val int) int {
	this.list = append(this.list, val)
	this.len++
	sort.Sort(sort.Reverse(sort.IntSlice(this.list)))

	if this.len <= this.k {
		return this.list[this.len-1]
	} else {
		result := this.list[this.k-1]
		this.list = this.list[0:this.k]
		this.len = this.k
		return result
	}
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
