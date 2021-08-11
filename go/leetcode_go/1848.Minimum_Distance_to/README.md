### 1848. Minimum Distance to the Target Element

Given an integer array nums (0-indexed) and two integers target and start, find an index i such that nums[i] == target and abs(i - start) is minimized. Note that abs(x) is the absolute value of x.

Return abs(i - start).

It is guaranteed that target exists in nums.

Example 1:

    Input: nums = [1,2,3,4,5], target = 5, start = 3
    Output: 1
    Explanation: nums[4] = 5 is the only value equal to target, so the answer is abs(4 - 3) = 1.

Example 2:

    Input: nums = [1], target = 1, start = 0
    Output: 0
    Explanation: nums[0] = 1 is the only value equal to target, so the answer is abs(0 - 0) = 0.

Example 3:

    Input: nums = [1,1,1,1,1,1,1,1,1,1], target = 1, start = 0
    Output: 0
    Explanation: Every value of nums is 1, but nums[0] minimizes abs(i - start), which is abs(0 - 0) = 0.

Constraints:

- 1 <= nums.length <= 1000
- 1 <= nums[i] <= 10^4
- 0 <= start < nums.length
- target is in nums.

---

### 1848. 到目标元素的最小距离

给你一个整数数组 nums （下标 从 0 开始 计数）以及两个整数 target 和 start ，请你找出一个下标 i ，满足 nums[i] == target 且 abs(i - start) 最小化 。注意：abs(x) 表示 x 的绝对值。

返回 abs(i - start) 。

题目数据保证 target 存在于 nums 中。

示例 1：

    输入：nums = [1,2,3,4,5], target = 5, start = 3
    输出：1
    解释：nums[4] = 5 是唯一一个等于 target 的值，所以答案是 abs(4 - 3) = 1 。

示例 2：

    输入：nums = [1], target = 1, start = 0
    输出：0
    解释：nums[0] = 1 是唯一一个等于 target 的值，所以答案是 abs(0 - 0) = 0 。

示例 3：

    输入：nums = [1,1,1,1,1,1,1,1,1,1], target = 1, start = 0
    输出：0
    解释：nums 中的每个值都是 1 ，但 nums[0] 使 abs(i - start) 的结果得以最小化，所以答案是 abs(0 - 0) = 0 。

提示：

- 1 <= nums.length <= 1000
- 1 <= nums[i] <= 10^4
- 0 <= start < nums.length
- target 存在于 nums 中
