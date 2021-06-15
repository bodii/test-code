### 237. Delete Node in a Linked List
Write a function to delete a node in a singly-linked list. You will not be given access to the head of the list, instead you will be given access to the node to be deleted directly.

It is guaranteed that the node to be deleted is not a tail node in the list.



Example 1:

	Input: head = [4,5,1,9], node = 5
	Output: [4,1,9]
	Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.

Example 2:

	Input: head = [4,5,1,9], node = 1
	Output: [4,5,9]
	Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.

Example 3:

	Input: head = [1,2,3,4], node = 3
	Output: [1,2,4]

Example 4:

	Input: head = [0,1], node = 0
	Output: [1]

Example 5:

	Input: head = [-3,5,-99], node = -3
	Output: [5,-99]



Constraints:

* The number of the nodes in the given list is in the range [2, 1000].
* -1000 <= Node.val <= 1000
* The value of each node in the list is unique.
* The node to be deleted is in the list and is not a tail node

----

### 237. 删除链表中的节点
请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点。传入函数的唯一参数为 要被删除的节点 。



现有一个链表 -- head = [4,5,1,9]，它可以表示为:



示例 1：

	输入：head = [4,5,1,9], node = 5
	输出：[4,1,9]
	解释：给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.

示例 2：

	输入：head = [4,5,1,9], node = 1
	输出：[4,5,9]
	解释：给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.



提示：

* 链表至少包含两个节点。
* 链表中所有节点的值都是唯一的。
* 给定的节点为非末尾节点并且一定是链表中的一个有效节点。
* 不要从你的函数中返回任何结果。
