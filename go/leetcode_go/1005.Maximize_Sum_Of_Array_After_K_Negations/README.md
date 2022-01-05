### 1005. Maximize Sum Of Array After K Negations
Given an integer array nums and an integer k, modify the array in the following way:

* choose an index i and replace nums[i] with -nums[i].

You should apply this process exactly k times. You may choose the same index i multiple times.

Return the largest possible sum of the array after modifying it in this way.



Example 1:

	Input: nums = [4,2,3], k = 1
	Output: 5
	Explanation: Choose index 1 and nums becomes [4,-2,3].

Example 2:

	Input: nums = [3,-1,0,2], k = 3
	Output: 6
	Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].

Example 3:

	Input: nums = [2,-3,-1,5,-4], k = 2
	Output: 13
	Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].



Constraints:

* 1 <= nums.length <= 10^4
* -100 <= nums[i] <= 100
* 1 <= k <= 10^4

----

### 1005. K 次取反后最大化的数组和
给你一个整数数组 nums 和一个整数 k ，按以下方法修改该数组：

* 选择某个下标 i 并将 nums[i] 替换为 -nums[i] 。

重复这个过程恰好 k 次。可以多次选择同一个下标 i 。

以这种方式修改数组后，返回数组 可能的最大和 。



示例 1：

	输入：nums = [4,2,3], k = 1
	输出：5
	解释：选择下标 1 ，nums 变为 [4,-2,3] 。

示例 2：

	输入：nums = [3,-1,0,2], k = 3
	输出：6
	解释：选择下标 (1, 2, 2) ，nums 变为 [3,1,0,2] 。

示例 3：

	输入：nums = [2,-3,-1,5,-4], k = 2
	输出：13
	解释：选择下标 (1, 4) ，nums 变为 [2,3,-1,5,4] 。



提示：

* 1 <= nums.length <= 10^4
* -100 <= nums[i] <= 100
* 1 <= k <= 10^4
