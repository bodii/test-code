### 1078. Occurrences After Bigram
Given words first and second, consider occurrences in some text of the form "first second third", where second comes immediately after first, and third comes immediately after second.

For each such occurrence, add "third" to the answer, and return the answer.



 Example 1:

	 Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
	 Output: ["girl","student"]

 Example 2:

	 Input: text = "we will we will rock you", first = "we", second = "will"
	 Output: ["we","rock"]



  Note:
	* 1 <= text.length <= 1000
	* text consists of space separated words, where each word consists of lowercase English letters.
	* 1 <= first.length, second.length <= 10
	* first and second consist of lowercase English letters.

--------------------

### 1078. Bigram 分词
给出第一个词 first 和第二个词 second，考虑在某些文本 text 中可能以 "first second third" 形式出现的情况，其中 second 紧随 first 出现，third 紧随 second 出现。

对于每种这样的情况，将第三个词 "third" 添加到答案中，并返回答案。



 示例 1：

	 输入：text = "alice is a good girl she is a good student", first = "a", second = "good"
	 输出：["girl","student"]

 示例 2：

	 输入：text = "we will we will rock you", first = "we", second = "will"
	 输出：["we","rock"]



提示：
* 1 <= text.length <= 1000
* text 由一些用空格分隔的单词组成，每个单词都由小写英文字母组成
* 1 <= first.length, second.length <= 10
* first 和 second 由小写英文字母组成

