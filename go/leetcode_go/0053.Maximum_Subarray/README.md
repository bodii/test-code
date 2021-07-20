### 53. Maximum Subarray
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.



Example 1:

	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: [4,-1,2,1] has the largest sum = 6.

Example 2:

	Input: nums = [1]
	Output: 1

Example 3:

	Input: nums = [5,4,-1,7,8]
	Output: 23



Constraints:

* 1 <= nums.length <= 3 * 104
* -105 <= nums[i] <= 105


Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.


----

### 53. 最大子序和
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。



示例 1：

	输入：nums = [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6 。

示例 2：

	输入：nums = [1]
	输出：1

示例 3：

	输入：nums = [0]
	输出：0

示例 4：

	输入：nums = [-1]
	输出：-1

示例 5：

	输入：nums = [-100000]
	输出：-100000



提示：

* 1 <= nums.length <= 3 * 104
* -105 <= nums[i] <= 105



进阶：如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的 分治法 求解。

