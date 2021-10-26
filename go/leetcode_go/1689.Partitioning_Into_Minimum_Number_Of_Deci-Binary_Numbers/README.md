### 1689. Partitioning Into Minimum Number Of Deci-Binary Numbers
A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros. For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary numbers needed so that they sum up to n.



Example 1:

	Input: n = "32"
	Output: 3
	Explanation: 10 + 11 + 11 = 32

Example 2:

	Input: n = "82734"
	Output: 8

Example 3:

	Input: n = "27346209830709182346"
	Output: 9

Constraints:
* 1 <= n.length <= 10^5
* n consists of only digits.
* n does not contain any leading zeros and represents a positive integer.

----
### 1689. 十-二进制数的最少数目
如果一个十进制数字不含任何前导零，且每一位上的数字不是 0 就是 1 ，那么该数字就是一个 十-二进制数 。例如，101 和 1100 都是 十-二进制数，而 112 和 3001 不是。

给你一个表示十进制整数的字符串 n ，返回和为 n 的 十-二进制数 的最少数目。



示例 1：

	输入：n = "32"
	输出：3
	解释：10 + 11 + 11 = 32

示例 2：

	输入：n = "82734"
	输出：8

示例 3：

	输入：n = "27346209830709182346"
	输出：9



提示：

* 1 <= n.length <= 105
* n 仅由数字组成
* n 不含任何前导零并总是表示正整数

