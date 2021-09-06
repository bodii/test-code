### 1446. Consecutive Characters
The power of the string is the maximum length of a non-empty substring that contains only one unique character.

Given a string s, return the power of s.



Example 1:

	Input: s = "leetcode"
	Output: 2
	Explanation: The substring "ee" is of length 2 with the character 'e' only.

Example 2:

	Input: s = "abbcccddddeeeeedcba"
	Output: 5
	Explanation: The substring "eeeee" is of length 5 with the character 'e' only.

Example 3:

	Input: s = "triplepillooooow"
	Output: 5

Example 4:

	Input: s = "hooraaaaaaaaaaay"
	Output: 11

Example 5:

	Input: s = "tourist"
	Output: 1



Constraints:

* 1 <= s.length <= 500
* s consists of only lowercase English letters.


----

### 1446. 连续字符
给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。

请你返回字符串的能量。



示例 1：

	输入：s = "leetcode"
	输出：2
	解释：子字符串 "ee" 长度为 2 ，只包含字符 'e' 。

示例 2：

	输入：s = "abbcccddddeeeeedcba"
	输出：5
	解释：子字符串 "eeeee" 长度为 5 ，只包含字符 'e' 。

示例 3：

	输入：s = "triplepillooooow"
	输出：5

示例 4：

	输入：s = "hooraaaaaaaaaaay"
	输出：11

示例 5：

	输入：s = "tourist"
	输出：1



提示：

* 1 <= s.length <= 500
* s 只包含小写英文字母。

