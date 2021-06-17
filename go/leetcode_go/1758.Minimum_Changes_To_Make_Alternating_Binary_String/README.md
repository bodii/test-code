### 1758. Minimum Changes To Make Alternating Binary String
You are given a string s consisting only of the characters '0' and '1'. In one operation, you can change any '0' to '1' or vice versa.

The string is called alternating if no two adjacent characters are equal. For example, the string "010" is alternating, while the string "0100" is not.

Return the minimum number of operations needed to make s alternating.



Example 1:

	Input: s = "0100"
	Output: 1
	Explanation: If you change the last character to '1', s will be "0101", which is alternating.

Example 2:

	Input: s = "10"
	Output: 0
	Explanation: s is already alternating.

Example 3:

	Input: s = "1111"
	Output: 2
	Explanation: You need two operations to reach "0101" or "1010".



Constraints:

* 1 <= s.length <= 10<sup>4</sup>
* s[i] is either '0' or '1'.

----

### 1758. 生成交替二进制字符串的最少操作数
给你一个仅由字符 '0' 和 '1' 组成的字符串 s 。一步操作中，你可以将任一 '0' 变成 '1' ，或者将 '1' 变成 '0' 。

交替字符串 定义为：如果字符串中不存在相邻两个字符相等的情况，那么该字符串就是交替字符串。例如，字符串 "010" 是交替字符串，而字符串 "0100" 不是。

返回使 s 变成 交替字符串 所需的 最少 操作数。



示例 1：

	输入：s = "0100"
	输出：1
	解释：如果将最后一个字符变为 '1' ，s 就变成 "0101" ，即符合交替字符串定义。

示例 2：

	输入：s = "10"
	输出：0
	解释：s 已经是交替字符串。

示例 3：

	输入：s = "1111"
	输出：2
	解释：需要 2 步操作得到 "0101" 或 "1010" 。



提示：

* 1 <= s.length <= 10<sup>4</sup>
* s[i] 是 '0' 或 '1'

