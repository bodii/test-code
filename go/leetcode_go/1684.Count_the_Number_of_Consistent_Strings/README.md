### 1684. Count the Number of Consistent Strings
You are given a string allowed consisting of distinct characters and an array of strings words. A string is consistent if all characters in the string appear in the string allowed.

Return the number of consistent strings in the array words.



Example 1:

	Input: allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
	Output: 2
	Explanation: Strings "aaab" and "baa" are consistent since they only contain characters 'a' and 'b'.

Example 2:

	Input: allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
	Output: 7
	Explanation: All strings are consistent.

Example 3:

	Input: allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
	Output: 4
	Explanation: Strings "cc", "acd", "ac", and "d" are consistent.



Constraints:

* 1 <= words.length <= 104
* 1 <= allowed.length <= 26
* 1 <= words[i].length <= 10
* The characters in allowed are distinct.
* words[i] and allowed contain only lowercase English letters.

----

### 1684. 统计一致字符串的数目
给你一个由不同字符组成的字符串 allowed 和一个字符串数组 words 。如果一个字符串的每一个字符都在 allowed 中，就称这个字符串是 一致字符串 。

请你返回 words 数组中 一致字符串 的数目。



示例 1：

	输入：allowed = "ab", words = ["ad","bd","aaab","baa","badab"]
	输出：2
	解释：字符串 "aaab" 和 "baa" 都是一致字符串，因为它们只包含字符 'a' 和 'b' 。

示例 2：

	输入：allowed = "abc", words = ["a","b","c","ab","ac","bc","abc"]
	输出：7
	解释：所有字符串都是一致的。

示例 3：

	输入：allowed = "cad", words = ["cc","acd","b","ba","bac","bad","ac","d"]
	输出：4
	解释：字符串 "cc"，"acd"，"ac" 和 "d" 是一致字符串。



提示：

* 1 <= words.length <= 104
* 1 <= allowed.length <= 26
* 1 <= words[i].length <= 10
* allowed 中的字符 互不相同 。
* words[i] 和 allowed 只包含小写英文字母。
