### 1805. Number of Different Integers in a String
You are given a string word that consists of digits and lowercase English letters.

You will replace every non-digit character with a space. For example, "a123bc34d8ef34" will become " 123  34 8  34". Notice that you are left with some integers that are separated by at least one space: "123", "34", "8", and "34".

Return the number of different integers after performing the replacement operations on word.

Two integers are considered different if their decimal representations without any leading zeros are different.



Example 1:

	Input: word = "a123bc34d8ef34"
	Output: 3
	Explanation: The three different integers are "123", "34", and "8". Notice that "34" is only counted once.

Example 2:

	Input: word = "leet1234code234"
	Output: 2

Example 3:

	Input: word = "a1b01c001"
	Output: 1
	Explanation: The three integers "1", "01", and "001" all represent the same integer because
	the leading zeros are ignored when comparing their decimal values.



Constraints:

* 1 <= word.length <= 1000
* word consists of digits and lowercase English letters.

----

### 1805. 字符串中不同整数的数目
给你一个字符串 word ，该字符串由数字和小写英文字母组成。

请你用空格替换每个不是数字的字符。例如，"a123bc34d8ef34" 将会变成 " 123  34 8  34" 。注意，剩下的这些整数为（相邻彼此至少有一个空格隔开）："123"、"34"、"8" 和 "34" 。

返回对 word 完成替换后形成的 不同 整数的数目。

只有当两个整数的 不含前导零 的十进制表示不同， 才认为这两个整数也不同。



示例 1：

	输入：word = "a123bc34d8ef34"
	输出：3
	解释：不同的整数有 "123"、"34" 和 "8" 。注意，"34" 只计数一次。

示例 2：

	输入：word = "leet1234code234"
	输出：2

示例 3：

	输入：word = "a1b01c001"
	输出：1
	解释："1"、"01" 和 "001" 视为同一个整数的十进制表示，因为在比较十进制值时会忽略前导零的存在。



提示：

* 1 <= word.length <= 1000
* word 由数字和小写英文字母组成

