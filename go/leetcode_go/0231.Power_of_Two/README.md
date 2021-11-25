### 231. Power of Two
Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == $2^x$.



Example 1:

	Input: n = 1
	Output: true
	Explanation: 20 = 1

Example 2:

	Input: n = 16
	Output: true
	Explanation: 24 = 16

Example 3:

	Input: n = 3
	Output: false

Example 4:

	Input: n = 4
	Output: true

Example 5:

	Input: n = 5
	Output: false



Constraints:

* -2<sup>31</sup> <= n <= 2<sup>31</sup> - 1


Follow up: Could you solve it without loops/recursion?

---- 
### 231. 2 的幂

给你一个整数 n，请你判断该整数是否是 2 的幂次方。如果是，返回 true ；否则，返回 false 。

如果存在一个整数 x 使得 n == $2^x$ ，则认为 n 是 2 的幂次方。

 

示例 1：

	输入：n = 1
	输出：true
	解释：20 = 1

示例 2：

	输入：n = 16
	输出：true
	解释：24 = 16

示例 3：

	输入：n = 3
	输出：false

示例 4：

	输入：n = 4
	输出：true

示例 5：

	输入：n = 5
	输出：false

 

提示：

* -231 <= n <= 231 - 1

 

进阶：你能够不使用循环/递归解决此问题吗？