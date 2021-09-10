### 1331. Rank Transform of an Array
Given an array of integers arr, replace each element with its rank.

The rank represents how large the element is. The rank has the following rules:

* Rank is an integer starting from 1.
* The larger the element, the larger the rank. If two elements are equal, their rank must be the same.
* Rank should be as small as possible.



Example 1:

	Input: arr = [40,10,20,30]
	Output: [4,1,2,3]
	Explanation: 40 is the largest element. 10 is the smallest. 20 is the second smallest. 30 is the third smallest.

Example 2:

	Input: arr = [100,100,100]
	Output: [1,1,1]
	Explanation: Same elements share the same rank.

Example 3:

	Input: arr = [37,12,28,9,100,56,80,5,12]
	Output: [5,3,4,2,8,6,7,1,3]



Constraints:

* 0 <= arr.length <= 10^5
* -109 <= arr[i] <= 10^9

----

### 1331. 数组序号转换
给你一个整数数组 arr ，请你将数组中的每个元素替换为它们排序后的序号。

序号代表了一个元素有多大。序号编号的规则如下：

* 序号从 1 开始编号。
* 一个元素越大，那么序号越大。如果两个元素相等，那么它们的序号相同。
* 每个数字的序号都应该尽可能地小。



示例 1：

	输入：arr = [40,10,20,30]
	输出：[4,1,2,3]
	解释：40 是最大的元素。 10 是最小的元素。 20 是第二小的数字。 30 是第三小的数字。

示例 2：

	输入：arr = [100,100,100]
	输出：[1,1,1]
	解释：所有元素有相同的序号。

示例 3：

	输入：arr = [37,12,28,9,100,56,80,5,12]
	输出：[5,3,4,2,8,6,7,1,3]



提示：

* 0 <= arr.length <= 10^5
* -109 <= arr[i] <= 10^9

