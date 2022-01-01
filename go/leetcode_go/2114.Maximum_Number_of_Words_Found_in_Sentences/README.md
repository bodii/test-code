### 2114. Maximum Number of Words Found in Sentences
A sentence is a list of words that are separated by a single space with no leading or trailing spaces.

You are given an array of strings sentences, where each sentences[i] represents a single sentence.

Return the maximum number of words that appear in a single sentence.

 

Example 1:

	Input: sentences = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"]
	Output: 6
	Explanation: 
	- The first sentence, "alice and bob love leetcode", has 5 words in total.
	- The second sentence, "i think so too", has 4 words in total.
	- The third sentence, "this is great thanks very much", has 6 words in total.
	Thus, the maximum number of words in a single sentence comes from the third sentence, which has 6 words.

Example 2:

	Input: sentences = ["please wait", "continue to fight", "continue to win"]
	Output: 3
	Explanation: It is possible that multiple sentences contain the same number of words. 
	In this example, the second and third sentences (underlined) have the same number of words.

 

Constraints:

* 1 <= sentences.length <= 100
* 1 <= sentences[i].length <= 100
* sentences[i] consists only of lowercase English letters and ' ' only.
* sentences[i] does not have leading or trailing spaces.
* All the words in sentences[i] are separated by a single space.

---- 
### 2114. 句子中的最多单词数
一个 句子 由一些 单词 以及它们之间的单个空格组成，句子的开头和结尾不会有多余空格。

给你一个字符串数组 sentences ，其中 sentences[i] 表示单个 句子 。

请你返回单个句子里 单词的最多数目 。

 

示例 1：

	输入：sentences = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"]
	输出：6
	解释：
	- 第一个句子 "alice and bob love leetcode" 总共有 5 个单词。
	- 第二个句子 "i think so too" 总共有 4 个单词。
	- 第三个句子 "this is great thanks very much" 总共有 6 个单词。
	所以，单个句子中有最多单词数的是第三个句子，总共有 6 个单词。

示例 2：

	输入：sentences = ["please wait", "continue to fight", "continue to win"]
	输出：3
	解释：可能有多个句子有相同单词数。
	这个例子中，第二个句子和第三个句子（加粗斜体）有相同数目的单词数。

 

提示：

* 1 <= sentences.length <= 100
* 1 <= sentences[i].length <= 100
* sentences[i] 只包含小写英文字母和 ' ' 。
* sentences[i] 的开头和结尾都没有空格。
* sentences[i] 中所有单词由单个空格隔开。


