### 2133. Check if Every Row and Column Contains All Numbers
An n x n matrix is valid if every row and every column contains all the integers from 1 to n (inclusive).

Given an n x n integer matrix matrix, return true if the matrix is valid. Otherwise, return false.



Example 1:

	Input: matrix = [[1,2,3],[3,1,2],[2,3,1]]
	Output: true
	Explanation: In this case, n = 3, and every row and column contains the numbers 1, 2, and 3.
	Hence, we return true.

Example 2:

	Input: matrix = [[1,1,1],[1,2,3],[1,2,3]]
	Output: false
	Explanation: In this case, n = 3, but the first row and the first column do not contain the numbers 2 or 3.
	Hence, we return false.



Constraints:

* n == matrix.length == matrix[i].length
* 1 <= n <= 100
* 1 <= matrix[i][j] <= n

----

### 2133. 检查是否每一行每一列都包含全部整数
对一个大小为 n x n 的矩阵而言，如果其每一行和每一列都包含从 1 到 n 的 全部 整数（含 1 和 n），则认为该矩阵是一个 有效 矩阵。

给你一个大小为 n x n 的整数矩阵 matrix ，请你判断矩阵是否为一个有效矩阵：如果是，返回 true ；否则，返回 false 。



示例 1：

	输入：matrix = [[1,2,3],[3,1,2],[2,3,1]]
	输出：true
	解释：在此例中，n = 3 ，每一行和每一列都包含数字 1、2、3 。
	因此，返回 true 。

示例 2：

	输入：matrix = [[1,1,1],[1,2,3],[1,2,3]]
	输出：false
	解释：在此例中，n = 3 ，但第一行和第一列不包含数字 2 和 3 。
	因此，返回 false 。



提示：

* n == matrix.length == matrix[i].length
* 1 <= n <= 100
* 1 <= matrix[i][j] <= n

