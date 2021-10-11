### 434. Number of Segments in a String
You are given a string s, return the number of segments in the string.

A segment is defined to be a contiguous sequence of non-space characters.



Example 1:

	Input: s = "Hello, my name is John"
	Output: 5
	Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]

Example 2:

	Input: s = "Hello"
	Output: 1

Example 3:

	Input: s = "love live! mu'sic forever"
	Output: 4

Example 4:

	Input: s = ""
	Output: 0



Constraints:

* 0 <= s.length <= 300
* s consists of lower-case and upper-case English letters, digits or one of the following characters "!@#$%^&*()_+-=',.:".
* The only space character in s is ' '.

----

### 434. 字符串中的单词数
统计字符串中的单词个数，这里的单词指的是连续的不是空格的字符。

请注意，你可以假定字符串里不包括任何不可打印的字符。

示例:

	输入: "Hello, my name is John"
	输出: 5
	解释: 这里的单词是指连续的不是空格的字符，所以 "Hello," 算作 1 个单词。

