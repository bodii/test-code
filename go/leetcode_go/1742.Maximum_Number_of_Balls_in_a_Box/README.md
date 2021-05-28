### 1742. Maximum Number of Balls in a Box
You are working in a ball factory where you have n balls numbered from lowLimit up to highLimit inclusive (i.e., n == highLimit - lowLimit + 1), and an infinite number of boxes numbered from 1 to infinity.

Your job at this factory is to put each ball in the box with a number equal to the sum of digits of the ball's number. For example, the ball number 321 will be put in the box number 3 + 2 + 1 = 6 and the ball number 10 will be put in the box number 1 + 0 = 1.

Given two integers lowLimit and highLimit, return the number of balls in the box with the most balls.



Example 1:

	Input: lowLimit = 1, highLimit = 10
	Output: 2
	Explanation:
	Box Number:  1 2 3 4 5 6 7 8 9 10 11 ...
	Ball Count:  2 1 1 1 1 1 1 1 1 0  0  ...
	Box 1 has the most number of balls with 2 balls.

Example 2:

	Input: lowLimit = 5, highLimit = 15
	Output: 2
	Explanation:
	Box Number:  1 2 3 4 5 6 7 8 9 10 11 ...
	Ball Count:  1 1 1 1 2 2 1 1 1 0  0  ...
	Boxes 5 and 6 have the most number of balls with 2 balls in each.

Example 3:

	Input: lowLimit = 19, highLimit = 28
	Output: 2
	Explanation:
	Box Number:  1 2 3 4 5 6 7 8 9 10 11 12 ...
	Ball Count:  0 1 1 1 1 1 1 1 1 2  0  0  ...
	Box 10 has the most number of balls with 2 balls.



Constraints:

* 1 <= lowLimit <= highLimit <= 10<sup>5</sup>

----

### 1742. 盒子中小球的最大数量
你在一家生产小球的玩具厂工作，有 n 个小球，编号从 lowLimit 开始，到 highLimit 结束（包括 lowLimit 和 highLimit ，即 n == highLimit - lowLimit + 1）。另有无限数量的盒子，编号从 1 到 infinity 。

你的工作是将每个小球放入盒子中，其中盒子的编号应当等于小球编号上每位数字的和。例如，编号 321 的小球应当放入编号 3 + 2 + 1 = 6 的盒子，而编号 10 的小球应当放入编号 1 + 0 = 1 的盒子。

给你两个整数 lowLimit 和 highLimit ，返回放有最多小球的盒子中的小球数量。如果有多个盒子都满足放有最多小球，只需返回其中任一盒子的小球数量。



示例 1：

	输入：lowLimit = 1, highLimit = 10
	输出：2
	解释：
	盒子编号：1 2 3 4 5 6 7 8 9 10 11 ...
	小球数量：2 1 1 1 1 1 1 1 1 0  0  ...
	编号 1 的盒子放有最多小球，小球数量为 2 。

示例 2：

	输入：lowLimit = 5, highLimit = 15
	输出：2
	解释：
	盒子编号：1 2 3 4 5 6 7 8 9 10 11 ...
	小球数量：1 1 1 1 2 2 1 1 1 0  0  ...
	编号 5 和 6 的盒子放有最多小球，每个盒子中的小球数量都是 2 。

示例 3：

	输入：lowLimit = 19, highLimit = 28
	输出：2
	解释：
	盒子编号：1 2 3 4 5 6 7 8 9 10 11 12 ...
	小球数量：0 1 1 1 1 1 1 1 1 2  0  0  ...
	编号 10 的盒子放有最多小球，小球数量为 2 。



提示：

* 1 <= lowLimit <= highLimit <= 10<sup>5</sup>

