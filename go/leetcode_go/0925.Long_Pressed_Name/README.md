### 925. Long Pressed Name
Your friend is typing his name into a keyboard. Sometimes, when typing a character c, the key might get long pressed, and the character will be typed 1 or more times.

You examine the typed characters of the keyboard. Return True if it is possible that it was your friends name, with some characters (possibly none) being long pressed.


Example 1:

	Input: name = "alex", typed = "aaleex"
	Output: true
	Explanation: 'a' and 'e' in 'alex' were long pressed.

Example 2:

	Input: name = "saeed", typed = "ssaaedd"
	Output: false
	Explanation: 'e' must have been pressed twice, but it wasn't in the typed output.

Example 3:

	Input: name = "leelee", typed = "lleeelee"
	Output: true

Example 4:

	Input: name = "laiden", typed = "laiden"
	Output: true
	Explanation: It's not necessary to long press any character.



Constraints:

* 1 <= name.length <= 1000
* 1 <= typed.length <= 1000
* name and typed contain only lowercase English letters.

----

### 925. 长按键入
你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。

你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。



示例 1：

	输入：name = "alex", typed = "aaleex"
	输出：true
	解释：'alex' 中的 'a' 和 'e' 被长按。

示例 2：

	输入：name = "saeed", typed = "ssaaedd"
	输出：false
	解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。

示例 3：

	输入：name = "leelee", typed = "lleeelee"
	输出：true

示例 4：

	输入：name = "laiden", typed = "laiden"
	输出：true
	解释：长按名字中的字符并不是必要的。



提示：

* name.length <= 1000
* typed.length <= 1000
* name 和 typed 的字符都是小写字母。

