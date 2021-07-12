### 1221. Split a String in Balanced Strings
Balanced strings are those that have an equal quantity of 'L' and 'R' characters.

Given a balanced string s, split it in the maximum amount of balanced strings.

Return the maximum amount of split balanced strings.



Example 1:

	Input: s = "RLRRLLRLRL"
	Output: 4
	Explanation: s can be split into "RL", "RRLL", "RL", "RL", each substring contains same number of 'L' and 'R'.

Example 2:

	Input: s = "RLLLLRRRLR"
	Output: 3
	Explanation: s can be split into "RL", "LLLRRR", "LR", each substring contains same number of 'L' and 'R'.

Example 3:

	Input: s = "LLLLRRRR"
	Output: 1
	Explanation: s can be split into "LLLLRRRR".

Example 4:

	Input: s = "RLRRRLLRLL"
	Output: 2
	Explanation: s can be split into "RL", "RRRLLRLL", since each substring contains an equal number of 'L' and 'R'



Constraints:

* 1 <= s.length <= 1000
* s[i] is either 'L' or 'R'.
* s is a balanced string.

----

### 1221. 分割平衡字符串
在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。

给你一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。

注意：分割得到的每个字符串都必须是平衡字符串。

返回可以通过分割得到的平衡字符串的 最大数量 。



示例 1：

	输入：s = "RLRRLLRLRL"
	输出：4
	解释：s 可以分割为 "RL"、"RRLL"、"RL"、"RL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。

示例 2：

	输入：s = "RLLLLRRRLR"
	输出：3
	解释：s 可以分割为 "RL"、"LLLRRR"、"LR" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。

示例 3：

	输入：s = "LLLLRRRR"
	输出：1
	解释：s 只能保持原样 "LLLLRRRR".

示例 4：

	输入：s = "RLRRRLLRLL"
	输出：2
	解释：s 可以分割为 "RL"、"RRRLLRLL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。



提示：

* 1 <= s.length <= 1000
* s[i] = 'L' 或 'R'
* s 是一个 平衡 字符串

