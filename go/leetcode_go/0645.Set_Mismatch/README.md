### 645. Set Mismatch
You have a set of integers s, which originally contains all the numbers from 1 to n. Unfortunately, due to some error, one of the numbers in s got duplicated to another number in the set, which results in repetition of one number and loss of another number.

You are given an integer array nums representing the data status of this set after the error.

Find the number that occurs twice and the number that is missing and return them in the form of an array.



Example 1:

	Input: nums = [1,2,2,4]
	Output: [2,3]

Example 2:

	Input: nums = [1,1]
	Output: [1,2]



Constraints:

* 2 <= nums.length <= 10^4
* 1 <= nums[i] <= 10^4

----

### 645. 错误的集合
集合 s 包含从 1 到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。

给定一个数组 nums 代表了集合 S 发生错误后的结果。

请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。



示例 1：

	输入：nums = [1,2,2,4]
	输出：[2,3]

示例 2：

	输入：nums = [1,1]
	输出：[1,2]



提示：

* 2 <= nums.length <= 10^4
* 1 <= nums[i] <= 10^4

