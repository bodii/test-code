package main

import "fmt"

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
func getKthFromEnd(head *ListNode, k int) *ListNode {
	len := 0
	for node := head; node != nil; node = node.Next {
		len++
	}

	for i := len; i > k; i-- {
		head = head.Next
	}

	return head
}

func insert(nodes *ListNode, nums []int) {
	nodes.Val = nums[0]
	for i := 1; i < len(nums); i++ {
		node := &ListNode{nums[i], nil}
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
	var head = new(ListNode)
	nums := []int{1, 2, 3, 4, 5}
	k := 2
	nodes := head
	insert(nodes, nums)
	nodes = head
	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := getKthFromEnd(nodes, k)
	fmt.Println("test01 success result: 4->5, result:=")
	print(result)
	fmt.Println()
}

func test02() {
	var head = new(ListNode)
	nums := []int{1}
	k := 1
	nodes := head
	insert(nodes, nums)
	nodes = head
	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	result := getKthFromEnd(nodes, k)
	fmt.Println("test01 success result: 1, result:=")
	print(result)
	fmt.Println()
}

func main() {
	test01()
	test02()
}
