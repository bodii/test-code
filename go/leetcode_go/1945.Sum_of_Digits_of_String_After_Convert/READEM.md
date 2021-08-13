### 1945. Sum of Digits of String After Convert
You are given a string s consisting of lowercase English letters, and an integer k.

First, convert s into an integer by replacing each letter with its position in the alphabet (i.e., replace 'a' with 1, 'b' with 2, ..., 'z' with 26). Then, transform the integer by replacing it with the sum of its digits. Repeat the transform operation k times in total.

For example, if s = "zbax" and k = 2, then the resulting integer would be 8 by the following operations:

* Convert: "zbax" ➝ "(26)(2)(1)(24)" ➝ "262124" ➝ 262124
* Transform #1: 262124 ➝ 2 + 6 + 2 + 1 + 2 + 4 ➝ 17
* Transform #2: 17 ➝ 1 + 7 ➝ 8

Return the resulting integer after performing the operations described above.

 

Example 1:

	Input: s = "iiii", k = 1
	Output: 36
	Explanation: The operations are as follows:
	- Convert: "iiii" ➝ "(9)(9)(9)(9)" ➝ "9999" ➝ 9999
	- Transform #1: 9999 ➝ 9 + 9 + 9 + 9 ➝ 36
	Thus the resulting integer is 36.

Example 2:

	Input: s = "leetcode", k = 2
	Output: 6
	Explanation: The operations are as follows:
	- Convert: "leetcode" ➝ "(12)(5)(5)(20)(3)(15)(4)(5)" ➝ "12552031545" ➝ 12552031545
	- Transform #1: 12552031545 ➝ 1 + 2 + 5 + 5 + 2 + 0 + 3 + 1 + 5 + 4 + 5 ➝ 33
	- Transform #2: 33 ➝ 3 + 3 ➝ 6
	Thus the resulting integer is 6.

Example 3:

	Input: s = "zbax", k = 2
	Output: 8

 

Constraints:

* 1 <= s.length <= 100
* 1 <= k <= 10
* s consists of lowercase English letters.

----

### 1945. 字符串转化后的各位数字之和
给你一个由小写字母组成的字符串 s ，以及一个整数 k 。

首先，用字母在字母表中的位置替换该字母，将 s 转化 为一个整数（也就是，'a' 用 1 替换，'b' 用 2 替换，... 'z' 用 26 替换）。接着，将整数 转换 为其 各位数字之和 。共重复 转换 操作 k 次 。

例如，如果 s = "zbax" 且 k = 2 ，那么执行下述步骤后得到的结果是整数 8 ：

* 转化："zbax" ➝ "(26)(2)(1)(24)" ➝ "262124" ➝ 262124
* 转换 #1：262124 ➝ 2 + 6 + 2 + 1 + 2 + 4 ➝ 17
* 转换 #2：17 ➝ 1 + 7 ➝ 8

返回执行上述操作后得到的结果整数。

 

示例 1：

	输入：s = "iiii", k = 1
	输出：36
	解释：操作如下：
	- 转化："iiii" ➝ "(9)(9)(9)(9)" ➝ "9999" ➝ 9999
	- 转换 #1：9999 ➝ 9 + 9 + 9 + 9 ➝ 36
	因此，结果整数为 36 。

示例 2：

	输入：s = "leetcode", k = 2
	输出：6
	解释：操作如下：
	- 转化："leetcode" ➝ "(12)(5)(5)(20)(3)(15)(4)(5)" ➝ "12552031545" ➝ 12552031545
	- 转换 #1：12552031545 ➝ 1 + 2 + 5 + 5 + 2 + 0 + 3 + 1 + 5 + 4 + 5 ➝ 33
	- 转换 #2：33 ➝ 3 + 3 ➝ 6
	因此，结果整数为 6 。

 

提示：

* 1 <= s.length <= 100
* 1 <= k <= 10
* s 由小写英文字母组成

