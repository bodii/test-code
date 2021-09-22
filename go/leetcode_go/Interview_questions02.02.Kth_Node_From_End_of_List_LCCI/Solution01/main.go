package main

import "fmt"

var head = new(ListNode)

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
func kthToLast(head *ListNode, k int) int {
	list := make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	return list[len(list)-k]
}

func print(nodes *ListNode) {
	for nodes != nil {
		fmt.Println(nodes.Val)
		nodes = nodes.Next
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	nodes := head
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
		nodes.Next = node
		nodes = node
	}

	nodes = head
	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("success result: 4, result:=", kthToLast(nodes, k))
	fmt.Println()
}
