### 908. Smallest Range I
Given an array A of integers, for each integer A[i] we may choose any x with -K <= x <= K, and add x to A[i].

After this process, we have some array B.

Return the smallest possible difference between the maximum value of B and the minimum value of B.



Example 1:

	Input: A = [1], K = 0
	Output: 0
	Explanation: B = [1]

Example 2:

	Input: A = [0,10], K = 2
	Output: 6
	Explanation: B = [2,8]

Example 3:

	Input: A = [1,3,6], K = 3
	Output: 0
	Explanation: B = [3,3,3] or B = [4,4,4]



Note:

* 1 <= A.length <= 10000
* 0 <= A[i] <= 10000
* 0 <= K <= 10000

----

### 908. 最小差值 I
给你一个整数数组 A，请你给数组中的每个元素 A[i] 都加上一个任意数字 x （-K <= x <= K），从而得到一个新数组 B 。

返回数组 B 的最大值和最小值之间可能存在的最小差值。



示例 1：

	输入：A = [1], K = 0
	输出：0
	解释：B = [1]

示例 2：

	输入：A = [0,10], K = 2
	输出：6
	解释：B = [2,8]

示例 3：

	输入：A = [1,3,6], K = 3
	输出：0
	解释：B = [3,3,3] 或 B = [4,4,4]



提示：

* 1 <= A.length <= 10000
* 0 <= A[i] <= 10000
* 0 <= K <= 10000

