### 1974. Minimum Time to Type Word Using Special Typewriter
There is a special typewriter with lowercase English letters 'a' to 'z' arranged in a circle with a pointer. A character can only be typed if the pointer is pointing to that character. The pointer is initially pointing to the character 'a'.

Each second, you may perform one of the following operations:

    Move the pointer one character counterclockwise or clockwise.
    Type the character the pointer is currently on.

Given a string word, return the minimum number of seconds to type out the characters in word.



Example 1:

	Input: word = "abc"
	Output: 5
	Explanation:
	The characters are printed as follows:
	- Type the character 'a' in 1 second since the pointer is initially on 'a'.
	- Move the pointer clockwise to 'b' in 1 second.
	- Type the character 'b' in 1 second.
	- Move the pointer clockwise to 'c' in 1 second.
	- Type the character 'c' in 1 second.

Example 2:

	Input: word = "bza"
	Output: 7
	Explanation:
	The characters are printed as follows:
	- Move the pointer clockwise to 'b' in 1 second.
	- Type the character 'b' in 1 second.
	- Move the pointer counterclockwise to 'z' in 2 seconds.
	- Type the character 'z' in 1 second.
	- Move the pointer clockwise to 'a' in 1 second.
	- Type the character 'a' in 1 second.

Example 3:

	Input: word = "zjpc"
	Output: 34
	Explanation:
	The characters are printed as follows:
	- Move the pointer counterclockwise to 'z' in 1 second.
	- Type the character 'z' in 1 second.
	- Move the pointer clockwise to 'j' in 10 seconds.
	- Type the character 'j' in 1 second.
	- Move the pointer clockwise to 'p' in 6 seconds.
	- Type the character 'p' in 1 second.
	- Move the pointer counterclockwise to 'c' in 13 seconds.
	- Type the character 'c' in 1 second.



Constraints:

* 1 <= word.length <= 100
* word consists of lowercase English letters.

----
### 1974. 使用特殊打字机键入单词的最少时间
有一个特殊打字机，它由一个 圆盘 和一个 指针 组成， 圆盘上标有小写英文字母 'a' 到 'z'。只有 当指针指向某个字母时，它才能被键入。指针 初始时 指向字符 'a' 。

每一秒钟，你可以执行以下操作之一：

    将指针 顺时针 或者 逆时针 移动一个字符。
    键入指针 当前 指向的字符。

给你一个字符串 word ，请你返回键入 word 所表示单词的 最少 秒数 。



示例 1：

	输入：word = "abc"
	输出：5
	解释：
	单词按如下操作键入：
	- 花 1 秒键入字符 'a' in 1 ，因为指针初始指向 'a' ，故不需移动指针。
	- 花 1 秒将指针顺时针移到 'b' 。
	- 花 1 秒键入字符 'b' 。
	- 花 1 秒将指针顺时针移到 'c' 。
	- 花 1 秒键入字符 'c' 。

示例 2：

	输入：word = "bza"
	输出：7
	解释：
	单词按如下操作键入：
	- 花 1 秒将指针顺时针移到 'b' 。
	- 花 1 秒键入字符 'b' 。
	- 花 2 秒将指针逆时针移到 'z' 。
	- 花 1 秒键入字符 'z' 。
	- 花 1 秒将指针顺时针移到 'a' 。
	- 花 1 秒键入字符 'a' 。

示例 3：

	输入：word = "zjpc"
	输出：34
	解释：
	单词按如下操作键入：
	- 花 1 秒将指针逆时针移到 'z' 。
	- 花 1 秒键入字符 'z' 。
	- 花 10 秒将指针顺时针移到 'j' 。
	- 花 1 秒键入字符 'j' 。
	- 花 6 秒将指针顺时针移到 'p' 。
	- 花 1 秒键入字符 'p' 。
	- 花 13 秒将指针逆时针移到 'c' 。
	- 花 1 秒键入字符 'c' 。



提示：

* 1 <= word.length <= 100
* word 只包含小写英文字母。

