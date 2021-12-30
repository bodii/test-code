### 696. Count Binary Substrings
Give a binary string s, return the number of non-empty substrings that have the same number of 0's and 1's, and all the 0's and all the 1's in these substrings are grouped consecutively.

Substrings that occur multiple times are counted the number of times they occur.



Example 1:

	Input: s = "00110011"
	Output: 6
	Explanation: There are 6 substrings that have equal number of consecutive 1's and 0's: "0011", "01", "1100", "10", "0011", and "01".
	Notice that some of these substrings repeat and are counted the number of times they occur.
	Also, "00110011" is not a valid substring because all the 0's (and 1's) are not grouped together.

Example 2:

	Input: s = "10101"
	Output: 4
	Explanation: There are 4 substrings: "10", "01", "10", "01" that have equal number of consecutive 1's and 0's.



Constraints:

* 1 <= s.length <= 10^5
* s[i] is either '0' or '1'.

----

### 696. 计数二进制子串
给定一个字符串 s，统计并返回具有相同数量 0 和 1 的非空（连续）子字符串的数量，并且这些子字符串中的所有 0 和所有 1 都是成组连续的。

重复出现（不同位置）的子串也要统计它们出现的次数。


示例 1：

	输入：s = "00110011"
	输出：6
	解释：6 个子串满足具有相同数量的连续 1 和 0 ："0011"、"01"、"1100"、"10"、"0011" 和 "01" 。
	注意，一些重复出现的子串（不同位置）要统计它们出现的次数。
	另外，"00110011" 不是有效的子串，因为所有的 0（还有 1 ）没有组合在一起。

示例 2：

	输入：s = "10101"
	输出：4
	解释：有 4 个子串："10"、"01"、"10"、"01" ，具有相同数量的连续 1 和 0 。



提示：

* 1 <= s.length <= 10^5
* s[i] 为 '0' 或 '1'

