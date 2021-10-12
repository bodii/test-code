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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root2 == nil {
		return root1
	}

	if root1 != nil {
		root1.Val += root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	} else {
		return root2
	}

	return root1
}

func test01() {
	root1 := []int{1, 3, 2, 5}
	root2 := []int{2, 1, 3, 4, 7}
	successResult := []int{3, 4, 5, 5, 4, 7}
	rootTree1, rootTree2 := &BinaryTree{}, &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root1 :=", root1)
	fmt.Println("root2 :=", root2)

	fmt.Println("create tree1:")
	for _, n := range root1 {
		rootTree1.add(n)
	}
	rootTree1.printTree()

	fmt.Println("create tree2:")
	for _, n := range root2 {
		rootTree2.add(n)
	}
	rootTree2.printTree()

	fmt.Println("tree sum success result:", successResult)
	rootTree := &BinaryTree{}
	rootTree.root = mergeTrees(rootTree1.root, rootTree2.root)
	rootTree.printTree()
	fmt.Printf("\n\n")
}

func test02() {
	root1 := []int{1}
	root2 := []int{1, 2}
	successResult := []int{2, 2}
	rootTree1, rootTree2 := &BinaryTree{}, &BinaryTree{}

	fmt.Println("test01: ")
	fmt.Println("root1 :=", root1)
	fmt.Println("root2 :=", root2)

	fmt.Println("create tree1:")
	for _, n := range root1 {
		rootTree1.add(n)
	}
	rootTree1.printTree()

	fmt.Println("create tree2:")
	for _, n := range root2 {
		rootTree2.add(n)
	}
	rootTree2.printTree()

	fmt.Println("tree sum success result:", successResult)
	rootTree := &BinaryTree{}
	rootTree.root = mergeTrees(rootTree1.root, rootTree2.root)
	rootTree.printTree()
	fmt.Printf("\n\n")
}

func main() {
	test01()
	test02()
}
