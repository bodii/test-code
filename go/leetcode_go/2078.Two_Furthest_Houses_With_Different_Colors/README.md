### 2078. Two Furthest Houses With Different Colors
There are n houses evenly lined up on the street, and each house is beautifully painted. You are given a 0-indexed integer array colors of length n, where colors[i] represents the color of the ith house.

Return the maximum distance between two houses with different colors.

The distance between the ith and jth houses is abs(i - j), where abs(x) is the absolute value of x.



Example 1:

	Input: colors = [1,1,1,6,1,1,1]
	Output: 3
	Explanation: In the above image, color 1 is blue, and color 6 is red.
	The furthest two houses with different colors are house 0 and house 3.
	House 0 has color 1, and house 3 has color 6. The distance between them is abs(0 - 3) = 3.
	Note that houses 3 and 6 can also produce the optimal answer.

Example 2:

	Input: colors = [1,8,3,8,3]
	Output: 4
	Explanation: In the above image, color 1 is blue, color 8 is yellow, and color 3 is green.
	The furthest two houses with different colors are house 0 and house 4.
	House 0 has color 1, and house 4 has color 3. The distance between them is abs(0 - 4) = 4.

Example 3:

	Input: colors = [0,1]
	Output: 1
	Explanation: The furthest two houses with different colors are house 0 and house 1.
	House 0 has color 0, and house 1 has color 1. The distance between them is abs(0 - 1) = 1.



Constraints:

* n == colors.length
* 2 <= n <= 100
* 0 <= colors[i] <= 100
* Test data are generated such that at least two houses have different colors.

----

### 2078. 两栋颜色不同且距离最远的房子
街上有 n 栋房子整齐地排成一列，每栋房子都粉刷上了漂亮的颜色。给你一个下标从 0 开始且长度为 n 的整数数组 colors ，其中 colors[i] 表示第  i 栋房子的颜色。

返回 两栋 颜色 不同 房子之间的 最大 距离。

第 i 栋房子和第 j 栋房子之间的距离是 abs(i - j) ，其中 abs(x) 是 x 的绝对值。



示例 1：

	输入：colors = [1,1,1,6,1,1,1]
	输出：3
	解释：上图中，颜色 1 标识成蓝色，颜色 6 标识成红色。
	两栋颜色不同且距离最远的房子是房子 0 和房子 3 。
	房子 0 的颜色是颜色 1 ，房子 3 的颜色是颜色 6 。两栋房子之间的距离是 abs(0 - 3) = 3 。
	注意，房子 3 和房子 6 也可以产生最佳答案。

示例 2：

	输入：colors = [1,8,3,8,3]
	输出：4
	解释：上图中，颜色 1 标识成蓝色，颜色 8 标识成黄色，颜色 3 标识成绿色。
	两栋颜色不同且距离最远的房子是房子 0 和房子 4 。
	房子 0 的颜色是颜色 1 ，房子 4 的颜色是颜色 3 。两栋房子之间的距离是 abs(0 - 4) = 4 。

示例 3：

	输入：colors = [0,1]
	输出：1
	解释：两栋颜色不同且距离最远的房子是房子 0 和房子 1 。
	房子 0 的颜色是颜色 0 ，房子 1 的颜色是颜色 1 。两栋房子之间的距离是 abs(0 - 1) = 1 。



提示：

* n == colors.length
* 2 <= n <= 100
* 0 <= colors[i] <= 100
* 生成的测试数据满足 至少 存在 2 栋颜色不同的房子


