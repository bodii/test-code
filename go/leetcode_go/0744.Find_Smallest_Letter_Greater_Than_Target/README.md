### 744. Find Smallest Letter Greater Than Target
Given a characters array letters that is sorted in non-decreasing order and a character target, return the smallest character in the array that is larger than target.

Note that the letters wrap around.

* For example, if target == 'z' and letters == ['a', 'b'], the answer is 'a'.



Example 1:

	Input: letters = ["c","f","j"], target = "a"
	Output: "c"

Example 2:

	Input: letters = ["c","f","j"], target = "c"
	Output: "f"

Example 3:

	Input: letters = ["c","f","j"], target = "d"
	Output: "f"



Constraints:

* 2 <= letters.length <= 10^4
* letters[i] is a lowercase English letter.
* letters is sorted in non-decreasing order.
* letters contains at least two different characters.
* target is a lowercase English letter.

----

### 744. 寻找比目标字母大的最小字母
给你一个排序后的字符列表 letters ，列表中只包含小写英文字母。另给出一个目标字母 target，请你寻找在这一有序列表里比目标字母大的最小字母。

在比较时，字母是依序循环出现的。举个例子：

* 如果目标字母 target = 'z' 并且字符列表为 letters = ['a', 'b']，则答案返回 'a'



示例：

	输入:
	letters = ["c", "f", "j"]
	target = "a"
	输出: "c"

输入:
	letters = ["c", "f", "j"]
	target = "c"
	输出: "f"

输入:
	letters = ["c", "f", "j"]
	target = "d"
	输出: "f"

输入:
	letters = ["c", "f", "j"]
	target = "g"
	输出: "j"

输入:
	letters = ["c", "f", "j"]
	target = "j"
	输出: "c"

输入:
	letters = ["c", "f", "j"]
	target = "k"
	输出: "c"



提示：

* letters长度范围在[2, 10000]区间内。
* letters 仅由小写字母组成，最少包含两个不同的字母。
* 目标字母target 是一个小写字母。


