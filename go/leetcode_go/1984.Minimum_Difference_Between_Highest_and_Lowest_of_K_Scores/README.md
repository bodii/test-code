### 1984. Minimum Difference Between Highest and Lowest of K Scores
You are given a 0-indexed integer array nums, where nums[i] represents the score of the ith student. You are also given an integer k.

Pick the scores of any k students from the array so that the difference between the highest and the lowest of the k scores is minimized.

Return the minimum possible difference.

 

Example 1:

	Input: nums = [90], k = 1
	Output: 0
	Explanation: There is one way to pick score(s) of one student:
	- [90]. The difference between the highest and lowest score is 90 - 90 = 0.
	The minimum possible difference is 0.

Example 2:

	Input: nums = [9,4,1,7], k = 2
	Output: 2
	Explanation: There are six ways to pick score(s) of two students:
	- [9,4,1,7]. The difference between the highest and lowest score is 9 - 4 = 5.
	- [9,4,1,7]. The difference between the highest and lowest score is 9 - 1 = 8.
	- [9,4,1,7]. The difference between the highest and lowest score is 9 - 7 = 2.
	- [9,4,1,7]. The difference between the highest and lowest score is 4 - 1 = 3.
	- [9,4,1,7]. The difference between the highest and lowest score is 7 - 4 = 3.
	- [9,4,1,7]. The difference between the highest and lowest score is 7 - 1 = 6.
	The minimum possible difference is 2.

 

Constraints:

* 1 <= k <= nums.length <= 1000
* 0 <= nums[i] <= 10^5

----

### 1984. 学生分数的最小差值
给你一个 下标从 0 开始 的整数数组 nums ，其中 nums[i] 表示第 i 名学生的分数。另给你一个整数 k 。

从数组中选出任意 k 名学生的分数，使这 k 个分数间 最高分 和 最低分 的 差值 达到 最小化 。

返回可能的 最小差值 。

 

示例 1：

	输入：nums = [90], k = 1
	输出：0
	解释：选出 1 名学生的分数，仅有 1 种方法：
	- [90] 最高分和最低分之间的差值是 90 - 90 = 0
	可能的最小差值是 0

示例 2：

	输入：nums = [9,4,1,7], k = 2
	输出：2
	解释：选出 2 名学生的分数，有 6 种方法：
	- [9,4,1,7] 最高分和最低分之间的差值是 9 - 4 = 5
	- [9,4,1,7] 最高分和最低分之间的差值是 9 - 1 = 8
	- [9,4,1,7] 最高分和最低分之间的差值是 9 - 7 = 2
	- [9,4,1,7] 最高分和最低分之间的差值是 4 - 1 = 3
	- [9,4,1,7] 最高分和最低分之间的差值是 7 - 4 = 3
	- [9,4,1,7] 最高分和最低分之间的差值是 7 - 1 = 6
	可能的最小差值是 2

 

提示：

* 1 <= k <= nums.length <= 1000
* 0 <= nums[i] <= 10^5

