### 1897. Redistribute Characters to Make All Strings Equal
You are given an array of strings words (0-indexed).

In one operation, pick two distinct indices i and j, where words[i] is a non-empty string, and move any character from words[i] to any position in words[j].

Return true if you can make every string in words equal using any number of operations, and false otherwise.



Example 1:

	Input: words = ["abc","aabc","bc"]
	Output: true
	Explanation: Move the first 'a' in words[1] to the front of words[2],
	to make words[1] = "abc" and words[2] = "abc".
	All the strings are now equal to "abc", so return true.

Example 2:

	Input: words = ["ab","a"]
	Output: false
	Explanation: It is impossible to make all the strings equal using the operation.



Constraints:

* 1 <= words.length <= 100
* 1 <= words[i].length <= 100
* words[i] consists of lowercase English letters.

----

### 1897. 重新分配字符使所有字符串都相等
给你一个字符串数组 words（下标 从 0 开始 计数）。

在一步操作中，需先选出两个 不同 下标 i 和 j，其中 words[i] 是一个非空字符串，接着将 words[i] 中的 任一 字符移动到 words[j] 中的 任一 位置上。

如果执行任意步操作可以使 words 中的每个字符串都相等，返回 true ；否则，返回 false 。



示例 1：

	输入：words = ["abc","aabc","bc"]
	输出：true
	解释：将 words[1] 中的第一个 'a' 移动到 words[2] 的最前面。
	使 words[1] = "abc" 且 words[2] = "abc" 。
	所有字符串都等于 "abc" ，所以返回 true 。

示例 2：

	输入：words = ["ab","a"]
	输出：false
	解释：执行操作无法使所有字符串都相等。



提示：

* 1 <= words.length <= 100
* 1 <= words[i].length <= 100
* words[i] 由小写英文字母组成

