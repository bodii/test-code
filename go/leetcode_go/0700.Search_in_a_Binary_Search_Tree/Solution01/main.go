package main

import "fmt"

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
func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}

	return searchBST(root.Right, val)
}

func test01() {
	root := []int{4, 2, 7, 1, 3}
	val := 2
	successResult := []int{2, 1, 3}
	rootTree := &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree sum success result:", successResult)
	tree2 := &BinaryTree{root: searchBST(rootTree.root, val)}
	tree2.printTree()
	fmt.Printf("\n\n")
}

func test02() {
	root := []int{4, 2, 7, 1, 3}
	val := 5
	successResult := []int{}
	rootTree := &BinaryTree{}

	fmt.Println("test02: ")
	fmt.Println("root :=", root)

	fmt.Println("create tree1:")
	for _, n := range root {
		rootTree.add(n)
	}
	rootTree.printTree()

	fmt.Println("tree sum success result:", successResult)
	tree2 := &BinaryTree{root: searchBST(rootTree.root, val)}
	tree2.printTree()
	fmt.Printf("\n\n")
}

func main() {
	test01()
	test02()
}
