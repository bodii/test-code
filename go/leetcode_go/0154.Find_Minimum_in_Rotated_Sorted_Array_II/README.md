### 154. Find Minimum in Rotated Sorted Array II
Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,4,4,5,6,7] might become:

* [4,5,6,7,0,1,4] if it was rotated 4 times.
* [0,1,4,4,5,6,7] if it was rotated 7 times.

Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.

You must decrease the overall operation steps as much as possible.



Example 1:

	Input: nums = [1,3,5]
	Output: 1

Example 2:

	Input: nums = [2,2,2,0,1]
	Output: 0



Constraints:

* n == nums.length
* 1 <= n <= 5000
* -5000 <= nums[i] <= 5000
* nums is sorted and rotated between 1 and n times.



Follow up: This problem is similar to Find Minimum in Rotated Sorted Array, but nums may contain duplicates. Would this affect the runtime complexity? How and why?

----
### 154. 寻找旋转排序数组中的最小值 II
已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：

* 若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
* 若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]

注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。

给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。



示例 1：

	输入：nums = [1,3,5]
	输出：1

示例 2：

	输入：nums = [2,2,2,0,1]
	输出：0



提示：

* n == nums.length
* 1 <= n <= 5000
* -5000 <= nums[i] <= 5000
* nums 原来是一个升序排序的数组，并进行了 1 至 n 次旋转



进阶：

* 这道题是 寻找旋转排序数组中的最小值 的延伸题目。
* 允许重复会影响算法的时间复杂度吗？会如何影响，为什么？

