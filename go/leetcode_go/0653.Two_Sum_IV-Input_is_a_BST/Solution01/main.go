package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BinaryTree struct {
	root  *TreeNode
	count int
}

func (tn *TreeNode) addNode(newNode *TreeNode) {
	if tn.Val > newNode.Val {
		if tn.Left == nil {
			tn.Left = newNode
		} else {
			tn.Left.addNode(newNode)
		}
	} else {
		if tn.Right == nil {
			tn.Right = newNode
		} else {
			tn.Right.addNode(newNode)
		}
	}
}

func (bt *BinaryTree) add(val int) {
	var newNode TreeNode
	newNode.Val = val
	if bt.root == nil {
		bt.root = &newNode
	} else {
		bt.root.addNode(&newNode)
	}
	bt.count++
}

func (tn *TreeNode) printNode() {
	fmt.Print(" ->val: ", tn.Val)
	if tn.Left != nil {
		tn.Left.printNode()
		fmt.Print(" l: ", tn.Left.Val)
	}

	if tn.Right != nil {
		tn.Right.printNode()
		fmt.Print(" r: ", tn.Right.Val)
	}
	fmt.Println()
}

func (bt *BinaryTree) printTree() {
	if bt.root == nil {
		return
	}

	bt.root.printNode()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
	values := make(map[int]int)

	index := 1
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}

		values[root.Val] = index
		index++

		dfs(root.Left)
		dfs(root.Right)
	}

	dfs(root)

	for v := range values {
		if values[k-v] > 0 && values[k-v] != values[v] {
			return true
		}
	}

	return false
}

func test01() {
	root := []int{5, 3, 6, 2, 4, 7}
	k := 9
	successResult := true
	rootTree := &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", findTarget(rootTree.root, k))
	fmt.Printf("\n\n")
}

func test02() {
	root := []int{5, 3, 6, 2, 4, 7}
	k := 28
	successResult := false
	rootTree := &BinaryTree{}

	fmt.Println("test02: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree postorder Traversal success result:", successResult)
	fmt.Println("tree postorder Traversal result:", findTarget(rootTree.root, k))
	fmt.Printf("\n\n")
}

func main() {
	test01()
	test02()
}
