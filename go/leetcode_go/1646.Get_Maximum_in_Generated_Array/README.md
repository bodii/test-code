### 1646. Get Maximum in Generated Array
You are given an integer n. An array nums of length n + 1 is generated in the following way:

* nums[0] = 0
* nums[1] = 1
* nums[2 * i] = nums[i] when 2 <= 2 * i <= n
* nums[2 * i + 1] = nums[i] + nums[i + 1] when 2 <= 2 * i + 1 <= n

Return the maximum integer in the array nums.



Example 1:

	Input: n = 7
	Output: 3
	Explanation: According to the given rules:
	  nums[0] = 0
	  nums[1] = 1
	  nums[(1 * 2) = 2] = nums[1] = 1
	  nums[(1 * 2) + 1 = 3] = nums[1] + nums[2] = 1 + 1 = 2
	  nums[(2 * 2) = 4] = nums[2] = 1
	  nums[(2 * 2) + 1 = 5] = nums[2] + nums[3] = 1 + 2 = 3
	  nums[(3 * 2) = 6] = nums[3] = 2
	  nums[(3 * 2) + 1 = 7] = nums[3] + nums[4] = 2 + 1 = 3
	Hence, nums = [0,1,1,2,1,3,2,3], and the maximum is 3.

Example 2:

	Input: n = 2
	Output: 1
	Explanation: According to the given rules, the maximum between nums[0], nums[1], and nums[2] is 1.

Example 3:

	Input: n = 3
	Output: 2
	Explanation: According to the given rules, the maximum between nums[0], nums[1], nums[2], and nums[3] is 2.



Constraints:

* 0 <= n <= 100

----

### 1646. 获取生成数组中的最大值
给你一个整数 n 。按下述规则生成一个长度为 n + 1 的数组 nums ：

* nums[0] = 0
* nums[1] = 1
* 当 2 <= 2 * i <= n 时，nums[2 * i] = nums[i]
* 当 2 <= 2 * i + 1 <= n 时，nums[2 * i + 1] = nums[i] + nums[i + 1]

返回生成数组 nums 中的 最大 值。



示例 1：

	输入：n = 7
	输出：3
	解释：根据规则：
	  nums[0] = 0
	  nums[1] = 1
	  nums[(1 * 2) = 2] = nums[1] = 1
	  nums[(1 * 2) + 1 = 3] = nums[1] + nums[2] = 1 + 1 = 2
	  nums[(2 * 2) = 4] = nums[2] = 1
	  nums[(2 * 2) + 1 = 5] = nums[2] + nums[3] = 1 + 2 = 3
	  nums[(3 * 2) = 6] = nums[3] = 2
	  nums[(3 * 2) + 1 = 7] = nums[3] + nums[4] = 2 + 1 = 3
	因此，nums = [0,1,1,2,1,3,2,3]，最大值 3

示例 2：

	输入：n = 2
	输出：1
	解释：根据规则，nums[0]、nums[1] 和 nums[2] 之中的最大值是 1

示例 3：

	输入：n = 3
	输出：2
	解释：根据规则，nums[0]、nums[1]、nums[2] 和 nums[3] 之中的最大值是 2



提示：

* 0 <= n <= 100

