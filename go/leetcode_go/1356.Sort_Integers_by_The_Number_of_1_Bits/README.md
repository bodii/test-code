### 1356. Sort Integers by The Number of 1 Bits
Given an integer array arr. You have to sort the integers in the array in ascending order by the number of 1's in their binary representation and in case of two or more integers have the same number of 1's you have to sort them in ascending order.

Return the sorted array.



Example 1:

	Input: arr = [0,1,2,3,4,5,6,7,8]
	Output: [0,1,2,4,8,3,5,6,7]
	Explantion: [0] is the only integer with 0 bits.
	[1,2,4,8] all have 1 bit.
	[3,5,6] have 2 bits.
	[7] has 3 bits.
	The sorted array by bits is [0,1,2,4,8,3,5,6,7]

Example 2:

	Input: arr = [1024,512,256,128,64,32,16,8,4,2,1]
	Output: [1,2,4,8,16,32,64,128,256,512,1024]
	Explantion: All integers have 1 bit in the binary representation, you should just sort them in ascending order.

Example 3:

	Input: arr = [10000,10000]
	Output: [10000,10000]

Example 4:

	Input: arr = [2,3,5,7,11,13,17,19]
	Output: [2,3,5,17,7,11,13,19]

Example 5:

	Input: arr = [10,100,1000,10000]
	Output: [10,100,10000,1000]



Constraints:

* 1 <= arr.length <= 500
* 0 <= arr[i] <= 10^4

----
### 1356. 根据数字二进制下 1 的数目排序
给你一个整数数组 arr 。请你将数组中的元素按照其二进制表示中数字 1 的数目升序排序。

如果存在多个数字二进制中 1 的数目相同，则必须将它们按照数值大小升序排列。

请你返回排序后的数组。



示例 1：

	输入：arr = [0,1,2,3,4,5,6,7,8]
	输出：[0,1,2,4,8,3,5,6,7]
	解释：[0] 是唯一一个有 0 个 1 的数。
	[1,2,4,8] 都有 1 个 1 。
	[3,5,6] 有 2 个 1 。
	[7] 有 3 个 1 。
	按照 1 的个数排序得到的结果数组为 [0,1,2,4,8,3,5,6,7]

示例 2：

	输入：arr = [1024,512,256,128,64,32,16,8,4,2,1]
	输出：[1,2,4,8,16,32,64,128,256,512,1024]
	解释：数组中所有整数二进制下都只有 1 个 1 ，所以你需要按照数值大小将它们排序。

示例 3：

	输入：arr = [10000,10000]
	输出：[10000,10000]

示例 4：

	输入：arr = [2,3,5,7,11,13,17,19]
	输出：[2,3,5,17,7,11,13,19]

示例 5：

	输入：arr = [10,100,1000,10000]
	输出：[10,100,10000,1000]



提示：

* 1 <= arr.length <= 500
* 0 <= arr[i] <= 10^4

