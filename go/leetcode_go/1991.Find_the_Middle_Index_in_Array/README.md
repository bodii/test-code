### 1991. Find the Middle Index in Array
Given a 0-indexed integer array nums, find the leftmost middleIndex (i.e., the smallest amongst all the possible ones).

A middleIndex is an index where nums[0] + nums[1] + ... + nums[middleIndex-1] == nums[middleIndex+1] + nums[middleIndex+2] + ... + nums[nums.length-1].

If middleIndex == 0, the left side sum is considered to be 0. Similarly, if middleIndex == nums.length - 1, the right side sum is considered to be 0.

Return the leftmost middleIndex that satisfies the condition, or -1 if there is no such index.



Example 1:

	Input: nums = [2,3,-1,8,4]
	Output: 3
	Explanation:
	The sum of the numbers before index 3 is: 2 + 3 + -1 = 4
	The sum of the numbers after index 3 is: 4 = 4

Example 2:

	Input: nums = [1,-1,4]
	Output: 2
	Explanation:
	The sum of the numbers before index 2 is: 1 + -1 = 0
	The sum of the numbers after index 2 is: 0

Example 3:

	Input: nums = [2,5]
	Output: -1
	Explanation:
	There is no valid middleIndex.

Example 4:

	Input: nums = [1]
	Output: 0
	Explantion:
	The sum of the numbers before index 0 is: 0
	The sum of the numbers after index 0 is: 0



Constraints:

* 1 <= nums.length <= 100
* -1000 <= nums[i] <= 1000



Note: This question is the same as 724: https://leetcode.com/problems/find-pivot-index/

----

### 1991. 找到数组的中间位置
给你一个下标从 0 开始的整数数组 nums ，请你找到 最左边 的中间位置 middleIndex （也就是所有可能中间位置下标最小的一个）。

中间位置 middleIndex 是满足 nums[0] + nums[1] + ... + nums[middleIndex-1] == nums[middleIndex+1] + nums[middleIndex+2] + ... + nums[nums.length-1] 的数组下标。

如果 middleIndex == 0 ，左边部分的和定义为 0 。类似的，如果 middleIndex == nums.length - 1 ，右边部分的和定义为 0 。

请你返回满足上述条件 最左边 的 middleIndex ，如果不存在这样的中间位置，请你返回 -1 。



示例 1：

	输入：nums = [2,3,-1,8,4]
	输出：3
	解释：
	下标 3 之前的数字和为：2 + 3 + -1 = 4
	下标 3 之后的数字和为：4 = 4

示例 2：

	输入：nums = [1,-1,4]
	输出：2
	解释：
	下标 2 之前的数字和为：1 + -1 = 0
	下标 2 之后的数字和为：0

示例 3：

	输入：nums = [2,5]
	输出：-1
	解释：
	不存在符合要求的 middleIndex 。

示例 4：

	输入：nums = [1]
	输出：0
	解释：
	下标 0 之前的数字和为：0
	下标 0 之后的数字和为：0



提示：

* 1 <= nums.length <= 100
* -1000 <= nums[i] <= 1000



注意：本题与主站 724 题相同：https://leetcode-cn.com/problems/find-pivot-index/

