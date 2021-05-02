### 1089. Duplicate Zeros

Given a fixed length array arr of integers, duplicate each occurrence of zero, shifting the remaining elements to the right.

Note that elements beyond the length of the original array are not written.

Do the above modifications to the input array in place, do not return anything from your function.



 Example 1:

	 Input: [1,0,2,3,0,4,5,0]
	 Output: null
	 Explanation: After calling your function, the input array is modified to: [1,0,0,2,3,0,0,4]

 Example 2:

	 Input: [1,2,3]
	 Output: null
	 Explanation: After calling your function, the input array is modified to: [1,2,3]



  Note:

	1. 1 <= arr.length <= 10000
	2. 0 <= arr[i] <= 9

---------

### 1089. 复写零

给你一个长度固定的整数数组 arr，请你将该数组中出现的每个零都复写一遍，并将其余的元素向右平移。

注意：请不要在超过该数组长度的位置写入元素。

要求：请对输入的数组 就地 进行上述修改，不要从函数返回任何东西。



 示例 1：

	 输入：[1,0,2,3,0,4,5,0]
	 输出：null
	 解释：调用函数后，输入的数组将被修改为：[1,0,0,2,3,0,0,4]

 示例 2：

	 输入：[1,2,3]
	 输出：null
	 解释：调用函数后，输入的数组将被修改为：[1,2,3]



  提示：

    1.  1 <= arr.length <= 10000
	2.  0 <= arr[i] <= 9

