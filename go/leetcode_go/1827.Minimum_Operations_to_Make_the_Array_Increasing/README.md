### 1827. Minimum Operations to Make the Array Increasing
You are given an integer array nums (0-indexed). In one operation, you can choose an element of the array and increment it by 1.

* For example, if nums = [1,2,3], you can choose to increment nums[1] to make nums = [1,3,3].

Return the minimum number of operations needed to make nums strictly increasing.

An array nums is strictly increasing if nums[i] < nums[i+1] for all 0 <= i < nums.length - 1. An array of length 1 is trivially strictly increasing.



Example 1:

	Input: nums = [1,1,1]
	Output: 3
	Explanation: You can do the following operations:
	1) Increment nums[2], so nums becomes [1,1,2].
	2) Increment nums[1], so nums becomes [1,2,2].
	3) Increment nums[2], so nums becomes [1,2,3].

Example 2:

	Input: nums = [1,5,2,4,1]
	Output: 14

Example 3:

	Input: nums = [8]
	Output: 0



Constraints:

* 1 <= nums.length <= 5000
* 1 <= nums[i] <= 104

----

### 1827. 最少操作使数组递增
给你一个整数数组 nums （下标从 0 开始）。每一次操作中，你可以选择数组中一个元素，并将它增加 1 。

* 比方说，如果 nums = [1,2,3] ，你可以选择增加 nums[1] 得到 nums = [1,3,3] 。

请你返回使 nums 严格递增 的 最少 操作次数。

我们称数组 nums 是 严格递增的 ，当它满足对于所有的 0 <= i < nums.length - 1 都有 nums[i] < nums[i+1] 。一个长度为 1 的数组是严格递增的一种特殊情况。



示例 1：

	输入：nums = [1,1,1]
	输出：3
	解释：你可以进行如下操作：
	1) 增加 nums[2] ，数组变为 [1,1,2] 。
	2) 增加 nums[1] ，数组变为 [1,2,2] 。
	3) 增加 nums[2] ，数组变为 [1,2,3] 。

示例 2：

	输入：nums = [1,5,2,4,1]
	输出：14

示例 3：

	输入：nums = [8]
	输出：0



提示：

* 1 <= nums.length <= 5000
* 1 <= nums[i] <= 104

