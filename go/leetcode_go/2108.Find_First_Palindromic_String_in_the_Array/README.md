### 2108. Find First Palindromic String in the Array
Given an array of strings words, return the first palindromic string in the array. If there is no such string, return an empty string "".

A string is palindromic if it reads the same forward and backward.



Example 1:

	Input: words = ["abc","car","ada","racecar","cool"]
	Output: "ada"
	Explanation: The first string that is palindromic is "ada".
	Note that "racecar" is also palindromic, but it is not the first.

Example 2:

	Input: words = ["notapalindrome","racecar"]
	Output: "racecar"
	Explanation: The first and only string that is palindromic is "racecar".

Example 3:

	Input: words = ["def","ghi"]
	Output: ""
	Explanation: There are no palindromic strings, so the empty string is returned.



Constraints:

* 1 <= words.length <= 100
* 1 <= words[i].length <= 100
* words[i] consists only of lowercase English letters.

----

### 2108. 找出数组中的第一个回文字符串
给你一个字符串数组 words ，找出并返回数组中的 第一个回文字符串 。如果不存在满足要求的字符串，返回一个 空字符串 "" 。

回文字符串 的定义为：如果一个字符串正着读和反着读一样，那么该字符串就是一个 回文字符串 。



示例 1：

	输入：words = ["abc","car","ada","racecar","cool"]
	输出："ada"
	解释：第一个回文字符串是 "ada" 。
	注意，"racecar" 也是回文字符串，但它不是第一个。

示例 2：

	输入：words = ["notapalindrome","racecar"]
	输出："racecar"
	解释：第一个也是唯一一个回文字符串是 "racecar" 。

示例 3：

	输入：words = ["def","ghi"]
	输出：""
	解释：不存在回文字符串，所以返回一个空字符串。



提示：

* 1 <= words.length <= 100
* 1 <= words[i].length <= 100
* words[i] 仅由小写英文字母组成

