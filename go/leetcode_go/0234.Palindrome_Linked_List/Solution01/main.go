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
func isPalindrome(head *ListNode) bool {
	nodeVals := make([]int, 0)

	for head != nil {
		nodeVals = append(nodeVals, head.Val)
		head = head.Next
	}

	nodeValsLen := len(nodeVals)
	for i := 0; i < nodeValsLen/2; i++ {
		if nodeVals[i] != nodeVals[nodeValsLen-1-i] {
			return false
		}
	}

	return true
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
	head := new(ListNode)
	l1 := []int{1, 2, 2, 1}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("test01 success result: true")
	fmt.Println("result: ", isPalindrome(nodes))
	fmt.Println()
}

func test02() {
	head := new(ListNode)
	l1 := []int{1, 2}
	nodes := head
	insert(nodes, l1)

	fmt.Println("print:")
	print(nodes)

	fmt.Println()
	fmt.Println("test02 success result: false")
	fmt.Println("result: ", isPalindrome(nodes))
	fmt.Println()
}

func main() {
	test01()
	test02()
}
