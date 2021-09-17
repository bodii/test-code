### 844. Backspace String Compare
Given two strings s and t, return true if they are equal when both are typed into empty text editors. '#' means a backspace character.

Note that after backspacing an empty text, the text will continue empty.



Example 1:

	Input: s = "ab#c", t = "ad#c"
	Output: true
	Explanation: Both s and t become "ac".

Example 2:

	Input: s = "ab##", t = "c#d#"
	Output: true
	Explanation: Both s and t become "".

Example 3:

	Input: s = "a##c", t = "#a#c"
	Output: true
	Explanation: Both s and t become "c".

Example 4:

	Input: s = "a#c", t = "b"
	Output: false
	Explanation: s becomes "c" while t becomes "b".



Constraints:

* 1 <= s.length, t.length <= 200
* s and t only contain lowercase letters and '#' characters.



Follow up: Can you solve it in O(n) time and O(1) space?

----

### 844. 比较含退格的字符串
给定 S 和 T 两个字符串，当它们分别被输入到空白的文本编辑器后，判断二者是否相等，并返回结果。 # 代表退格字符。

注意：如果对空文本输入退格字符，文本继续为空。



示例 1：

	输入：S = "ab#c", T = "ad#c"
	输出：true
	解释：S 和 T 都会变成 “ac”。

示例 2：

	输入：S = "ab##", T = "c#d#"
	输出：true
	解释：S 和 T 都会变成 “”。

示例 3：

	输入：S = "a##c", T = "#a#c"
	输出：true
	解释：S 和 T 都会变成 “c”。

示例 4：

	输入：S = "a#c", T = "b"
	输出：false
	解释：S 会变成 “c”，但 T 仍然是 “b”。



提示：

* 1 <= S.length <= 200
* 1 <= T.length <= 200
* S 和 T 只含有小写字母以及字符 '#'。



进阶：

你可以用 O(N) 的时间复杂度和 O(1) 的空间复杂度解决该问题吗？

