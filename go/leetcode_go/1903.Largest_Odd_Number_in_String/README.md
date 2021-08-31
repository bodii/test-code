### 1903. Largest Odd Number in String
You are given a string num, representing a large integer. Return the largest-valued odd integer (as a string) that is a non-empty substring of num, or an empty string "" if no odd integer exists.

A substring is a contiguous sequence of characters within a string.



Example 1:

	Input: num = "52"
	Output: "5"
	Explanation: The only non-empty substrings are "5", "2", and "52". "5" is the only odd number.

Example 2:

	Input: num = "4206"
	Output: ""
	Explanation: There are no odd numbers in "4206".

Example 3:

	Input: num = "35427"
	Output: "35427"
	Explanation: "35427" is already an odd number.



Constraints:

* 1 <= num.length <= 10^5
* num only consists of digits and does not contain any leading zeros.


----

### 1903. 字符串中的最大奇数
给你一个字符串 num ，表示一个大整数。请你在字符串 num 的所有 非空子字符串 中找出 值最大的奇数 ，并以字符串形式返回。如果不存在奇数，则返回一个空字符串 "" 。

子字符串 是字符串中的一个连续的字符序列。



示例 1：

	输入：num = "52"
	输出："5"
	解释：非空子字符串仅有 "5"、"2" 和 "52" 。"5" 是其中唯一的奇数。

示例 2：

	输入：num = "4206"
	输出：""
	解释：在 "4206" 中不存在奇数。

示例 3：

	输入：num = "35427"
	输出："35427"
	解释："35427" 本身就是一个奇数。



提示：

* 1 <= num.length <= 10^5
* num 仅由数字组成且不含前导零

