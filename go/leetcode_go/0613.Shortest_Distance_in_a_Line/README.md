### 613. Shortest Distance in a Line
Table: Point

+-------------+------+
| Column Name | Type |
+-------------+------+
| x           | int  |
+-------------+------+
x is the primary key column for this table.
Each row of this table indicates the position of a point on the X-axis.



Write an SQL query to report the shortest distance between any two points from the Point table.

The query result format is in the following example.



Example 1:

Input:
Point table:
+----+
| x  |
+----+
| -1 |
| 0  |
| 2  |
+----+
Output:
+----------+
| shortest |
+----------+
| 1        |
+----------+
Explanation: The shortest distance is between points -1 and 0 which is |(-1) - 0| = 1.



Follow up: How could you optimize your query if the Point table is ordered in ascending order?

----

### 613. 直线上的最近距离
表 point 保存了一些点在 x 轴上的坐标，这些坐标都是整数。



写一个查询语句，找到这些点中最近两个点之间的距离。



| x   |
|-----|
| -1  |
| 0   |
| 2   |



最近距离显然是 '1' ，是点 '-1' 和 '0' 之间的距离。所以输出应该如下：



| shortest|
|---------|
| 1       |



注意：每个点都与其他点坐标不同，表 table 不会有重复坐标出现。



进阶：如果这些点在 x 轴上从左到右都有一个编号，输出结果时需要输出最近点对的编号呢？

