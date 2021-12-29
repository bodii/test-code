### 2119. A Number After a Double Reversal
Reversing an integer means to reverse all its digits.

* For example, reversing 2021 gives 1202. Reversing 12300 gives 321 as the leading zeros are not retained.

Given an integer num, reverse num to get reversed1, then reverse reversed1 to get reversed2. Return true if reversed2 equals num. Otherwise return false.



Example 1:

	Input: num = 526
	Output: true
	Explanation: Reverse num to get 625, then reverse 625 to get 526, which equals num.

Example 2:

	Input: num = 1800
	Output: false
	Explanation: Reverse num to get 81, then reverse 81 to get 18, which does not equal num.

Example 3:

	Input: num = 0
	Output: true
	Explanation: Reverse num to get 0, then reverse 0 to get 0, which equals num.



Constraints:

* 0 <= num <= 10^6

----

### 2119. 反转两次的数字
反转 一个整数意味着倒置它的所有位。

* 例如，反转 2021 得到 1202 。反转 12300 得到 321 ，不保留前导零 。

给你一个整数 num ，反转 num 得到 reversed1 ，接着反转 reversed1 得到 reversed2 。如果 reversed2 等于 num ，返回 true ；否则，返回 false 。



示例 1：

	输入：num = 526
	输出：true
	解释：反转 num 得到 625 ，接着反转 625 得到 526 ，等于 num 。

示例 2：

	输入：num = 1800
	输出：false
	解释：反转 num 得到 81 ，接着反转 81 得到 18 ，不等于 num 。

示例 3：

	输入：num = 0
	输出：true
	解释：反转 num 得到 0 ，接着反转 0 得到 0 ，等于 num 。



提示：

* 0 <= num <= 10^6

