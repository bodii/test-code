### 1309. Decrypt String from Alphabet to Integer Mapping
Given a string s formed by digits ('0' - '9') and '#' . We want to map s to English lowercase characters as follows:

* Characters ('a' to 'i') are represented by ('1' to '9') respectively.
* Characters ('j' to 'z') are represented by ('10#' to '26#') respectively. 

Return the string formed after mapping.

It's guaranteed that a unique mapping will always exist.

 

Example 1:

	Input: s = "10#11#12"
	Output: "jkab"
	Explanation: "j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".

Example 2:

	Input: s = "1326#"
	Output: "acz"

Example 3:

	Input: s = "25#"
	Output: "y"

Example 4:

	Input: s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
	Output: "abcdefghijklmnopqrstuvwxyz"

 

Constraints:

* 1 <= s.length <= 1000
* s[i] only contains digits letters ('0'-'9') and '#' letter.
* s will be valid string such that mapping is always possible.

----
### 1309. 解码字母到整数映射
给你一个字符串 s，它由数字（'0' - '9'）和 '#' 组成。我们希望按下述规则将 s 映射为一些小写英文字符：

* 字符（'a' - 'i'）分别用（'1' - '9'）表示。
* 字符（'j' - 'z'）分别用（'10#' - '26#'）表示。 

返回映射之后形成的新字符串。

题目数据保证映射始终唯一。

 

示例 1：

	输入：s = "10#11#12"
	输出："jkab"
	解释："j" -> "10#" , "k" -> "11#" , "a" -> "1" , "b" -> "2".

示例 2：

	输入：s = "1326#"
	输出："acz"

示例 3：

	输入：s = "25#"
	输出："y"

示例 4：

	输入：s = "12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"
	输出："abcdefghijklmnopqrstuvwxyz"

 

提示：

* 1 <= s.length <= 1000
* s[i] 只包含数字（'0'-'9'）和 '#' 字符。
* s 是映射始终存在的有效字符串。

