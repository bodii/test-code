### 1909. Remove One Element to Make the Array Strictly Increasing
Given a 0-indexed integer array nums, return true if it can be made strictly increasing after removing exactly one element, or false otherwise. If the array is already strictly increasing, return true.

The array nums is strictly increasing if nums[i - 1] < nums[i] for each index (1 <= i < nums.length).



Example 1:

	Input: nums = [1,2,10,5,7]
	Output: true
	Explanation: By removing 10 at index 2 from nums, it becomes [1,2,5,7].
	[1,2,5,7] is strictly increasing, so return true.

Example 2:

	Input: nums = [2,3,1,2]
	Output: false
	Explanation:
	[3,1,2] is the result of removing the element at index 0.
	[2,1,2] is the result of removing the element at index 1.
	[2,3,2] is the result of removing the element at index 2.
	[2,3,1] is the result of removing the element at index 3.
	No resulting array is strictly increasing, so return false.

Example 3:

	Input: nums = [1,1,1]
	Output: false
	Explanation: The result of removing any element is [1,1].
	[1,1] is not strictly increasing, so return false.



Constraints:

* 2 <= nums.length <= 1000
* 1 <= nums[i] <= 1000

----

### 1909. 删除一个元素使数组严格递增
给你一个下标从 0 开始的整数数组 nums ，如果 恰好 删除 一个 元素后，数组 严格递增 ，那么请你返回 true ，否则返回 false 。如果数组本身已经是严格递增的，请你也返回 true 。

数组 nums 是 严格递增 的定义为：对于任意下标的 1 <= i < nums.length 都满足 nums[i - 1] < nums[i] 。



示例 1：

	输入：nums = [1,2,10,5,7]
	输出：true
	解释：从 nums 中删除下标 2 处的 10 ，得到 [1,2,5,7] 。
	[1,2,5,7] 是严格递增的，所以返回 true 。

示例 2：

	输入：nums = [2,3,1,2]
	输出：false
	解释：
	[3,1,2] 是删除下标 0 处元素后得到的结果。
	[2,1,2] 是删除下标 1 处元素后得到的结果。
	[2,3,2] 是删除下标 2 处元素后得到的结果。
	[2,3,1] 是删除下标 3 处元素后得到的结果。
	没有任何结果数组是严格递增的，所以返回 false 。

示例 3：

	输入：nums = [1,1,1]
	输出：false
	解释：删除任意元素后的结果都是 [1,1] 。
	[1,1] 不是严格递增的，所以返回 false 。

示例 4：

	输入：nums = [1,2,3]
	输出：true
	解释：[1,2,3] 已经是严格递增的，所以返回 true 。



提示：

* 2 <= nums.length <= 1000
* 1 <= nums[i] <= 1000

