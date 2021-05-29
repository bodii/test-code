### 509. Fibonacci Number
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.

Given n, calculate F(n).



Example 1:

	Input: n = 2
	Output: 1
	Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

Example 2:

	Input: n = 3
	Output: 2
	Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

Example 3:

	Input: n = 4
	Output: 3
	Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.



Constraints:

*    0 <= n <= 30

----

### 509. 斐波那契数
斐波那契数，通常用 F(n) 表示，形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：

F(0) = 0，F(1) = 1
F(n) = F(n - 1) + F(n - 2)，其中 n > 1

给你 n ，请计算 F(n) 。



示例 1：

	输入：2
	输出：1
	解释：F(2) = F(1) + F(0) = 1 + 0 = 1

示例 2：

	输入：3
	输出：2
	解释：F(3) = F(2) + F(1) = 1 + 1 = 2

示例 3：

	输入：4
	输出：3
	解释：F(4) = F(3) + F(2) = 2 + 1 = 3



提示：

* 0 <= n <= 30

