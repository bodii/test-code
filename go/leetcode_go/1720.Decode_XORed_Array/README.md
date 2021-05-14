### 1720. Decode XORed Array
There is a hidden integer array arr that consists of n non-negative integers.

It was encoded into another integer array encoded of length n - 1, such that encoded[i] = arr[i] XOR arr[i + 1]. For example, if arr = [1,0,2,1], then encoded = [1,2,3].

You are given the encoded array. You are also given an integer first, that is the first element of arr, i.e. arr[0].

Return the original array arr. It can be proved that the answer exists and is unique.



Example 1:

	Input: encoded = [1,2,3], first = 1
	Output: [1,0,2,1]
	Explanation: If arr = [1,0,2,1], then first = 1 and encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]

Example 2:

	Input: encoded = [6,2,7,3], first = 4
	Output: [4,2,0,7,4]



Constraints:

* 2 <= n <= 104
* encoded.length == n - 1
* 0 <= encoded[i] <= 105
* 0 <= first <= 105

----

### 1720. 解码异或后的数组
未知 整数数组 arr 由 n 个非负整数组成。

经编码后变为长度为 n - 1 的另一个整数数组 encoded ，其中 encoded[i] = arr[i] XOR arr[i + 1] 。例如，arr = [1,0,2,1] 经编码后得到 encoded = [1,2,3] 。

给你编码后的数组 encoded 和原数组 arr 的第一个元素 first（arr[0]）。

请解码返回原数组 arr 。可以证明答案存在并且是唯一的。



示例 1：

	输入：encoded = [1,2,3], first = 1
	输出：[1,0,2,1]
	解释：若 arr = [1,0,2,1] ，那么 first = 1 且 encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]

示例 2：

	输入：encoded = [6,2,7,3], first = 4
	输出：[4,2,0,7,4]



提示：

* 2 <= n <= 104
* encoded.length == n - 1
* 0 <= encoded[i] <= 105
* 0 <= first <= 105

