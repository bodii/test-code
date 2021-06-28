### 893. Groups of Special-Equivalent Strings
You are given an array words of strings.

A move onto s consists of swapping any two even indexed characters of s, or any two odd indexed characters of s.

Two strings s and t are special-equivalent if after any number of moves onto s, s == t.

For example, s = "zzxy" and t = "xyzz" are special-equivalent because we may make the moves "zzxy" -> "xzzy" -> "xyzz" that swap s[0] and s[2], then s[1] and s[3].

Now, a group of special-equivalent strings from words is a non-empty subset of words such that:

* Every pair of strings in the group are special equivalent, and;
* The group is the largest size possible (ie., there isn't a string s not in the group such that s is special equivalent to every string in the group)

Return the number of groups of special-equivalent strings from words.


Example 1:

	Input: words = ["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
	Output: 3
	Explanation:
	One group is ["abcd", "cdab", "cbad"], since they are all pairwise special equivalent, and none of the other strings are all pairwise special equivalent to these.

	The other two groups are ["xyzz", "zzxy"] and ["zzyx"].  Note that in particular, "zzxy" is not special equivalent to "zzyx".

Example 2:

	Input: words = ["abc","acb","bac","bca","cab","cba"]
	Output: 3



Note:

* 1 <= words.length <= 1000
* 1 <= words[i].length <= 20
* All words[i] have the same length.
* All words[i] consist of only lowercase letters.

----

### 893. 特殊等价字符串组
你将得到一个字符串数组 A。

每次移动都可以交换 S 的任意两个偶数下标的字符或任意两个奇数下标的字符。

如果经过任意次数的移动，S == T，那么两个字符串 S 和 T 是 特殊等价 的。

例如，S = "zzxy" 和 T = "xyzz" 是一对特殊等价字符串，因为可以先交换 S[0] 和 S[2]，然后交换 S[1] 和 S[3]，使得 "zzxy" -> "xzzy" -> "xyzz" 。

现在规定，A 的 一组特殊等价字符串 就是 A 的一个同时满足下述条件的非空子集：

* 该组中的每一对字符串都是 特殊等价 的
* 该组字符串已经涵盖了该类别中的所有特殊等价字符串，容量达到理论上的最大值（也就是说，如果一个字符串不在该组中，那么这个字符串就 不会 与该组内任何字符串特殊等价）

返回 A 中特殊等价字符串组的数量。



示例 1：

	输入：["abcd","cdab","cbad","xyzz","zzxy","zzyx"]
	输出：3
	解释：
	其中一组为 ["abcd", "cdab", "cbad"]，因为它们是成对的特殊等价字符串，且没有其他字符串与这些字符串特殊等价。
	另外两组分别是 ["xyzz", "zzxy"] 和 ["zzyx"]。特别需要注意的是，"zzxy" 不与 "zzyx" 特殊等价。

示例 2：

	输入：["abc","acb","bac","bca","cab","cba"]
	输出：3
	解释：3 组 ["abc","cba"]，["acb","bca"]，["bac","cab"]



提示：

* 1 <= A.length <= 1000
* 1 <= A[i].length <= 20
* 所有 A[i] 都具有相同的长度。
* 所有 A[i] 都只由小写字母组成。

