### 1002. Find Common Characters
Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates). You may return the answer in any order.

 

Example 1:

	Input: words = ["bella","label","roller"]
	Output: ["e","l","l"]

Example 2:

	Input: words = ["cool","lock","cook"]
	Output: ["c","o"]

 

Constraints:

* 1 <= words.length <= 100
* 1 <= words[i].length <= 100
* words[i] consists of lowercase English letters.

----

### 1002. 查找常用字符
给定仅有小写字母组成的字符串数组 A，返回列表中的每个字符串中都显示的全部字符（包括重复字符）组成的列表。例如，如果一个字符在每个字符串中出现 3 次，但不是 4 次，则需要在最终答案中包含该字符 3 次。

你可以按任意顺序返回答案。

 

示例 1：

	输入：["bella","label","roller"]
	输出：["e","l","l"]

示例 2：

	输入：["cool","lock","cook"]
	输出：["c","o"]

 

提示：

* 1 <= A.length <= 100
* 1 <= A[i].length <= 100
* A[i][j] 是小写字母
