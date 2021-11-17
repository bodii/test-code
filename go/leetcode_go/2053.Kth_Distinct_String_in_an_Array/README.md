### 2053. Kth Distinct String in an Array
A distinct string is a string that is present only once in an array.

Given an array of strings arr, and an integer k, return the kth distinct string present in arr. If there are fewer than k distinct strings, return an empty string "".

Note that the strings are considered in the order in which they appear in the array.



Example 1:

	Input: arr = ["d","b","c","b","c","a"], k = 2
	Output: "a"
	Explanation:
	The only distinct strings in arr are "d" and "a".
	"d" appears 1st, so it is the 1st distinct string.
	"a" appears 2nd, so it is the 2nd distinct string.
	Since k == 2, "a" is returned.

Example 2:

	Input: arr = ["aaa","aa","a"], k = 1
	Output: "aaa"
	Explanation:
	All strings in arr are distinct, so the 1st string "aaa" is returned.

Example 3:

	Input: arr = ["a","b","a"], k = 3
	Output: ""
	Explanation:
	The only distinct string is "b". Since there are fewer than 3 distinct strings, we return an empty string "".



Constraints:

* 1 <= k <= arr.length <= 1000
* 1 <= arr[i].length <= 5
* arr[i] consists of lowercase English letters.

----

### 2053. 数组中第 K 个独一无二的字符串
独一无二的字符串 指的是在一个数组中只出现过 一次 的字符串。

给你一个字符串数组 arr 和一个整数 k ，请你返回 arr 中第 k 个 独一无二的字符串 。如果 少于 k 个独一无二的字符串，那么返回 空字符串 "" 。

注意，按照字符串在原数组中的 顺序 找到第 k 个独一无二字符串。



示例 1:

	输入：arr = ["d","b","c","b","c","a"], k = 2
	输出："a"
	解释：
	arr 中独一无二字符串包括 "d" 和 "a" 。
	"d" 首先出现，所以它是第 1 个独一无二字符串。
	"a" 第二个出现，所以它是 2 个独一无二字符串。
	由于 k == 2 ，返回 "a" 。

示例 2:

	输入：arr = ["aaa","aa","a"], k = 1
	输出："aaa"
	解释：
	arr 中所有字符串都是独一无二的，所以返回第 1 个字符串 "aaa" 。

示例 3：

	输入：arr = ["a","b","a"], k = 3
	输出：""
	解释：
	唯一一个独一无二字符串是 "b" 。由于少于 3 个独一无二字符串，我们返回空字符串 "" 。



提示：

* 1 <= k <= arr.length <= 1000
* 1 <= arr[i].length <= 5
* arr[i] 只包含小写英文字母。


