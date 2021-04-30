### 944. Delete Columns to Make Sorted

You are given an array of n strings strs, all of the same length.

The strings can be arranged such that there is one on each line, making a grid. For example, strs = ["abc", "bce", "cae"] can be arranged as:

	abc
	bce
	cae

You want to delete the columns that are not sorted lexicographically. In the above example (0-indexed), columns 0 ('a', 'b', 'c') and 2 ('c', 'e', 'e') are sorted while column 1 ('b', 'c', 'a') is not, so you would delete column 1.

Return the number of columns that you will delete.



Example 1:

Input: strs = ["cba","daf","ghi"]
	Output: 1
	Explanation: The grid looks as follows:
	  cba
	  daf
	  ghi
	Columns 0 and 2 are sorted, but column 1 is not, so you only need to delete 1 column.

Example 2:

	Input: strs = ["a","b"]
	Output: 0
	Explanation: The grid looks as follows:
	  a
	  b
	Column 0 is the only column and is sorted, so you will not delete any columns.

Example 3:

	Input: strs = ["zyx","wvu","tsr"]
	Output: 3
	Explanation: The grid looks as follows:
	  zyx
	  wvu
	  tsr
	All 3 columns are not sorted, so you will delete all 3.



Constraints:

* n == strs.length
* 1 <= n <= 100
* 1 <= strs[i].length <= 1000
* strs[i] consists of lowercase English letters.

--------------

### 944. 删列造序
给你由 n 个小写字母字符串组成的数组 strs，其中每个字符串长度相等。

这些字符串可以每个一行，排成一个网格。例如，strs = ["abc", "bce", "cae"] 可以排列为：

	abc
	bce
	cae

你需要找出并删除 不是按字典序升序排列的 列。在上面的例子（下标从 0 开始）中，列 0（'a', 'b', 'c'）和列 2（'c', 'e', 'e'）都是按升序排列的，而列 1（'b', 'c', 'a'）不是，所以要删除列 1 。

返回你需要删除的列数。



示例 1：

	输入：strs = ["cba","daf","ghi"]
	输出：1
	解释：网格示意如下：
	  cba
	  daf
	  ghi
	列 0 和列 2 按升序排列，但列 1 不是，所以只需要删除列 1 。

示例 2：

	输入：strs = ["a","b"]
	输出：0
	解释：网格示意如下：
	  a
	  b
	只有列 0 这一列，且已经按升序排列，所以不用删除任何列。

示例 3：

	输入：strs = ["zyx","wvu","tsr"]
	输出：3
	解释：网格示意如下：
	  zyx
	  wvu
	  tsr
	所有 3 列都是非升序排列的，所以都要删除。



提示：

* n == strs.length
* 1 <= n <= 100
* 1 <= strs[i].length <= 1000
* strs[i] 由小写英文字母组成

