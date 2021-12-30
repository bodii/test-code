### 868. Binary Gap
Given a positive integer n, find and return the longest distance between any two adjacent 1's in the binary representation of n. If there are no two adjacent 1's, return 0.

Two 1's are adjacent if there are only 0's separating them (possibly no 0's). The distance between two 1's is the absolute difference between their bit positions. For example, the two 1's in "1001" have a distance of 3.



Example 1:

	Input: n = 22
	Output: 2
	Explanation: 22 in binary is "10110".
	The first adjacent pair of 1's is "10110" with a distance of 2.
	The second adjacent pair of 1's is "10110" with a distance of 1.
	The answer is the largest of these two distances, which is 2.
	Note that "10110" is not a valid pair since there is a 1 separating the two 1's underlined.

Example 2:

	Input: n = 8
	Output: 0
	Explanation: 8 in binary is "1000".
	There are not any adjacent pairs of 1's in the binary representation of 8, so we return 0.

Example 3:

	Input: n = 5
	Output: 2
	Explanation: 5 in binary is "101".



Constraints:

* 1 <= n <= 10^9

----

### 868. 二进制间距
给定一个正整数 n，找到并返回 n 的二进制表示中两个 相邻 1 之间的 最长距离 。如果不存在两个相邻的 1，返回 0 。

如果只有 0 将两个 1 分隔开（可能不存在 0 ），则认为这两个 1 彼此 相邻 。两个 1 之间的距离是它们的二进制表示中位置的绝对差。例如，"1001" 中的两个 1 的距离为 3 。



示例 1：

	输入：n = 22
	输出：2
	解释：
	22 的二进制是 "10110" 。
	在 22 的二进制表示中，有三个 1，组成两对相邻的 1 。
	第一对相邻的 1 中，两个 1 之间的距离为 2 。
	第二对相邻的 1 中，两个 1 之间的距离为 1 。
	答案取两个距离之中最大的，也就是 2 。

示例 2：

	输入：n = 5
	输出：2
	解释：
	5 的二进制是 "101" 。

示例 3：

	输入：n = 6
	输出：1
	解释：
	6 的二进制是 "110" 。

示例 4：

	输入：n = 8
	输出：0
	解释：
	8 的二进制是 "1000" 。
	在 8 的二进制表示中没有相邻的两个 1，所以返回 0 。

示例 5：

	输入：n = 1
	输出：0



提示：

* 1 <= N <= 10^9

