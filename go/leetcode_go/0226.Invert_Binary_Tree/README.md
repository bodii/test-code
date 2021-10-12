### 226. Invert Binary Tree
Given the root of a binary tree, invert the tree, and return its root.



Example 1:
	Input: root = [4,2,7,1,3,6,9]
	Output: [4,7,2,9,6,3,1]

Example 2:

	Input: root = [2,1,3]
	Output: [2,3,1]

Example 3:

	Input: root = []
	Output: []

 

Constraints:

* The number of nodes in the tree is in the range [0, 100].
* -100 <= Node.val <= 100

----

### 226. 翻转二叉树
翻转一棵二叉树。

示例：

输入：

		 4
	   /   \
	  2     7
	 / \   / \
	1   3 6   9

输出：

		 4
	   /   \
	  7     2
	 / \   / \
	9   6 3   1

备注:
这个问题是受到 Max Howell 的 原问题 启发的 ：

> 谷歌：我们90％的工程师使用您编写的软件(Homebrew)，但是您却无法在面试时在白板上写出翻转二叉树这道题，这太糟糕了。

