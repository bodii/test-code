### 1071. Greatest Common Divisor of Strings
or two strings s and t, we say "t divides s" if and only if s = t + ... + t  (t concatenated with itself 1 or more times)

Given two strings str1 and str2, return the largest string x such that x divides both str1 and str2.



Example 1:

	Input: str1 = "ABCABC", str2 = "ABC"
	Output: "ABC"

Example 2:

	Input: str1 = "ABABAB", str2 = "ABAB"
	Output: "AB"

Example 3:

	Input: str1 = "LEET", str2 = "CODE"
	Output: ""

Example 4:

	Input: str1 = "ABCDEF", str2 = "ABC"
	Output: ""



Constraints:

* 1 <= str1.length <= 1000
* 1 <= str2.length <= 1000
* str1 and str2 consist of English uppercase letters.


----

### 1071. 字符串的最大公因子
对于字符串 S 和 T，只有在 S = T + ... + T（T 自身连接 1 次或多次）时，我们才认定 “T 能除尽 S”。

返回最长字符串 X，要求满足 X 能除尽 str1 且 X 能除尽 str2。



示例 1：

	输入：str1 = "ABCABC", str2 = "ABC"
	输出："ABC"

示例 2：

	输入：str1 = "ABABAB", str2 = "ABAB"
	输出："AB"

示例 3：

	输入：str1 = "LEET", str2 = "CODE"
	输出：""



提示：

* 1 <= str1.length <= 1000
* 1 <= str2.length <= 1000
* str1[i] 和 str2[i] 为大写英文字母


