### 219. Contains Duplicate II
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.



Example 1:

	Input: nums = [1,2,3,1], k = 3
	Output: true

Example 2:

	Input: nums = [1,0,1,1], k = 1
	Output: true

Example 3:

	Input: nums = [1,2,3,1,2,3], k = 2
	Output: false



Constraints:

* 1 <= nums.length <= 10^5
* -10^9 <= nums[i] <= 10^9
* 0 <= k <= 10^5

----

### 219. 存在重复元素 II
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的 绝对值 至多为 k。



示例 1:

	输入: nums = [1,2,3,1], k = 3
	输出: true

示例 2:

	输入: nums = [1,0,1,1], k = 1
	输出: true

示例 3:

	输入: nums = [1,2,3,1,2,3], k = 2
	输出: false

