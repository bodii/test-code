### 88. Merge Sorted Array
You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n, representing the number of elements in nums1 and nums2 respectively.

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function, but instead be stored inside the array nums1. To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged, and the last n elements are set to 0 and should be ignored. nums2 has a length of n.



Example 1:

	Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	Output: [1,2,2,3,5,6]
	Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
	The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

Example 2:

	Input: nums1 = [1], m = 1, nums2 = [], n = 0
	Output: [1]
	Explanation: The arrays we are merging are [1] and [].
	The result of the merge is [1].

Example 3:

	Input: nums1 = [0], m = 0, nums2 = [1], n = 1
	Output: [1]
	Explanation: The arrays we are merging are [] and [1].
	The result of the merge is [1].
	Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.



Constraints:

* nums1.length == m + n
* nums2.length == n
* 0 <= m, n <= 200
* 1 <= m + n <= 200
* -109 <= nums1[i], nums2[j] <= 10^9


Follow up: Can you come up with an algorithm that runs in O(m + n) time?


----

### 88. 合并两个有序数组
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。

初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。你可以假设 nums1 的空间大小等于 m + n，这样它就有足够的空间保存来自 nums2 的元素。



示例 1：

	输入：nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
	输出：[1,2,2,3,5,6]

示例 2：

	输入：nums1 = [1], m = 1, nums2 = [], n = 0
	输出：[1]



提示：

* nums1.length == m + n
* nums2.length == n
* 0 <= m, n <= 200
* 1 <= m + n <= 200
* -109 <= nums1[i], nums2[i] <= 10^9

