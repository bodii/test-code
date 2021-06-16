### 1624. Largest Substring Between Two Equal Characters
Given a string s, return the length of the longest substring between two equal characters, excluding the two characters. If there is no such substring return -1.

A substring is a contiguous sequence of characters within a string.



Example 1:

	Input: s = "aa"
	Output: 0
	Explanation: The optimal substring here is an empty substring between the two 'a's.

Example 2:

	Input: s = "abca"
	Output: 2
	Explanation: The optimal substring here is "bc".

Example 3:

	Input: s = "cbzxy"
	Output: -1
	Explanation: There are no characters that appear twice in s.

Example 4:

	Input: s = "cabbac"
	Output: 4
	Explanation: The optimal substring here is "abba". Other non-optimal substrings include "bb" and "".



Constraints:

* 1 <= s.length <= 300
* s contains only lowercase English letters.

----

### 1624. 两个相同字符之间的最长子字符串
给你一个字符串 s，请你返回 两个相同字符之间的最长子字符串的长度 ，计算长度时不含这两个字符。如果不存在这样的子字符串，返回 -1 。

子字符串 是字符串中的一个连续字符序列。



示例 1：

	输入：s = "aa"
	输出：0
	解释：最优的子字符串是两个 'a' 之间的空子字符串。

示例 2：

	输入：s = "abca"
	输出：2
	解释：最优的子字符串是 "bc" 。

示例 3：

	输入：s = "cbzxy"
	输出：-1
	解释：s 中不存在出现出现两次的字符，所以返回 -1 。

示例 4：

	输入：s = "cabbac"
	输出：4
	解释：最优的子字符串是 "abba" ，其他的非最优解包括 "bb" 和 "" 。



提示：

* 1 <= s.length <= 300
* s 只含小写英文字母

