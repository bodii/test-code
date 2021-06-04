### 922. Sort Array By Parity II
Given an array of integers nums, half of the integers in nums are odd, and the other half are even.

Sort the array so that whenever nums[i] is odd, i is odd, and whenever nums[i] is even, i is even.

Return any answer array that satisfies this condition.



Example 1:

	Input: nums = [4,2,5,7]
	Output: [4,5,2,7]
	Explanation: [4,7,2,5], [2,5,4,7], [2,7,4,5] would also have been accepted.

Example 2:

	Input: nums = [2,3]
	Output: [2,3]



Constraints:

* 2 <= nums.length <= 2 * 104
* nums.length is even.
* Half of the integers in nums are even.
* 0 <= nums[i] <= 1000

---- 

### 922. 按奇偶排序数组 II
给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。



示例：

	输入：[4,2,5,7]
	输出：[4,5,2,7]
	解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。



提示：

* 2 <= A.length <= 20000
* A.length % 2 == 0
* 0 <= A[i] <= 1000
