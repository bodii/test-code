### 2026. Low-Quality Problems
Table: Problems

+-------------+------+
| Column Name | Type |
+-------------+------+
| problem_id  | int  |
| likes       | int  |
| dislikes    | int  |
+-------------+------+
problem_id is the primary key column for this table.
Each row of this table indicates the number of likes and dislikes for a LeetCode problem.



Write an SQL query to report the IDs of the low-quality problems. A LeetCode problem is low-quality if the like percentage of the problem (number of likes divided by the total number of votes) is strictly less than 60%.

Return the result table ordered by problem_id in ascending order.

The query result format is in the following example.



Example 1:

Input:
Problems table:
+------------+-------+----------+
| problem_id | likes | dislikes |
+------------+-------+----------+
| 6          | 1290  | 425      |
| 11         | 2677  | 8659     |
| 1          | 4446  | 2760     |
| 7          | 8569  | 6086     |
| 13         | 2050  | 4164     |
| 10         | 9002  | 7446     |
+------------+-------+----------+
Output:
+------------+
| problem_id |
+------------+
| 7          |
| 10         |
| 11         |
| 13         |
+------------+
Explanation: The like percentages are as follows:
- Problem 1: (4446 / (4446 + 2760)) * 100 = 61.69858%
- Problem 6: (1290 / (1290 + 425)) * 100 = 75.21866%
- Problem 7: (8569 / (8569 + 6086)) * 100 = 58.47151%
- Problem 10: (9002 / (9002 + 7446)) * 100 = 54.73006%
- Problem 11: (2677 / (2677 + 8659)) * 100 = 23.61503%
- Problem 13: (2050 / (2050 + 4164)) * 100 = 32.99002%
Problems 7, 10, 11, and 13 are low-quality problems because their like percentages are less than 60%.

----

### 2026. 低质量的问题
表： Problems

+-------------+------+
| 列名         | 类型 |
+-------------+------+
| problem_id  | int  |
| likes       | int  |
| dislikes    | int  |
+-------------+------+
problem_id 是这张表的主键。
该表的每一行都表示一个力扣问题的喜欢和不喜欢的数量。

写一个 SQL 查询低质量问题的 ID 集合。如果一个力扣问题的喜欢率（喜欢数除以总投票数）严格低于60% ，则该问题为低质量问题。

按 problem_id 升序排列返回结果表。

查询结果的格式在下面的例子中。



示例 1:

输入:
Problems 表:
+------------+-------+----------+
| problem_id | likes | dislikes |
+------------+-------+----------+
| 6          | 1290  | 425      |
| 11         | 2677  | 8659     |
| 1          | 4446  | 2760     |
| 7          | 8569  | 6086     |
| 13         | 2050  | 4164     |
| 10         | 9002  | 7446     |
+------------+-------+----------+
输出:
+------------+
| problem_id |
+------------+
| 7          |
| 10         |
| 11         |
| 13         |
+------------+
解释: 喜欢的比率如下:
- 问题 1: (4446 / (4446 + 2760)) * 100 = 61.69858%
- 问题 6: (1290 / (1290 + 425)) * 100 = 75.21866%
- 问题 7: (8569 / (8569 + 6086)) * 100 = 58.47151%
- 问题 10: (9002 / (9002 + 7446)) * 100 = 54.73006%
- 问题 11: (2677 / (2677 + 8659)) * 100 = 23.61503%
- 问题 13: (2050 / (2050 + 4164)) * 100 = 32.99002%
问题 7, 10, 11, 和 13 是低质量问题，因为它们的同类百分比低于60%。

