### 1967. Number of Strings That Appear as Substrings in Word
Given an array of strings patterns and a string word, return the number of strings in patterns that exist as a substring in word.

A substring is a contiguous sequence of characters within a string.



Example 1:

	Input: patterns = ["a","abc","bc","d"], word = "abc"
	Output: 3
	Explanation:
	- "a" appears as a substring in "abc".
	- "abc" appears as a substring in "abc".
	- "bc" appears as a substring in "abc".
	- "d" does not appear as a substring in "abc".
	3 of the strings in patterns appear as a substring in word.

Example 2:

	Input: patterns = ["a","b","c"], word = "aaaaabbbbb"
	Output: 2
	Explanation:
	- "a" appears as a substring in "aaaaabbbbb".
	- "b" appears as a substring in "aaaaabbbbb".
	- "c" does not appear as a substring in "aaaaabbbbb".
	2 of the strings in patterns appear as a substring in word.

Example 3:

	Input: patterns = ["a","a","a"], word = "ab"
	Output: 3
	Explanation: Each of the patterns appears as a substring in word "ab".



Constraints:

* 1 <= patterns.length <= 100
* 1 <= patterns[i].length <= 100
* 1 <= word.length <= 100
* patterns[i] and word consist of lowercase English letters.

----

### 1967. 作为子字符串出现在单词中的字符串数目
给你一个字符串数组 patterns 和一个字符串 word ，统计 patterns 中有多少个字符串是 word 的子字符串。返回字符串数目。

子字符串 是字符串中的一个连续字符序列。



示例 1：

	输入：patterns = ["a","abc","bc","d"], word = "abc"
	输出：3
	解释：
	- "a" 是 "abc" 的子字符串。
	- "abc" 是 "abc" 的子字符串。
	- "bc" 是 "abc" 的子字符串。
	- "d" 不是 "abc" 的子字符串。
	patterns 中有 3 个字符串作为子字符串出现在 word 中。

示例 2：

	输入：patterns = ["a","b","c"], word = "aaaaabbbbb"
	输出：2
	解释：
	- "a" 是 "aaaaabbbbb" 的子字符串。
	- "b" 是 "aaaaabbbbb" 的子字符串。
	- "c" 不是 "aaaaabbbbb" 的字符串。
	patterns 中有 2 个字符串作为子字符串出现在 word 中。

示例 3：

	输入：patterns = ["a","a","a"], word = "ab"
	输出：3
	解释：patterns 中的每个字符串都作为子字符串出现在 word "ab" 中。



提示：

* 1 <= patterns.length <= 100
* 1 <= patterns[i].length <= 100
* 1 <= word.length <= 100
* patterns[i] 和 word 由小写英文字母组成

