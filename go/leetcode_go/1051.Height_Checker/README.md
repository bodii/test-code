### 1051. Height Checker
A school is trying to take an annual photo of all the students. The students are asked to stand in a single file line in non-decreasing order by height. Let this ordering be represented by the integer array expected where expected[i] is the expected height of the ith student in line.

You are given an integer array heights representing the current order that the students are standing in. Each heights[i] is the height of the ith student in line (0-indexed).

Return the number of indices where heights[i] != expected[i].



 Example 1:

	 Input: heights = [1,1,4,2,1,3]
	 Output: 3
	 Explanation:
	 heights:  [1,1,4,2,1,3]
	 expected: [1,1,1,2,3,4]
	 Indices 2, 4, and 5 do not match.

 Example 2:

	 Input: heights = [5,1,2,3,4]
	 Output: 5
	 Explanation:
	 heights:  [5,1,2,3,4]
	 expected: [1,2,3,4,5]
	 All indices do not match.

 Example 3:

	 Input: heights = [1,2,3,4,5]
	 Output: 0
	 Explanation:
	 heights:  [1,2,3,4,5]
	 expected: [1,2,3,4,5]
	 All indices match.



  Constraints:

*   1 <= heights.length <= 100
*   1 <= heights[i] <= 100

----

### 1051. 高度检查器
学校在拍年度纪念照时，一般要求学生按照 非递减 的高度顺序排列。

请你返回能让所有学生以 非递减 高度排列的最小必要移动人数。

注意，当一组学生被选中时，他们之间可以以任何可能的方式重新排序，而未被选中的学生应该保持不动。



 示例：

	 输入：heights = [1,1,4,2,1,3]
	 输出：3
	 解释：
	 当前数组：[1,1,4,2,1,3]
	 目标数组：[1,1,1,2,3,4]
	 在下标 2 处（从 0 开始计数）出现 4 vs 1 ，所以我们必须移动这名学生。
	 在下标 4 处（从 0 开始计数）出现 1 vs 3 ，所以我们必须移动这名学生。
	 在下标 5 处（从 0 开始计数）出现 3 vs 4 ，所以我们必须移动这名学生。

 示例 2：

	 输入：heights = [5,1,2,3,4]
	 输出：5

 示例 3：

	 输入：heights = [1,2,3,4,5]
	 输出：0



  提示：

* 1 <= heights.length <= 100
* 1 <= heights[i] <= 100
