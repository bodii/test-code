### 1837. Sum of Digits in Base K
Given an integer n (in base 10) and a base k, return the sum of the digits of n after converting n from base 10 to base k.

After converting, each digit should be interpreted as a base 10 number, and the sum should be returned in base 10.



Example 1:

	Input: n = 34, k = 6
	Output: 9
	Explanation: 34 (base 10) expressed in base 6 is 54. 5 + 4 = 9.

Example 2:

	Input: n = 10, k = 10
	Output: 1
	Explanation: n is already in base 10. 1 + 0 = 1.



Constraints:

* 1 <= n <= 100
* 2 <= k <= 10

----

### 1837. K 进制表示下的各位数字总和
给你一个整数 n（10 进制）和一个基数 k ，请你将 n 从 10 进制表示转换为 k 进制表示，计算并返回转换后各位数字的 总和 。

转换后，各位数字应当视作是 10 进制数字，且它们的总和也应当按 10 进制表示返回。



示例 1：

	输入：n = 34, k = 6
	输出：9
	解释：34 (10 进制) 在 6 进制下表示为 54 。5 + 4 = 9 。

示例 2：

	输入：n = 10, k = 10
	输出：1
	解释：n 本身就是 10 进制。 1 + 0 = 1 。



提示：

* 1 <= n <= 100
* 2 <= k <= 10

