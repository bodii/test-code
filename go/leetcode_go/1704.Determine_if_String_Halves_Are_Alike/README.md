### 1704. Determine if String Halves Are Alike
You are given a string s of even length. Split this string into two halves of equal lengths, and let a be the first half and b be the second half.

Two strings are alike if they have the same number of vowels ('a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'). Notice that s contains uppercase and lowercase letters.

Return true if a and b are alike. Otherwise, return false.



Example 1:

	Input: s = "book"
	Output: true
	Explanation: a = "bo" and b = "ok". a has 1 vowel and b has 1 vowel. Therefore, they are alike.

Example 2:

	Input: s = "textbook"
	Output: false
	Explanation: a = "text" and b = "book". a has 1 vowel whereas b has 2. Therefore, they are not alike.
	Notice that the vowel o is counted twice.

Example 3:

	Input: s = "MerryChristmas"
	Output: false

Example 4:

	Input: s = "AbCdEfGh"
	Output: true



Constraints:

* 2 <= s.length <= 1000
* s.length is even.
* s consists of uppercase and lowercase letters.

----

### 1704. 判断字符串的两半是否相似
给你一个偶数长度的字符串 s 。将其拆分成长度相同的两半，前一半为 a ，后一半为 b 。

两个字符串 相似 的前提是它们都含有相同数目的元音（'a'，'e'，'i'，'o'，'u'，'A'，'E'，'I'，'O'，'U'）。注意，s 可能同时含有大写和小写字母。

如果 a 和 b 相似，返回 true ；否则，返回 false 。



示例 1：

	输入：s = "book"
	输出：true
	解释：a = "bo" 且 b = "ok" 。a 中有 1 个元音，b 也有 1 个元音。所以，a 和 b 相似。

示例 2：

	输入：s = "textbook"
	输出：false
	解释：a = "text" 且 b = "book" 。a 中有 1 个元音，b 中有 2 个元音。因此，a 和 b 不相似。
	注意，元音 o 在 b 中出现两次，记为 2 个。

示例 3：

	输入：s = "MerryChristmas"
	输出：false

示例 4：

	输入：s = "AbCdEfGh"
	输出：true



提示：

* 2 <= s.length <= 1000
* s.length 是偶数
* s 由 大写和小写 字母组成
