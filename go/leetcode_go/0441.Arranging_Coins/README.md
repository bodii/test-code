### 441. Arranging Coins
You have n coins and you want to build a staircase with these coins. The staircase consists of k rows where the ith row has exactly i coins. The last row of the staircase may be incomplete.

Given the integer n, return the number of complete rows of the staircase you will build.



Example 1:

	Input: n = 5
	Output: 2
	Explanation: Because the 3rd row is incomplete, we return 2.

Example 2:

	Input: n = 8
	Output: 3
	Explanation: Because the 4th row is incomplete, we return 3.



Constraints:

* 1 <= n <= 2<sup>31</sup> - 1

----

### 441. 排列硬币
你总共有 n 枚硬币，你需要将它们摆成一个阶梯形状，第 k 行就必须正好有 k 枚硬币。

给定一个数字 n，找出可形成完整阶梯行的总行数。

n 是一个非负整数，并且在32位有符号整型的范围内。

示例 1:

	n = 5

	硬币可排列成以下几行:
	¤
	¤ ¤
	¤ ¤

	因为第三行不完整，所以返回2.

示例 2:

	n = 8

	硬币可排列成以下几行:
	¤
	¤ ¤
	¤ ¤ ¤
	¤ ¤

	因为第四行不完整，所以返回3.


