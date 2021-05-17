### 1512. Number of Good Pairs
Given an array of integers nums.

A pair (i,j) is called good if nums[i] == nums[j] and i < j.

Return the number of good pairs.



Example 1:

	Input: nums = [1,2,3,1,1,3]
	Output: 4
	Explanation: There are 4 good pairs (0,3), (0,4), (3,4), (2,5) 0-indexed.

Example 2:

	Input: nums = [1,1,1,1]
	Output: 6
	Explanation: Each pair in the array are good.

Example 3:

	Input: nums = [1,2,3]
	Output: 0



Constraints:

* 1 <= nums.length <= 100
* 1 <= nums[i] <= 100

----

### 1512. 好数对的数目
给你一个整数数组 nums 。

如果一组数字 (i,j) 满足 nums[i] == nums[j] 且 i < j ，就可以认为这是一组 好数对 。

返回好数对的数目。



示例 1：

	输入：nums = [1,2,3,1,1,3]
	输出：4
	解释：有 4 组好数对，分别是 (0,3), (0,4), (3,4), (2,5) ，下标从 0 开始

示例 2：

	输入：nums = [1,1,1,1]
	输出：6
	解释：数组中的每组数字都是好数对

示例 3：

	输入：nums = [1,2,3]
	输出：0



提示：

* 1 <= nums.length <= 100
* 1 <= nums[i] <= 100
