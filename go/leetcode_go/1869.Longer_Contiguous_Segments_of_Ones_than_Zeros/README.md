### 1869. Longer Contiguous Segments of Ones than Zeros
Given a binary string s, return true if the longest contiguous segment of 1s is strictly longer than the longest contiguous segment of 0s in s. Return false otherwise.

* For example, in s = "110100010" the longest contiguous segment of 1s has length 2, and the longest contiguous segment of 0s has length 3.

Note that if there are no 0s, then the longest contiguous segment of 0s is considered to have length 0. The same applies if there are no 1s.



Example 1:

	Input: s = "1101"
	Output: true
	Explanation:
	The longest contiguous segment of 1s has length 2: "1101"
	The longest contiguous segment of 0s has length 1: "1101"
	The segment of 1s is longer, so return true.

Example 2:

	Input: s = "111000"
	Output: false
	Explanation:
	The longest contiguous segment of 1s has length 3: "111000"
	The longest contiguous segment of 0s has length 3: "111000"
	The segment of 1s is not longer, so return false.

Example 3:

	Input: s = "110100010"
	Output: false
	Explanation:
	The longest contiguous segment of 1s has length 2: "110100010"
	The longest contiguous segment of 0s has length 3: "110100010"
	The segment of 1s is not longer, so return false.



Constraints:

* 1 <= s.length <= 100
* s[i] is either '0' or '1'.

----

### 1869. 哪种连续子字符串更长
给你一个二进制字符串 s 。如果字符串中由 1 组成的 最长 连续子字符串 严格长于 由 0 组成的 最长 连续子字符串，返回 true ；否则，返回 false 。

* 例如，s = "110100010" 中，由 1 组成的最长连续子字符串的长度是 2 ，由 0 组成的最长连续子字符串的长度是 3 。

注意，如果字符串中不存在 0 ，此时认为由 0 组成的最长连续子字符串的长度是 0 。字符串中不存在 1 的情况也适用此规则。



示例 1：

	输入：s = "1101"
	输出：true
	解释：
	由 1 组成的最长连续子字符串的长度是 2："1101"
	由 0 组成的最长连续子字符串的长度是 1："1101"
	由 1 组成的子字符串更长，故返回 true 。

示例 2：

	输入：s = "111000"
	输出：false
	解释：
	由 1 组成的最长连续子字符串的长度是 3："111000"
	由 0 组成的最长连续子字符串的长度是 3："111000"
	由 1 组成的子字符串不比由 0 组成的子字符串长，故返回 false 。

示例 3：

	输入：s = "110100010"
	输出：false
	解释：
	由 1 组成的最长连续子字符串的长度是 2："110100010"
	由 0 组成的最长连续子字符串的长度是 3："110100010"
	由 1 组成的子字符串不比由 0 组成的子字符串长，故返回 false 。



提示：

* 1 <= s.length <= 100
* s[i] 不是 '0' 就是 '1'

