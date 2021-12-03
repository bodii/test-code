### Search in Rotated Sorted Array II
There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

You must decrease the overall operation steps as much as possible.



Example 1:

	Input: nums = [2,5,6,0,0,1,2], target = 0
	Output: true

Example 2:

	Input: nums = [2,5,6,0,0,1,2], target = 3
	Output: false



Constraints:

    1 <= nums.length <= 5000
    -10^4 <= nums[i] <= 10^4
    nums is guaranteed to be rotated at some pivot.
    -10^4 <= target <= 10^4



Follow up: This problem is similar to Search in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?

----
### 81. 搜索旋转排序数组 II
已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。

在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。

给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。



示例 1：

	输入：nums = [2,5,6,0,0,1,2], target = 0
	输出：true

示例 2：

	输入：nums = [2,5,6,0,0,1,2], target = 3
	输出：false



提示：

* 1 <= nums.length <= 5000
* -10^4 <= nums[i] <= 10^4
* 题目数据保证 nums 在预先未知的某个下标上进行了旋转
* -10^4 <= target <= 10^4



进阶：

* 这是 搜索旋转排序数组 的延伸题目，本题中的 nums  可能包含重复元素。
* 这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？

