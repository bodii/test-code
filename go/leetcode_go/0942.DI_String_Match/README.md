### 942. DI String Match
A permutation perm of n + 1 integers of all the integers in the range [0, n] can be represented as a string s of length n where:

* s[i] == 'I' if perm[i] < perm[i + 1], and
* s[i] == 'D' if perm[i] > perm[i + 1].

Given a string s, reconstruct the permutation perm and return it. If there are multiple valid permutations perm, return any of them.



Example 1:

	Input: s = "IDID"
	Output: [0,4,1,3,2]

Example 2:

	Input: s = "III"
	Output: [0,1,2,3]

Example 3:

	Input: s = "DDI"
	Output: [3,2,0,1]



Constraints:

* 1 <= s.length <= 105
* s[i] is either 'I' or 'D'.

----

### 942. 增减字符串匹配
给定只含 "I"（增大）或 "D"（减小）的字符串 S ，令 N = S.length。

返回 [0, 1, ..., N] 的任意排列 A 使得对于所有 i = 0, ..., N-1，都有：

* 如果 S[i] == "I"，那么 A[i] < A[i+1]
* 如果 S[i] == "D"，那么 A[i] > A[i+1]



示例 1：

	输入："IDID"
	输出：[0,4,1,3,2]

示例 2：

	输入："III"
	输出：[0,1,2,3]

示例 3：

	输入："DDI"
	输出：[3,2,0,1]



提示：

* 1 <= S.length <= 10000
* S 只包含字符 "I" 或 "D"。

