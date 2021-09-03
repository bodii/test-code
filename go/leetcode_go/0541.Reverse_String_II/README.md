### 541. Reverse String II
Given a string s and an integer k, reverse the first k characters for every 2k characters counting from the start of the string.

If there are fewer than k characters left, reverse all of them. If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.



Example 1:

	Input: s = "abcdefg", k = 2
	Output: "bacdfeg"

Example 2:

	Input: s = "abcd", k = 2
	Output: "bacd"



Constraints:

* 1 <= s.length <= 104
* s consists of only lowercase English letters.
* 1 <= k <= 104


----

### 541. 反转字符串 II
给定一个字符串 s 和一个整数 k，从字符串开头算起，每 2k 个字符反转前 k 个字符。

    如果剩余字符少于 k 个，则将剩余字符全部反转。
    如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。



示例 1：

	输入：s = "abcdefg", k = 2
	输出："bacdfeg"

示例 2：

	输入：s = "abcd", k = 2
	输出："bacd"



提示：

* 1 <= s.length <= 104
* s 仅由小写英文组成
* 1 <= k <= 104

