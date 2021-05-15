### 1480. Running Sum of 1d Array
Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.



Example 1:

	Input: nums = [1,2,3,4]
	Output: [1,3,6,10]
	Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

Example 2:

	Input: nums = [1,1,1,1,1]
	Output: [1,2,3,4,5]
	Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

Example 3:

	Input: nums = [3,1,2,10,1]
	Output: [3,4,6,16,17]



Constraints:

* 1 <= nums.length <= 1000
* -10^6 <= nums[i] <= 10^6

----

### 1480. 一维数组的动态和
给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。

请返回 nums 的动态和。



示例 1：

	输入：nums = [1,2,3,4]
	输出：[1,3,6,10]
	解释：动态和计算过程为 [1, 1+2, 1+2+3, 1+2+3+4] 。

示例 2：

	输入：nums = [1,1,1,1,1]
	输出：[1,2,3,4,5]
	解释：动态和计算过程为 [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1] 。

示例 3：

	输入：nums = [3,1,2,10,1]
	输出：[3,4,6,16,17]



提示：

* 1 <= nums.length <= 1000
* -10^6 <= nums[i] <= 10^6
