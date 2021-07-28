### 1941. Check if All Characters Have Equal Number of Occurrences
Given a string s, return true if s is a good string, or false otherwise.

A string s is good if all the characters that appear in s have the same number of occurrences (i.e., the same frequency).



Example 1:

	Input: s = "abacbc"
	Output: true
	Explanation: The characters that appear in s are 'a', 'b', and 'c'. All characters occur 2 times in s.

Example 2:

	Input: s = "aaabb"
	Output: false
	Explanation: The characters that appear in s are 'a' and 'b'.
	'a' occurs 3 times while 'b' occurs 2 times, which is not the same number of times.



Constraints:

* 1 <= s.length <= 1000
* s consists of lowercase English letters.

----

### 1941. 检查是否所有字符出现次数相同
给你一个字符串 s ，如果 s 是一个 好 字符串，请你返回 true ，否则请返回 false 。

如果 s 中出现过的 所有 字符的出现次数 相同 ，那么我们称字符串 s 是 好 字符串。



示例 1：

	输入：s = "abacbc"
	输出：true
	解释：s 中出现过的字符为 'a'，'b' 和 'c' 。s 中所有字符均出现 2 次。

示例 2：

	输入：s = "aaabb"
	输出：false
	解释：s 中出现过的字符为 'a' 和 'b' 。
	'a' 出现了 3 次，'b' 出现了 2 次，两者出现次数不同。



提示：

* 1 <= s.length <= 1000
* s 只包含小写英文字母。

