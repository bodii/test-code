### 521. Longest Uncommon Subsequence I
Given two strings a and b, return the length of the longest uncommon subsequence between a and b. If the longest uncommon subsequence does not exist, return -1.

An uncommon subsequence between two strings is a string that is a subsequence of one but not the other.

A subsequence of a string s is a string that can be obtained after deleting any number of characters from s.

* For example, "abc" is a subsequence of "aebdc" because you can delete the underlined characters in "aebdc" to get "abc". Other subsequences of "aebdc" include "aebdc", "aeb", and "" (empty string).



Example 1:

	Input: a = "aba", b = "cdc"
	Output: 3
	Explanation: One longest uncommon subsequence is "aba" because "aba" is a subsequence of "aba" but not "cdc".
	Note that "cdc" is also a longest uncommon subsequence.

Example 2:

	Input: a = "aaa", b = "bbb"
	Output: 3
	Explanation: The longest uncommon subsequences are "aaa" and "bbb".

Example 3:

	Input: a = "aaa", b = "aaa"
	Output: -1
	Explanation: Every subsequence of string a is also a subsequence of string b. Similarly, every subsequence of string b is also a subsequence of string a.



Constraints:

* 1 <= a.length, b.length <= 100
* a and b consist of lower-case English letters.

----

### 521. 最长特殊序列 Ⅰ
给你两个字符串，请你从这两个字符串中找出最长的特殊序列。

「最长特殊序列」定义如下：该序列为某字符串独有的最长子序列（即不能是其他字符串的子序列）。

子序列 可以通过删去字符串中的某些字符实现，但不能改变剩余字符的相对顺序。空序列为所有字符串的子序列，任何字符串为其自身的子序列。

输入为两个字符串，输出最长特殊序列的长度。如果不存在，则返回 -1。



示例 1：

	输入: "aba", "cdc"
	输出: 3
	解释: 最长特殊序列可为 "aba" (或 "cdc")，两者均为自身的子序列且不是对方的子序列。

示例 2：

	输入：a = "aaa", b = "bbb"
	输出：3

示例 3：

	输入：a = "aaa", b = "aaa"
	输出：-1



提示：

* 两个字符串长度均处于区间 [1 - 100] 。
* 字符串中的字符仅含有 'a'~'z' 。

