### 1725. Number Of Rectangles That Can Form The Largest Square
You are given an array rectangles where rectangles[i] = [$l_i$, $w_i$] represents the ith rectangle of length $l_i$ and width $w_i$.

You can cut the ith rectangle to form a square with a side length of k if both k <= $l_i$ and k <= $w_i$. For example, if you have a rectangle [4,6], you can cut it to get a square with a side length of at most 4.

Let maxLen be the side length of the largest square you can obtain from any of the given rectangles.

Return the number of rectangles that can make a square with a side length of maxLen.



Example 1:

	Input: rectangles = [[5,8],[3,9],[5,12],[16,5]]
	Output: 3
	Explanation: The largest squares you can get from each rectangle are of lengths [5,3,5,5].
	The largest possible square is of length 5, and you can get it out of 3 rectangles.

Example 2:

	Input: rectangles = [[2,3],[3,7],[4,3],[3,7]]
	Output: 3



Constraints:

* 1 <= rectangles.length <= 1000
* rectangles[i].length == 2
* 1 <= $l_i$, $w_i$ <= $10^9$
* $l_i$ != $w_i$

----

### 1725. 可以形成最大正方形的矩形数目
给你一个数组 rectangles ，其中 rectangles[i] = [$l_i$, $w_i$] 表示第 i 个矩形的长度为 $l_i$ 、宽度为 $w_i$ 。

如果存在 k 同时满足 k <= $l_i$ 和 k <= $w_i$ ，就可以将第 i 个矩形切成边长为 k 的正方形。例如，矩形 [4,6] 可以切成边长最大为 4 的正方形。

设 maxLen 为可以从矩形数组 rectangles 切分得到的 最大正方形 的边长。

请你统计有多少个矩形能够切出边长为 maxLen 的正方形，并返回矩形 数目 。



示例 1：

	输入：rectangles = [[5,8],[3,9],[5,12],[16,5]]
	输出：3
	解释：能从每个矩形中切出的最大正方形边长分别是 [5,3,5,5] 。
	最大正方形的边长为 5 ，可以由 3 个矩形切分得到。

示例 2：

	输入：rectangles = [[2,3],[3,7],[4,3],[3,7]]
	输出：3



提示：

* 1 <= rectangles.length <= 1000
* rectangles[i].length == 2
* 1 <= $l_i$, $w_i$ <= $10^9$
* $l_i$ != $w_i$

