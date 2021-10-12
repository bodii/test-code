### 104. Maximum Depth of Binary Tree
Given the root of a binary tree, return its maximum depth.

A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.



Example 1:

	Input: root = [3,9,20,null,null,15,7]
	Output: 3

Example 2:

	Input: root = [1,null,2]
	Output: 2

Example 3:

	Input: root = []
	Output: 0

Example 4:

	Input: root = [0]
	Output: 1



Constraints:

* The number of nodes in the tree is in the range [0, 10^4].
* -100 <= Node.val <= 100

----

### 104. 二叉树的最大深度
给定一个二叉树，找出其最大深度。

二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。

说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

		3
	   / \
	  9  20
		/  \
	   15   7

返回它的最大深度 3 。

