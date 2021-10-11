### 859. Buddy Strings
Given two strings s and goal, return true if you can swap two letters in s so the result is equal to goal, otherwise, return false.

Swapping letters is defined as taking two indices i and j (0-indexed) such that i != j and swapping the characters at s[i] and s[j].

* For example, swapping at indices 0 and 2 in "abcd" results in "cbad".



Example 1:

	Input: s = "ab", goal = "ba"
	Output: true
	Explanation: You can swap s[0] = 'a' and s[1] = 'b' to get "ba", which is equal to goal.

Example 2:

	Input: s = "ab", goal = "ab"
	Output: false
	Explanation: The only letters you can swap are s[0] = 'a' and s[1] = 'b', which results in "ba" != goal.

Example 3:

	Input: s = "aa", goal = "aa"
	Output: true
	Explanation: You can swap s[0] = 'a' and s[1] = 'a' to get "aa", which is equal to goal.

Example 4:

	Input: s = "aaaaaaabc", goal = "aaaaaaacb"
	Output: true



Constraints:

    1 <= s.length, goal.length <= 2 * 10^4
    s and goal consist of lowercase letters.

----
### 859. 亲密字符串
给定两个由小写字母构成的字符串 A 和 B ，只要我们可以通过交换 A 中的两个字母得到与 B 相等的结果，就返回 true ；否则返回 false 。

交换字母的定义是取两个下标 i 和 j （下标从 0 开始），只要 i!=j 就交换 A[i] 和 A[j] 处的字符。例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。



示例 1：

	输入： A = "ab", B = "ba"
	输出： true
	解释： 你可以交换 A[0] = 'a' 和 A[1] = 'b' 生成 "ba"，此时 A 和 B 相等。

示例 2：

	输入： A = "ab", B = "ab"
	输出： false
	解释： 你只能交换 A[0] = 'a' 和 A[1] = 'b' 生成 "ba"，此时 A 和 B 不相等。

示例 3:

	输入： A = "aa", B = "aa"
	输出： true
	解释： 你可以交换 A[0] = 'a' 和 A[1] = 'a' 生成 "aa"，此时 A 和 B 相等。

示例 4：

	输入： A = "aaaaaaabc", B = "aaaaaaacb"
	输出： true

示例 5：

	输入： A = "", B = "aa"
	输出： false



提示：

* 0 <= A.length <= 20000
* 0 <= B.length <= 20000
* A 和 B 仅由小写字母构成。
