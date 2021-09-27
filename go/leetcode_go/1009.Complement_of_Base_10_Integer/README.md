### 1009. Complement of Base 10 Integer
The complement of an integer is the integer you get when you flip all the 0's to 1's and all the 1's to 0's in its binary representation.

* For example, The integer 5 is "101" in binary and its complement is "010" which is the integer 2.

Given an integer n, return its complement.



Example 1:

	Input: n = 5
	Output: 2
	Explanation: 5 is "101" in binary, with complement "010" in binary, which is 2 in base-10.

Example 2:

	Input: n = 7
	Output: 0
	Explanation: 7 is "111" in binary, with complement "000" in binary, which is 0 in base-10.

Example 3:

	Input: n = 10
	Output: 5
	Explanation: 10 is "1010" in binary, with complement "0101" in binary, which is 5 in base-10.



Constraints:

* 0 <= n < 10^9^



Note: This question is the same as 476: https://leetcode.com/problems/number-complement/

----

### 1009. 十进制整数的反码
每个非负整数 N 都有其二进制表示。例如， 5 可以被表示为二进制 "101"，11 可以用二进制 "1011" 表示，依此类推。注意，除 N = 0 外，任何二进制表示中都不含前导零。

二进制的反码表示是将每个 1 改为 0 且每个 0 变为 1。例如，二进制数 "101" 的二进制反码为 "010"。

给你一个十进制数 N，请你返回其二进制表示的反码所对应的十进制整数。



示例 1：

	输入：5
	输出：2
	解释：5 的二进制表示为 "101"，其二进制反码为 "010"，也就是十进制中的 2 。

示例 2：

	输入：7
	输出：0
	解释：7 的二进制表示为 "111"，其二进制反码为 "000"，也就是十进制中的 0 。

示例 3：

	输入：10
	输出：5
	解释：10 的二进制表示为 "1010"，其二进制反码为 "0101"，也就是十进制中的 5 。



提示：

* 0 <= N < 10^9
* 本题与 476：https://leetcode-cn.com/problems/number-complement/ 相同

