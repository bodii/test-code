### 1380. Lucky Numbers in a Matrix
Given an m x n matrix of distinct numbers, return all lucky numbers in the matrix in any order.

A lucky number is an element of the matrix such that it is the minimum element in its row and maximum in its column.



Example 1:

	Input: matrix = [[3,7,8],[9,11,13],[15,16,17]]
	Output: [15]
	Explanation: 15 is the only lucky number since it is the minimum in its row and the maximum in its column

Example 2:

	Input: matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
	Output: [12]
	Explanation: 12 is the only lucky number since it is the minimum in its row and the maximum in its column.

Example 3:

	Input: matrix = [[7,8],[1,2]]
	Output: [7]
	Explanation: 7 is the only lucky number since it is the minimum in its row and the maximum in its column.

Example 4:

	Input: matrix = [[3,6],[7,1],[5,2],[4,8]]
	Output: []
	Explanation: There is no lucky number.



Constraints:

* m == mat.length
* n == mat[i].length
* 1 <= n, m <= 50
* 1 <= matrix[i][j] <= 105.
* All elements in the matrix are distinct.


----

### 1380. 矩阵中的幸运数
给你一个 m * n 的矩阵，矩阵中的数字 各不相同 。请你按 任意 顺序返回矩阵中的所有幸运数。

幸运数是指矩阵中满足同时下列两个条件的元素：

* 在同一行的所有元素中最小
* 在同一列的所有元素中最大



示例 1：

	输入：matrix = [[3,7,8],[9,11,13],[15,16,17]]
	输出：[15]
	解释：15 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。

示例 2：

	输入：matrix = [[1,10,4,2],[9,3,8,7],[15,16,17,12]]
	输出：[12]
	解释：12 是唯一的幸运数，因为它是其所在行中的最小值，也是所在列中的最大值。

示例 3：

	输入：matrix = [[7,8],[1,2]]
	输出：[7]



提示：

* m == mat.length
* n == mat[i].length
* 1 <= n, m <= 50
* 1 <= matrix[i][j] <= 10^5
* 矩阵中的所有元素都是不同的

