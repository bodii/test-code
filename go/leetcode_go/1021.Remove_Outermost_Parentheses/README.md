### 1021. Remove Outermost Parentheses
A valid parentheses string is either empty "", "(" + A + ")", or A + B, where A and B are valid parentheses strings, and + represents string concatenation.

* For example, "", "()", "(())()", and "(()(()))" are all valid parentheses strings.

A valid parentheses string s is primitive if it is nonempty, and there does not exist a way to split it into s = A + B, with A and B nonempty valid parentheses strings.

Given a valid parentheses string s, consider its primitive decomposition: s = P1 + P2 + ... + Pk, where Pi are primitive valid parentheses strings.

Return s after removing the outermost parentheses of every primitive string in the primitive decomposition of s.



Example 1:

	Input: s = "(()())(())"
	Output: "()()()"
	Explanation:
	The input string is "(()())(())", with primitive decomposition "(()())" + "(())".
	After removing outer parentheses of each part, this is "()()" + "()" = "()()()".

Example 2:

	Input: s = "(()())(())(()(()))"
	Output: "()()()()(())"
	Explanation:
	The input string is "(()())(())(()(()))", with primitive decomposition "(()())" + "(())" + "(()(()))".
	After removing outer parentheses of each part, this is "()()" + "()" + "()(())" = "()()()()(())".

Example 3:

	Input: s = "()()"
	Output: ""
	Explanation:
	The input string is "()()", with primitive decomposition "()" + "()".
	After removing outer parentheses of each part, this is "" + "" = "".



Constraints:

* 1 <= s.length <= 10^5
* s[i] is either '(' or ')'.
* s is a valid parentheses string.

---- 

### 1021. 删除最外层的括号
有效括号字符串为空 ("")、"(" + A + ")" 或 A + B，其中 A 和 B 都是有效的括号字符串，+ 代表字符串的连接。例如，""，"()"，"(())()" 和 "(()(()))" 都是有效的括号字符串。

如果有效字符串 S 非空，且不存在将其拆分为 S = A+B 的方法，我们称其为原语（primitive），其中 A 和 B 都是非空有效括号字符串。

给出一个非空有效字符串 S，考虑将其进行原语化分解，使得：S = P_1 + P_2 + ... + P_k，其中 P_i 是有效括号字符串原语。

对 S 进行原语化分解，删除分解中每个原语字符串的最外层括号，返回 S 。



示例 1：

	输入："(()())(())"
	输出："()()()"
	解释：
	输入字符串为 "(()())(())"，原语化分解得到 "(()())" + "(())"，
	删除每个部分中的最外层括号后得到 "()()" + "()" = "()()()"。

示例 2：

	输入："(()())(())(()(()))"
	输出："()()()()(())"
	解释：
	输入字符串为 "(()())(())(()(()))"，原语化分解得到 "(()())" + "(())" + "(()(()))"，
	删除每个部分中的最外层括号后得到 "()()" + "()" + "()(())" = "()()()()(())"。

示例 3：

	输入："()()"
	输出：""
	解释：
	输入字符串为 "()()"，原语化分解得到 "()" + "()"，
	删除每个部分中的最外层括号后得到 "" + "" = ""。



提示：

* S.length <= 10000
* S[i] 为 "(" 或 ")"
* S 是一个有效括号字符串
