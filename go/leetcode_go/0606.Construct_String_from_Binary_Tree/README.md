### 606. Construct String from Binary Tree
Given the root of a binary tree, construct a string consisting of parenthesis and integers from a binary tree with the preorder traversal way, and return it.

Omit all the empty parenthesis pairs that do not affect the one-to-one mapping relationship between the string and the original binary tree.



Example 1:

	Input: root = [1,2,3,4]
	Output: "1(2(4))(3)"
	Explanation: Originally, it needs to be "1(2(4)())(3()())", but you need to omit all the unnecessary empty parenthesis pairs. And it will be "1(2(4))(3)"

Example 2:

	Input: root = [1,2,3,null,4]
	Output: "1(2()(4))(3)"
	Explanation: Almost the same as the first example, except we cannot omit the first parenthesis pair to break the one-to-one mapping relationship between the input and the output.



Constraints:

* The number of nodes in the tree is in the range [1, 10^4].
* -1000 <= Node.val <= 1000

----

### 606. 根据二叉树创建字符串
你需要采用前序遍历的方式，将一个二叉树转换成一个由括号和整数组成的字符串。

空节点则用一对空括号 "()" 表示。而且你需要省略所有不影响字符串与原始二叉树之间的一对一映射关系的空括号对。

示例 1:

	输入: 二叉树: [1,2,3,4]
		   1
		 /   \
		2     3
	   /
	  4

	输出: "1(2(4))(3)"

解释: 原本将是“1(2(4)())(3())”，
在你省略所有不必要的空括号对之后，
它将是“1(2(4))(3)”。

示例 2:

	输入: 二叉树: [1,2,3,null,4]
		   1
		 /   \
		2     3
		 \
		  4

	输出: "1(2()(4))(3)"

解释: 和第一个示例相似，
除了我们不能省略第一个对括号来中断输入和输出之间的一对一映射关系。

