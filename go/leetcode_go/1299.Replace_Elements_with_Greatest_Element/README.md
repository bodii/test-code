### 1299. Replace Elements with Greatest Element on Right Side
Given an array arr, replace every element in that array with the greatest element among the elements to its right, and replace the last element with -1.

After doing so, return the array.



Example 1:

	Input: arr = [17,18,5,4,6,1]
	Output: [18,6,6,6,1,-1]
	Explanation:
	- index 0 --> the greatest element to the right of index 0 is index 1 (18).
	- index 1 --> the greatest element to the right of index 1 is index 4 (6).
	- index 2 --> the greatest element to the right of index 2 is index 4 (6).
	- index 3 --> the greatest element to the right of index 3 is index 4 (6).
	- index 4 --> the greatest element to the right of index 4 is index 5 (1).
	- index 5 --> there are no elements to the right of index 5, so we put -1.

Example 2:

	Input: arr = [400]
	Output: [-1]
	Explanation: There are no elements to the right of index 0.



Constraints:

* 1 <= arr.length <= 10^4
* 1 <= arr[i] <= 10^5

----

### 1299. 将每个元素替换为右侧最大元素
给你一个数组 arr ，请你将每个元素用它右边最大的元素替换，如果是最后一个元素，用 -1 替换。

完成所有替换操作后，请你返回这个数组。



示例 1：

	输入：arr = [17,18,5,4,6,1]
	输出：[18,6,6,6,1,-1]
	解释：
	- 下标 0 的元素 --> 右侧最大元素是下标 1 的元素 (18)
	- 下标 1 的元素 --> 右侧最大元素是下标 4 的元素 (6)
	- 下标 2 的元素 --> 右侧最大元素是下标 4 的元素 (6)
	- 下标 3 的元素 --> 右侧最大元素是下标 4 的元素 (6)
	- 下标 4 的元素 --> 右侧最大元素是下标 5 的元素 (1)
	- 下标 5 的元素 --> 右侧没有其他元素，替换为 -1

示例 2：

	输入：arr = [400]
	输出：[-1]
	解释：下标 0 的元素右侧没有其他元素。



提示：

* 1 <= arr.length <= 10^4
* 1 <= arr[i] <= 10^5
