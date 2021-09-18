### 720. Longest Word in Dictionary
Given an array of strings words representing an English Dictionary, return the longest word in words that can be built one character at a time by other words in words.

If there is more than one possible answer, return the longest word with the smallest lexicographical order. If there is no answer, return the empty string.



Example 1:

	Input: words = ["w","wo","wor","worl","world"]
	Output: "world"
	Explanation: The word "world" can be built one character at a time by "w", "wo", "wor", and "worl".

Example 2:

	Input: words = ["a","banana","app","appl","ap","apply","apple"]
	Output: "apple"
	Explanation: Both "apply" and "apple" can be built from other words in the dictionary. However, "apple" is lexicographically smaller than "apply".



Constraints:

* 1 <= words.length <= 1000
* 1 <= words[i].length <= 30
* ords[i] consists of lowercase English letters.

----
### 720. 词典中最长的单词
给出一个字符串数组words组成的一本英语词典。从中找出最长的一个单词，该单词是由words词典中其他单词逐步添加一个字母组成。若其中有多个可行的答案，则返回答案中字典序最小的单词。

若无答案，则返回空字符串。



示例 1：

	输入：
	words = ["w","wo","wor","worl", "world"]
	输出："world"
	解释：
	单词"world"可由"w", "wo", "wor", 和 "worl"添加一个字母组成。

示例 2：

	输入：
	words = ["a", "banana", "app", "appl", "ap", "apply", "apple"]
	输出："apple"
	解释：
	"apply"和"apple"都能由词典中的单词组成。但是"apple"的字典序小于"apply"。



提示：

* 所有输入的字符串都只包含小写字母。
* words数组长度范围为[1,1000]。
* words[i]的长度范围为[1,30]。

