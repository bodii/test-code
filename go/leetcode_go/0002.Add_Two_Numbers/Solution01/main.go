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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num := 0
	root := &ListNode{Next: l1}
	prev := root
	for l1 != nil || l2 != nil || num > 0 {
		if prev.Next == nil {
			prev.Next = new(ListNode)
		}
		prev = prev.Next

		if l1 != nil {
			num += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num += l2.Val
			l2 = l2.Next
		}

		prev.Val = num % 10
		num /= 10
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
		fmt.Println("print node:", nodes.Val)
		nodes = nodes.Next
	}
}

func test01() {
	head1, head2 := new(ListNode), new(ListNode)
	l1 := []int{2, 4, 3}
	l2 := []int{5, 6, 4}
	nodes1 := head1
	insert(nodes1, l1)

	nodes2 := head2
	insert(nodes2, l2)

	fmt.Println("print:")
	print(nodes1)

	print(nodes2)

	fmt.Println("result")
	result := addTwoNumbers(nodes1, nodes2)
	print(result)
	fmt.Println("test01 success result: [7,0,8]")
	fmt.Println()
}

func test02() {
	head1, head2 := new(ListNode), new(ListNode)
	l1 := []int{0}
	l2 := []int{0}
	nodes1 := head1
	insert(nodes1, l1)

	nodes2 := head2
	insert(nodes2, l2)

	fmt.Println("print:")
	print(nodes1)

	print(nodes2)

	fmt.Println("result")
	result := addTwoNumbers(nodes1, nodes2)
	print(result)
	fmt.Println("test02 success result: [0]")
	fmt.Println()
}

func test03() {
	head1, head2 := new(ListNode), new(ListNode)
	l1 := []int{9, 9, 9, 9, 9, 9, 9}
	l2 := []int{9, 9, 9, 9}
	nodes1 := head1
	insert(nodes1, l1)

	nodes2 := head2
	insert(nodes2, l2)

	fmt.Println("print:")
	print(nodes1)

	print(nodes2)

	fmt.Println("result")
	result := addTwoNumbers(nodes1, nodes2)
	print(result)
	fmt.Println("test03 success result: [8,9,9,9,0,0,0,1]")
	fmt.Println()
}

func main() {
	test01()
	test02()
	test03()
}
