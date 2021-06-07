### 1399. Count Largest Group
Given an integer n. Each number from 1 to n is grouped according to the sum of its digits.

Return how many groups have the largest size.



Example 1:

	Input: n = 13
	Output: 4
	Explanation: There are 9 groups in total, they are grouped according sum of its digits of numbers from 1 to 13:
	[1,10], [2,11], [3,12], [4,13], [5], [6], [7], [8], [9]. There are 4 groups with largest size.

Example 2:

	Input: n = 2
	Output: 2
	Explanation: There are 2 groups [1], [2] of size 1.

Example 3:

	Input: n = 15
	Output: 6

Example 4:

	Input: n = 24
	Output: 5



Constraints:

* 1 <= n <= 10^4

----

### 1399. 统计最大组的数目
给你一个整数 n 。请你先求出从 1 到 n 的每个整数 10 进制表示下的数位和（每一位上的数字相加），然后把数位和相等的数字放到同一个组中。

请你统计每个组中的数字数目，并返回数字数目并列最多的组有多少个。



示例 1：

	输入：n = 13
	输出：4
	解释：总共有 9 个组，将 1 到 13 按数位求和后这些组分别是：
	[1,10]，[2,11]，[3,12]，[4,13]，[5]，[6]，[7]，[8]，[9]。总共有 4 个组拥有的数字并列最多。

示例 2：

	输入：n = 2
	输出：2
	解释：总共有 2 个大小为 1 的组 [1]，[2]。

示例 3：

	输入：n = 15
	输出：6

示例 4：

	输入：n = 24
	输出：5



提示：

* 1 <= n <= 10^4

