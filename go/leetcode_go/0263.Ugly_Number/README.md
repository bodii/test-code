### 263. Ugly Number
An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return true if n is an ugly number.



Example 1:

	Input: n = 6
	Output: true
	Explanation: 6 = 2 × 3

Example 2:

	Input: n = 8
	Output: true
	Explanation: 8 = 2 × 2 × 2

Example 3:

	Input: n = 14
	Output: false
	Explanation: 14 is not ugly since it includes the prime factor 7.

Example 4:

	Input: n = 1
	Output: true
	Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.



Constraints:

    -2^31 <= n <= 2^31 - 1

----

### 263. 丑数
给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。

丑数 就是只包含质因数 2、3 和/或 5 的正整数。



示例 1：

	输入：n = 6
	输出：true
	解释：6 = 2 × 3

示例 2：

	输入：n = 8
	输出：true
	解释：8 = 2 × 2 × 2

示例 3：

	输入：n = 14
	输出：false
	解释：14 不是丑数，因为它包含了另外一个质因数 7 。

示例 4：

	输入：n = 1
	输出：true
	解释：1 通常被视为丑数。



提示：

    -2^31 <= n <= 2^31 - 1

