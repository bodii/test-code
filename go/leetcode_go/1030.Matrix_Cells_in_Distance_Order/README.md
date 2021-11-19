### 1030. Matrix Cells in Distance Order
You are given four integers row, cols, rCenter, and cCenter. There is a rows x cols matrix and you are on the cell with the coordinates (rCenter, cCenter).

Return the coordinates of all cells in the matrix, sorted by their distance from (rCenter, cCenter) from the smallest distance to the largest distance. You may return the answer in any order that satisfies this condition.

The distance between two cells (r1, c1) and (r2, c2) is |r1 - r2| + |c1 - c2|.



Example 1:

	Input: rows = 1, cols = 2, rCenter = 0, cCenter = 0
	Output: [[0,0],[0,1]]
	Explanation: The distances from (0, 0) to other cells are: [0,1]

Example 2:

	Input: rows = 2, cols = 2, rCenter = 0, cCenter = 1
	Output: [[0,1],[0,0],[1,1],[1,0]]
	Explanation: The distances from (0, 1) to other cells are: [0,1,1,2]
	The answer [[0,1],[1,1],[0,0],[1,0]] would also be accepted as correct.

Example 3:

	Input: rows = 2, cols = 3, rCenter = 1, cCenter = 2
	Output: [[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
	Explanation: The distances from (1, 2) to other cells are: [0,1,1,2,2,3]
	There are other answers that would also be accepted as correct, such as [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]].



Constraints:

* 1 <= rows, cols <= 100
* 0 <= rCenter < rows
* 0 <= cCenter < cols

----

### 1030. 距离顺序排列矩阵单元格
给出 R 行 C 列的矩阵，其中的单元格的整数坐标为 (r, c)，满足 0 <= r < R 且 0 <= c < C。

另外，我们在该矩阵中给出了一个坐标为 (r0, c0) 的单元格。

返回矩阵中的所有单元格的坐标，并按到 (r0, c0) 的距离从最小到最大的顺序排，其中，两单元格(r1, c1) 和 (r2, c2) 之间的距离是曼哈顿距离，|r1 - r2| + |c1 - c2|。（你可以按任何满足此条件的顺序返回答案。）



示例 1：

	输入：R = 1, C = 2, r0 = 0, c0 = 0
	输出：[[0,0],[0,1]]
	解释：从 (r0, c0) 到其他单元格的距离为：[0,1]

示例 2：

	输入：R = 2, C = 2, r0 = 0, c0 = 1
	输出：[[0,1],[0,0],[1,1],[1,0]]
	解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2]
	[[0,1],[1,1],[0,0],[1,0]] 也会被视作正确答案。

示例 3：

	输入：R = 2, C = 3, r0 = 1, c0 = 2
	输出：[[1,2],[0,2],[1,1],[0,1],[1,0],[0,0]]
	解释：从 (r0, c0) 到其他单元格的距离为：[0,1,1,2,2,3]
	其他满足题目要求的答案也会被视为正确，例如 [[1,2],[1,1],[0,2],[1,0],[0,1],[0,0]]。



提示：

* 1 <= R <= 100
* 1 <= C <= 100
* 0 <= r0 < R
* 0 <= c0 < C

