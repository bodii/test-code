### 228. Summary Ranges
You are given a sorted unique integer array nums.

Return the smallest sorted list of ranges that cover all the numbers in the array exactly. That is, each element of nums is covered by exactly one of the ranges, and there is no integer x such that x is in one of the ranges but not in nums.

Each range [a,b] in the list should be output as:

    "a->b" if a != b
    "a" if a == b



Example 1:

	Input: nums = [0,1,2,4,5,7]
	Output: ["0->2","4->5","7"]
	Explanation: The ranges are:
	[0,2] --> "0->2"
	[4,5] --> "4->5"
	[7,7] --> "7"

Example 2:

	Input: nums = [0,2,3,4,6,8,9]
	Output: ["0","2->4","6","8->9"]
	Explanation: The ranges are:
	[0,0] --> "0"
	[2,4] --> "2->4"
	[6,6] --> "6"
	[8,9] --> "8->9"

Example 3:

	Input: nums = []
	Output: []

Example 4:

	Input: nums = [-1]
	Output: ["-1"]

Example 5:

	Input: nums = [0]
	Output: ["0"]



Constraints:

* 0 <= nums.length <= 20
* -2<sup>31</sup> <= nums[i] <= 2<sup>31</sup> - 1
* All the values of nums are unique.
* nums is sorted in ascending order.

----

### 228. 汇总区间
给定一个无重复元素的有序整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

    "a->b" ，如果 a != b
    "a" ，如果 a == b



示例 1：

	输入：nums = [0,1,2,4,5,7]
	输出：["0->2","4->5","7"]
	解释：区间范围是：
	[0,2] --> "0->2"
	[4,5] --> "4->5"
	[7,7] --> "7"

示例 2：

	输入：nums = [0,2,3,4,6,8,9]
	输出：["0","2->4","6","8->9"]
	解释：区间范围是：
	[0,0] --> "0"
	[2,4] --> "2->4"
	[6,6] --> "6"
	[8,9] --> "8->9"

示例 3：

	输入：nums = []
	输出：[]

示例 4：

	输入：nums = [-1]
	输出：["-1"]

示例 5：

	输入：nums = [0]
	输出：["0"]



提示：

* 0 <= nums.length <= 20
* -2<sup>31</sup> <= nums[i] <= 2<sup>31</sup> - 1
* nums 中的所有值都 互不相同
* nums 按升序排列

