### 1539. Kth Missing Positive Number
Given an array arr of positive integers sorted in a strictly increasing order, and an integer k.

Find the kth positive integer that is missing from this array.



Example 1:

	Input: arr = [2,3,4,7,11], k = 5
	Output: 9
	Explanation: The missing positive integers are [1,5,6,8,9,10,12,13,...]. The 5th missing positive integer is 9.

Example 2:

	Input: arr = [1,2,3,4], k = 2
	Output: 6
	Explanation: The missing positive integers are [5,6,7,...]. The 2nd missing positive integer is 6.



Constraints:

    1 <= arr.length <= 1000
    1 <= arr[i] <= 1000
    1 <= k <= 1000
    arr[i] < arr[j] for 1 <= i < j <= arr.length

----

### 1539. 第 k 个缺失的正整数
给你一个 严格升序排列 的正整数数组 arr 和一个整数 k 。

请你找到这个数组里第 k 个缺失的正整数。



示例 1：

	输入：arr = [2,3,4,7,11], k = 5
	输出：9
	解释：缺失的正整数包括 [1,5,6,8,9,10,12,13,...] 。第 5 个缺失的正整数为 9 。

示例 2：

	输入：arr = [1,2,3,4], k = 2
	输出：6
	解释：缺失的正整数包括 [5,6,7,...] 。第 2 个缺失的正整数为 6 。



提示：

* 1 <= arr.length <= 1000
* 1 <= arr[i] <= 1000
* 1 <= k <= 1000
* 对于所有 1 <= i < j <= arr.length 的 i 和 j 满足 arr[i] < arr[j]

