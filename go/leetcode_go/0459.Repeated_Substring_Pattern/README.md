### 459. Repeated Substring Pattern
Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.



Example 1:

	Input: s = "abab"
	Output: true
	Explanation: It is the substring "ab" twice.

Example 2:

	Input: s = "aba"
	Output: false

Example 3:

	Input: s = "abcabcabcabc"
	Output: true
	Explanation: It is the substring "abc" four times or the substring "abcabc" twice.



Constraints:

* 1 <= s.length <= 10^4
* s consists of lowercase English letters.

----

### 459. 重复的子字符串
给定一个非空的字符串，判断它是否可以由它的一个子串重复多次构成。给定的字符串只含有小写英文字母，并且长度不超过10000。

示例 1:

	输入: "abab"

	输出: True

	解释: 可由子字符串 "ab" 重复两次构成。

示例 2:

	输入: "aba"

	输出: False

示例 3:

	输入: "abcabcabcabc"

	输出: True

	解释: 可由子字符串 "abc" 重复四次构成。 (或者子字符串 "abcabc" 重复两次构成。)

