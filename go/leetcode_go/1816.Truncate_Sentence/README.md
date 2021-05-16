### 1816. Truncate Sentence
A sentence is a list of words that are separated by a single space with no leading or trailing spaces. Each of the words consists of only uppercase and lowercase English letters (no punctuation).

* For example, "Hello World", "HELLO", and "hello world hello world" are all sentences.

You are given a sentence s and an integer k. You want to truncate s such that it contains only the first k words. Return s after truncating it.



Example 1:

	Input: s = "Hello how are you Contestant", k = 4
	Output: "Hello how are you"
	Explanation:
	The words in s are ["Hello", "how" "are", "you", "Contestant"].
	The first 4 words are ["Hello", "how", "are", "you"].
	Hence, you should return "Hello how are you".

Example 2:

	Input: s = "What is the solution to this problem", k = 4
	Output: "What is the solution"
	Explanation:
	The words in s are ["What", "is" "the", "solution", "to", "this", "problem"].
	The first 4 words are ["What", "is", "the", "solution"].
	Hence, you should return "What is the solution".

Example 3:

	Input: s = "chopper is not a tanuki", k = 5
	Output: "chopper is not a tanuki"



Constraints:

* 1 <= s.length <= 500
* k is in the range [1, the number of words in s].
* s consist of only lowercase and uppercase English letters and spaces.
* The words in s are separated by a single space.
* There are no leading or trailing spaces.

----

### 1816. 截断句子
句子 是一个单词列表，列表中的单词之间用单个空格隔开，且不存在前导或尾随空格。每个单词仅由大小写英文字母组成（不含标点符号）。

    例如，"Hello World"、"HELLO" 和 "hello world hello world" 都是句子。

给你一个句子 s 和一个整数 k ，请你将 s 截断 ，使截断后的句子仅含 前 k 个单词。返回 截断 s 后得到的句子。



示例 1：

	输入：s = "Hello how are you Contestant", k = 4
	输出："Hello how are you"
	解释：
	s 中的单词为 ["Hello", "how" "are", "you", "Contestant"]
	前 4 个单词为 ["Hello", "how", "are", "you"]
	因此，应当返回 "Hello how are you"

示例 2：

	输入：s = "What is the solution to this problem", k = 4
	输出："What is the solution"
	解释：
	s 中的单词为 ["What", "is" "the", "solution", "to", "this", "problem"]
	前 4 个单词为 ["What", "is", "the", "solution"]
	因此，应当返回 "What is the solution"

示例 3：

	输入：s = "chopper is not a tanuki", k = 5
	输出："chopper is not a tanuki"



提示：

* 1 <= s.length <= 500
* k 的取值范围是 [1,  s 中单词的数目]
* s 仅由大小写英文字母和空格组成
* s 中的单词之间由单个空格隔开
* 不存在前导或尾随空格
