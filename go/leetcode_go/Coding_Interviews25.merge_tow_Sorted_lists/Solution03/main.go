package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	prev := root
	nums := []int{}

	for l1 != nil {
		nums = append(nums, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		nums = append(nums, l2.Val)
		l2 = l2.Next
	}

	sort.Ints(nums)

	for _, v := range nums {
		prev.Next = &ListNode{Val: v}
		prev = prev.Next
	}

	return root.Next
}

func insert(nodes *ListNode, nums []int) {
	numsLen := len(nums)
	if numsLen < 1 {
		return
	}

	nodes.Val = nums[0]
	for i := 1; i < numsLen; i++ {
		node := &ListNode{Val: nums[i]}
		nodes.Next = node
		nodes = node
	}
}

func print(nodes *ListNode) {
	for nodes != nil {
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

func test01() {
	var head1, head2 *ListNode
	head1, head2 = new(ListNode), new(ListNode)
	l1, l2 := []int{1, 2, 4}, []int{1, 3, 4}
	nodes1, nodes2 := head1, head2
	insert(nodes1, l1)
	insert(nodes2, l2)

	fmt.Println("print:")
	print(nodes1)
	print(nodes2)

	fmt.Println()
	result := mergeTwoLists(nodes1, nodes2)
	print(result)
	fmt.Println("test01 success result: [1,1,2,3,4,4]")
	fmt.Println()
}

func test02() {
	var head1, head2 *ListNode
	head1, head2 = new(ListNode), new(ListNode)
	l1, l2 := []int{}, []int{}
	nodes1, nodes2 := head1, head2
	insert(nodes1, l1)
	insert(nodes2, l2)

	fmt.Println("print:")
	print(nodes1)
	print(nodes2)

	fmt.Println()
	result := mergeTwoLists(nodes1, nodes2)
	print(result)
	fmt.Println("test02 success result: []")
	fmt.Println()
}

func test03() {
	var head1, head2 *ListNode
	head1, head2 = new(ListNode), new(ListNode)
	l1, l2 := []int{}, []int{0}
	nodes1, nodes2 := head1, head2
	insert(nodes1, l1)
	insert(nodes2, l2)
	nodes1, nodes2 = head1, head2
	fmt.Println("print:")
	print(nodes1)
	print(nodes2)

	fmt.Println()
	result := mergeTwoLists(nodes1, nodes2)
	print(result)
	fmt.Println("test03 success result: [0]")
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
