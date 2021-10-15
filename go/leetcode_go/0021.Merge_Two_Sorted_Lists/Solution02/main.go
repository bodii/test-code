package main

import (
	"fmt"
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
	head := new(ListNode)
	prev := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}

		prev = prev.Next
	}

	if l1 == nil {
		prev.Next = l2
	} else if l2 == nil {
		prev.Next = l1
	}

	return head.Next
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
