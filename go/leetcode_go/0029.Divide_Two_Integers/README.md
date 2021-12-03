### 29. Divide Two Integers
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.

The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.

Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.



Example 1:

	Input: dividend = 10, divisor = 3
	Output: 3
	Explanation: 10/3 = 3.33333.. which is truncated to 3.

Example 2:

	Input: dividend = 7, divisor = -3
	Output: -2
	Explanation: 7/-3 = -2.33333.. which is truncated to -2.

Example 3:

	Input: dividend = 0, divisor = 1
	Output: 0

Example 4:

	Input: dividend = 1, divisor = 1
	Output: 1



Constraints:

* -2^31 <= dividend, divisor <= 2^31 - 1
* divisor != 0

----

### 29. 两数相除
给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2



示例 1:

	输入: dividend = 10, divisor = 3
	输出: 3
	解释: 10/3 = truncate(3.33333..) = truncate(3) = 3

示例 2:

	输入: dividend = 7, divisor = -3
	输出: -2
	解释: 7/-3 = truncate(-2.33333..) = -2



提示：

* 被除数和除数均为 32 位有符号整数。
* 除数不为 0。
* 假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−2^31,  2^31 − 1]。本题中，如果除法结果溢出，则返回 2^31 − 1。

