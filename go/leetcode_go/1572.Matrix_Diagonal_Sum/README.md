### 1572. Matrix Diagonal Sum
Given a square matrix mat, return the sum of the matrix diagonals.

Only include the sum of all the elements on the primary diagonal and all the elements on the secondary diagonal that are not part of the primary diagonal.



Example 1:
	Input: mat = [[1,2,3],
				  [4,5,6],
				  [7,8,9]]
	Output: 25
	Explanation: Diagonals sum: 1 + 5 + 9 + 3 + 7 = 25
	Notice that element mat[1][1] = 5 is counted only once.

Example 2:

	Input: mat = [[1,1,1,1],
				  [1,1,1,1],
				  [1,1,1,1],
				  [1,1,1,1]]
	Output: 8

Example 3:

	Input: mat = [[5]]
	Output: 5

 

Constraints:

* n == mat.length == mat[i].length
* 1 <= n <= 100
* 1 <= mat[i][j] <= 100

----
### 1572. 矩阵对角线元素的和
给你一个正方形矩阵 mat，请你返回矩阵对角线元素的和。

请你返回在矩阵主对角线上的元素和副对角线上且不在主对角线上元素的和。



示例  1：
	输入：mat = [[1,2,3],
				[4,5,6],
				[7,8,9]]
	输出：25
	解释：对角线的和为：1 + 5 + 9 + 3 + 7 = 25
	请注意，元素 mat[1][1] = 5 只会被计算一次。

示例  2：

	输入：mat = [[1,1,1,1],
				[1,1,1,1],
				[1,1,1,1],
				[1,1,1,1]]
	输出：8

示例 3：

	输入：mat = [[5]]
	输出：5



提示：

* n == mat.length == mat[i].length
* 1 <= n <= 100
* 1 <= mat[i][j] <= 100

