### 1668. Maximum Repeating Substring
For a string sequence, a string word is k-repeating if word concatenated k times is a substring of sequence. The word's maximum k-repeating value is the highest value k where word is k-repeating in sequence. If word is not a substring of sequence, word's maximum k-repeating value is 0.

Given strings sequence and word, return the maximum k-repeating value of word in sequence.

Example 1:

	Input: sequence = "ababc", word = "ab"
	Output: 2
	Explanation: "abab" is a substring in "ababc".

Example 2:

	Input: sequence = "ababc", word = "ba"
	Output: 1
	Explanation: "ba" is a substring in "ababc". "baba" is not a substring in "ababc".

Example 3:

	Input: sequence = "ababc", word = "ac"
	Output: 0
	Explanation: "ac" is not a substring in "ababc".



Constraints:

* 1 <= sequence.length <= 100
* 1 <= word.length <= 100
* sequence and word contains only lowercase English letters.

----

### 1668. 最大重复子字符串
给你一个字符串 sequence ，如果字符串 word 连续重复 k 次形成的字符串是 sequence 的一个子字符串，那么单词 word 的 重复值为 k 。单词 word 的 最大重复值 是单词 word 在 sequence 中最大的重复值。如果 word 不是 sequence 的子串，那么重复值 k 为 0 。

给你一个字符串 sequence 和 word ，请你返回 最大重复值 k 。



示例 1：

	输入：sequence = "ababc", word = "ab"
	输出：2
	解释："abab" 是 "ababc" 的子字符串。

示例 2：

	输入：sequence = "ababc", word = "ba"
	输出：1
	解释："ba" 是 "ababc" 的子字符串，但 "baba" 不是 "ababc" 的子字符串。

示例 3：

	输入：sequence = "ababc", word = "ac"
	输出：0
	解释："ac" 不是 "ababc" 的子字符串。



提示：

* 1 <= sequence.length <= 100
* 1 <= word.length <= 100
* sequence 和 word 都只包含小写英文字母。

