### 1961. Check If String Is a Prefix of Array
Given a string s and an array of strings words, determine whether s is a prefix string of words.

A string s is a prefix string of words if s can be made by concatenating the first k strings in words for some positive k no larger than words.length.

Return true if s is a prefix string of words, or false otherwise.



Example 1:

	Input: s = "iloveleetcode", words = ["i","love","leetcode","apples"]
	Output: true
	Explanation:
	s can be made by concatenating "i", "love", and "leetcode" together.

Example 2:

	Input: s = "iloveleetcode", words = ["apples","i","love","leetcode"]
	Output: false
	Explanation:
	It is impossible to make s using a prefix of arr.



Constraints:

* 1 <= words.length <= 100
* 1 <= words[i].length <= 20
* 1 <= s.length <= 1000
* words[i] and s consist of only lowercase English letters.

----

### 1961. 检查字符串是否为数组前缀
给你一个字符串 s 和一个字符串数组 words ，请你判断 s 是否为 words 的 前缀字符串 。

字符串 s 要成为 words 的 前缀字符串 ，需要满足：s 可以由 words 中的前 k（k 为 正数 ）个字符串按顺序相连得到，且 k 不超过 words.length 。

如果 s 是 words 的 前缀字符串 ，返回 true ；否则，返回 false 。



示例 1：

	输入：s = "iloveleetcode", words = ["i","love","leetcode","apples"]
	输出：true
	解释：
	s 可以由 "i"、"love" 和 "leetcode" 相连得到。

示例 2：

	输入：s = "iloveleetcode", words = ["apples","i","love","leetcode"]
	输出：false
	解释：
	数组的前缀相连无法得到 s 。



提示：

* 1 <= words.length <= 100
* 1 <= words[i].length <= 20
* 1 <= s.length <= 1000
* words[i] 和 s 仅由小写英文字母组成

