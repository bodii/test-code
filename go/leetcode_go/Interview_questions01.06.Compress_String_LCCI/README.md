###	面试题 01.06. Compress String LCCI
Implement a method to perform basic string compression using the counts of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the "compressed" string would not become smaller than the original string, your method should return the original string. You can assume the string has only uppercase and lowercase letters (a - z).

Example 1:

	Input: "aabcccccaaa"
	Output: "a2b1c5a3"

Example 2:

	Input: "abbccd"
	Output: "abbccd"
	Explanation:
	The compressed string is "a1b2c2d1", which is longer than the original string.

Note:

* 0 <= S.length <= 50000

----

### 面试题 01.06. 字符串压缩
字符串压缩。利用字符重复出现的次数，编写一种方法，实现基本的字符串压缩功能。比如，字符串aabcccccaaa会变为a2b1c5a3。若“压缩”后的字符串没有变短，则返回原先的字符串。你可以假设字符串中只包含大小写英文字母（a至z）。

示例1:

	输入："aabcccccaaa"
	输出："a2b1c5a3"

示例2:

	输入："abbccd"
	输出："abbccd"
	解释："abbccd"压缩后为"a1b2c2d1"，比原字符串长度更长。

提示：

* 字符串长度在[0, 50000]范围内。

