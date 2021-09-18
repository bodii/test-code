### 724. Find Pivot Index
Given an array of integers nums, calculate the pivot index of this array.

The pivot index is the index where the sum of all the numbers strictly to the left of the index is equal to the sum of all the numbers strictly to the index's right.

If the index is on the left edge of the array, then the left sum is 0 because there are no elements to the left. This also applies to the right edge of the array.

Return the leftmost pivot index. If no such index exists, return -1.



Example 1:

	Input: nums = [1,7,3,6,5,6]
	Output: 3
	Explanation:
	The pivot index is 3.
	Left sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11
	Right sum = nums[4] + nums[5] = 5 + 6 = 11

Example 2:

	Input: nums = [1,2,3]
	Output: -1
	Explanation:
	There is no index that satisfies the conditions in the problem statement.

Example 3:

	Input: nums = [2,1,-1]
	Output: 0
	Explanation:
	The pivot index is 0.
	Left sum = 0 (no elements to the left of index 0)
	Right sum = nums[1] + nums[2] = 1 + -1 = 0



Constraints:

* 1 <= nums.length <= 10^4
* -1000 <= nums[i] <= 1000



Note: This question is the same as 1991: https://leetcode.com/problems/find-the-middle-index-in-array/

----

### 724. 寻找数组的中心下标
给你一个整数数组 nums ，请计算数组的 中心下标 。

数组 中心下标 是数组的一个下标，其左侧所有元素相加的和等于右侧所有元素相加的和。

如果中心下标位于数组最左端，那么左侧数之和视为 0 ，因为在下标的左侧不存在元素。这一点对于中心下标位于数组最右端同样适用。

如果数组有多个中心下标，应该返回 最靠近左边 的那一个。如果数组不存在中心下标，返回 -1 。



示例 1：

	输入：nums = [1, 7, 3, 6, 5, 6]
	输出：3
	解释：
	中心下标是 3 。
	左侧数之和 sum = nums[0] + nums[1] + nums[2] = 1 + 7 + 3 = 11 ，
	右侧数之和 sum = nums[4] + nums[5] = 5 + 6 = 11 ，二者相等。

示例 2：

	输入：nums = [1, 2, 3]
	输出：-1
	解释：
	数组中不存在满足此条件的中心下标。

示例 3：

	输入：nums = [2, 1, -1]
	输出：0
	解释：
	中心下标是 0 。
	左侧数之和 sum = 0 ，（下标 0 左侧不存在元素），
	右侧数之和 sum = nums[1] + nums[2] = 1 + -1 = 0 。



提示：

* 1 <= nums.length <= 10^4
* -1000 <= nums[i] <= 1000
