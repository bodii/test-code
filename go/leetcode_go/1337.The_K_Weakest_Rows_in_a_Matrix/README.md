### 1337. The K Weakest Rows in a Matrix
You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians). The soldiers are positioned in front of the civilians. That is, all the 1's will appear to the left of all the 0's in each row.

A row i is weaker than a row j if one of the following is true:

* The number of soldiers in row i is less than the number of soldiers in row j.
* Both rows have the same number of soldiers and i < j.

Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.



Example 1:

Input: mat =
	[[1,1,0,0,0],
	 [1,1,1,1,0],
	 [1,0,0,0,0],
	 [1,1,0,0,0],
	 [1,1,1,1,1]],
	k = 3
	Output: [2,0,3]
	Explanation:
	The number of soldiers in each row is:
	- Row 0: 2
	- Row 1: 4
	- Row 2: 1
	- Row 3: 2
	- Row 4: 5
	The rows ordered from weakest to strongest are [2,0,3,1,4].

Example 2:

	Input: mat =
	[[1,0,0,0],
	 [1,1,1,1],
	 [1,0,0,0],
	 [1,0,0,0]],
	k = 2
	Output: [0,2]
	Explanation:
	The number of soldiers in each row is:
	- Row 0: 1
	- Row 1: 4
	- Row 2: 1
	- Row 3: 1
	The rows ordered from weakest to strongest are [0,2,3,1].



Constraints:

* m == mat.length
* n == mat[i].length
* 2 <= n, m <= 100
* 1 <= k <= m
* matrix[i][j] is either 0 or 1.

----

### 1337. 矩阵中战斗力最弱的 K 行
给你一个大小为 m * n 的矩阵 mat，矩阵由若干军人和平民组成，分别用 1 和 0 表示。

请你返回矩阵中战斗力最弱的 k 行的索引，按从最弱到最强排序。

如果第 i 行的军人数量少于第 j 行，或者两行军人数量相同但 i 小于 j，那么我们认为第 i 行的战斗力比第 j 行弱。

军人 总是 排在一行中的靠前位置，也就是说 1 总是出现在 0 之前。



示例 1：

	输入：mat =
	[[1,1,0,0,0],
	 [1,1,1,1,0],
	 [1,0,0,0,0],
	 [1,1,0,0,0],
	 [1,1,1,1,1]],
	k = 3
	输出：[2,0,3]
	解释：
	每行中的军人数目：
	行 0 -> 2
	行 1 -> 4
	行 2 -> 1
	行 3 -> 2
	行 4 -> 5
	从最弱到最强对这些行排序后得到 [2,0,3,1,4]

示例 2：

	输入：mat =
	[[1,0,0,0],
	 [1,1,1,1],
	 [1,0,0,0],
	 [1,0,0,0]],
	k = 2
	输出：[0,2]
	解释：
	每行中的军人数目：
	行 0 -> 1
	行 1 -> 4
	行 2 -> 1
	行 3 -> 1
	从最弱到最强对这些行排序后得到 [0,2,3,1]



提示：

* m == mat.length
* n == mat[i].length
* 2 <= n, m <= 100
* 1 <= k <= m
* matrix[i][j] 不是 0 就是 1

