### 1768. Merge Strings Alternately
You are given two strings word1 and word2. Merge the strings by adding letters in alternating order, starting with word1. If a string is longer than the other, append the additional letters onto the end of the merged string.

Return the merged string.



Example 1:

	Input: word1 = "abc", word2 = "pqr"
	Output: "apbqcr"
	Explanation: The merged string will be merged as so:
	word1:  a   b   c
	word2:    p   q   r
	merged: a p b q c r

Example 2:

	Input: word1 = "ab", word2 = "pqrs"
	Output: "apbqrs"
	Explanation: Notice that as word2 is longer, "rs" is appended to the end.
	word1:  a   b
	word2:    p   q   r   s
	merged: a p b q   r   s

Example 3:

	Input: word1 = "abcd", word2 = "pq"
	Output: "apbqcd"
	Explanation: Notice that as word1 is longer, "cd" is appended to the end.
	word1:  a   b   c   d
	word2:    p   q
	merged: a p b q c   d



Constraints:

* 1 <= word1.length, word2.length <= 100
* word1 and word2 consist of lowercase English letters.

----

### 1768. 交替合并字符串
给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。

返回 合并后的字符串 。



示例 1：

	输入：word1 = "abc", word2 = "pqr"
	输出："apbqcr"
	解释：字符串合并情况如下所示：
	word1：  a   b   c
	word2：    p   q   r
	合并后：  a p b q c r

示例 2：

	输入：word1 = "ab", word2 = "pqrs"
	输出："apbqrs"
	解释：注意，word2 比 word1 长，"rs" 需要追加到合并后字符串的末尾。
	word1：  a   b
	word2：    p   q   r   s
	合并后：  a p b q   r   s

示例 3：

	输入：word1 = "abcd", word2 = "pq"
	输出："apbqcd"
	解释：注意，word1 比 word2 长，"cd" 需要追加到合并后字符串的末尾。
	word1：  a   b   c   d
	word2：    p   q
	合并后：  a p b q c   d



提示：

* 1 <= word1.length, word2.length <= 100
* word1 和 word2 由小写英文字母组成

