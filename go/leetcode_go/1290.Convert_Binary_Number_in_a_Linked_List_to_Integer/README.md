### 1290. Convert Binary Number in a Linked List to Integer
Given head which is a reference node to a singly-linked list. The value of each node in the linked list is either 0 or 1. The linked list holds the binary representation of a number.

Return the decimal value of the number in the linked list.



Example 1:
	Input: head = [1,0,1]
	Output: 5
	Explanation: (101) in base 2 = (5) in base 10

Example 2:

	Input: head = [0]
	Output: 0

Example 3:

	Input: head = [1]
	Output: 1

Example 4:

	Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
	Output: 18880

Example 5:

	Input: head = [0,0]
	Output: 0

 

Constraints:

* The Linked List is not empty.
* Number of nodes will not exceed 30.
* Each node's value is either 0 or 1.

----

### 1290. 二进制链表转整数
给你一个单链表的引用结点 head。链表中每个结点的值不是 0 就是 1。已知此链表是一个整数数字的二进制表示形式。

请你返回该链表所表示数字的 十进制值 。



示例 1：
	输入：head = [1,0,1]
	输出：5
	解释：二进制数 (101) 转化为十进制数 (5)

示例 2：

	输入：head = [0]
	输出：0

示例 3：

	输入：head = [1]
	输出：1

示例 4：

	输入：head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
	输出：18880

示例 5：

	输入：head = [0,0]
	输出：0

 
提示：

* 链表不为空。
* 链表的结点总数不超过 30。
* 每个结点的值不是 0 就是 1。

