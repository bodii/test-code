### 1957. Delete Characters to Make Fancy String
A fancy string is a string where no three consecutive characters are equal.

Given a string s, delete the minimum possible number of characters from s to make it fancy.

Return the final string after the deletion. It can be shown that the answer will always be unique.



Example 1:

	Input: s = "leeetcode"
	Output: "leetcode"
	Explanation:
	Remove an 'e' from the first group of 'e's to create "leetcode".
	No three consecutive characters are equal, so return "leetcode".

Example 2:

	Input: s = "aaabaaaa"
	Output: "aabaa"
	Explanation:
	Remove an 'a' from the first group of 'a's to create "aabaaaa".
	Remove two 'a's from the second group of 'a's to create "aabaa".
	No three consecutive characters are equal, so return "aabaa".

Example 3:

	Input: s = "aab"
	Output: "aab"
	Explanation: No three consecutive characters are equal, so return "aab".



Constraints:

* 1 <= s.length <= 10^5
* s consists only of lowercase English letters.

----

### 1957. 删除字符使字符串变好
一个字符串如果没有 三个连续 相同字符，那么它就是一个 好字符串 。

给你一个字符串 s ，请你从 s 删除 最少 的字符，使它变成一个 好字符串 。

请你返回删除后的字符串。题目数据保证答案总是 唯一的 。



示例 1：

	输入：s = "leeetcode"
	输出："leetcode"
	解释：
	从第一组 'e' 里面删除一个 'e' ，得到 "leetcode" 。
	没有连续三个相同字符，所以返回 "leetcode" 。

示例 2：

	输入：s = "aaabaaaa"
	输出："aabaa"
	解释：
	从第一组 'a' 里面删除一个 'a' ，得到 "aabaaaa" 。
	从第二组 'a' 里面删除两个 'a' ，得到 "aabaa" 。
	没有连续三个相同字符，所以返回 "aabaa" 。

示例 3：

	输入：s = "aab"
	输出："aab"
	解释：没有连续三个相同字符，所以返回 "aab" 。



提示：

* 1 <= s.length <= 10^5
* s 只包含小写英文字母。

