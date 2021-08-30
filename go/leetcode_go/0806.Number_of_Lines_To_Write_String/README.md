### 806. Number of Lines To Write String
You are given a string s of lowercase English letters and an array widths denoting how many pixels wide each lowercase English letter is. Specifically, widths[0] is the width of 'a', widths[1] is the width of 'b', and so on.

You are trying to write s across several lines, where each line is no longer than 100 pixels. Starting at the beginning of s, write as many letters on the first line such that the total width does not exceed 100 pixels. Then, from where you stopped in s, continue writing as many letters as you can on the second line. Continue this process until you have written all of s.

Return an array result of length 2 where:

* result[0] is the total number of lines.
* result[1] is the width of the last line in pixels.



Example 1:

	Input: widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10], s = "abcdefghijklmnopqrstuvwxyz"
	Output: [3,60]
	Explanation: You can write s as follows:
	abcdefghij  // 100 pixels wide
	klmnopqrst  // 100 pixels wide
	uvwxyz      // 60 pixels wide
	There are a total of 3 lines, and the last line is 60 pixels wide.

Example 2:

	Input: widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10], s = "bbbcccdddaaa"
	Output: [2,4]
	Explanation: You can write s as follows:
	bbbcccdddaa  // 98 pixels wide
	a            // 4 pixels wide
	There are a total of 2 lines, and the last line is 4 pixels wide.



Constraints:

* widths.length == 26
* 2 <= widths[i] <= 10
* 1 <= s.length <= 1000
* s contains only lowercase English letters.

----

### 806. 写字符串需要的行数
我们要把给定的字符串 S 从左到右写到每一行上，每一行的最大宽度为100个单位，如果我们在写某个字母的时候会使这行超过了100 个单位，那么我们应该把这个字母写到下一行。我们给定了一个数组 widths ，这个数组 widths[0] 代表 'a' 需要的单位， widths[1] 代表 'b' 需要的单位，...， widths[25] 代表 'z' 需要的单位。

现在回答两个问题：至少多少行能放下S，以及最后一行使用的宽度是多少个单位？将你的答案作为长度为2的整数列表返回。

示例 1:
输入:
	widths = [10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
	S = "abcdefghijklmnopqrstuvwxyz"
	输出: [3, 60]
	解释:
	所有的字符拥有相同的占用单位10。所以书写所有的26个字母，
	我们需要2个整行和占用60个单位的一行。

示例 2:
输入:
	widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10]
	S = "bbbcccdddaaa"
	输出: [2, 4]
	解释:
	除去字母'a'所有的字符都是相同的单位10，并且字符串 "bbbcccdddaa" 将会覆盖 9 * 10 + 2 * 4 = 98 个单位.
	最后一个字母 'a' 将会被写到第二行，因为第一行只剩下2个单位了。
	所以，这个答案是2行，第二行有4个单位宽度。



注:

* 字符串 S 的长度在 [1, 1000] 的范围。
* S 只包含小写字母。
* widths 是长度为 26的数组。
* widths[i] 值的范围在 [2, 10]。

