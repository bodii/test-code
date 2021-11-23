### 2022. Convert 1D Array Into 2D Array
You are given a 0-indexed 1-dimensional (1D) integer array original, and two integers, m and n. You are tasked with creating a 2-dimensional (2D) array with m rows and n columns using all the elements from original.

The elements from indices 0 to n - 1 (inclusive) of original should form the first row of the constructed 2D array, the elements from indices n to 2 * n - 1 (inclusive) should form the second row of the constructed 2D array, and so on.

Return an m x n 2D array constructed according to the above procedure, or an empty 2D array if it is impossible.



Example 1:
	Input: original = [1,2,3,4], m = 2, n = 2
	Output: [[1,2],[3,4]]
	Explanation:
	The constructed 2D array should contain 2 rows and 2 columns.
	The first group of n=2 elements in original, [1,2], becomes the first row in the constructed 2D array.
	The second group of n=2 elements in original, [3,4], becomes the second row in the constructed 2D array.

Example 2:

	Input: original = [1,2,3], m = 1, n = 3
	Output: [[1,2,3]]
	Explanation:
	The constructed 2D array should contain 1 row and 3 columns.
	Put all three elements in original into the first row of the constructed 2D array.

Example 3:

	Input: original = [1,2], m = 1, n = 1
	Output: []
	Explanation:
	There are 2 elements in original.
	It is impossible to fit 2 elements in a 1x1 2D array, so return an empty 2D array.

Example 4:

	Input: original = [3], m = 1, n = 2
	Output: []
	Explanation:
	There is 1 element in original.
	It is impossible to make 1 element fill all the spots in a 1x2 2D array, so return an empty 2D array.

 

Constraints:
* 1 <= original.length <= 5 * 10^4
* 1 <= original[i] <= 10^5
* 1 <= m, n <= 4 * 10^4

----
### 2022. 将一维数组转变成二维数组
给你一个下标从 0 开始的一维整数数组 original 和两个整数 m 和  n 。你需要使用 original 中 所有 元素创建一个 m 行 n 列的二维数组。

original 中下标从 0 到 n - 1 （都 包含 ）的元素构成二维数组的第一行，下标从 n 到 2 * n - 1 （都 包含 ）的元素构成二维数组的第二行，依此类推。

请你根据上述过程返回一个 m x n 的二维数组。如果无法构成这样的二维数组，请你返回一个空的二维数组。



示例 1：
	输入：original = [1,2,3,4], m = 2, n = 2
	输出：[[1,2],[3,4]]
	解释：
	构造出的二维数组应该包含 2 行 2 列。
	original 中第一个 n=2 的部分为 [1,2] ，构成二维数组的第一行。
	original 中第二个 n=2 的部分为 [3,4] ，构成二维数组的第二行。

示例 2：

	输入：original = [1,2,3], m = 1, n = 3
	输出：[[1,2,3]]
	解释：
	构造出的二维数组应该包含 1 行 3 列。
	将 original 中所有三个元素放入第一行中，构成要求的二维数组。

示例 3：

	输入：original = [1,2], m = 1, n = 1
	输出：[]
	解释：
	original 中有 2 个元素。
	无法将 2 个元素放入到一个 1x1 的二维数组中，所以返回一个空的二维数组。

示例 4：

	输入：original = [3], m = 1, n = 2
	输出：[]
	解释：
	original 中只有 1 个元素。
	无法将 1 个元素放满一个 1x2 的二维数组，所以返回一个空的二维数组。

 

提示：

* 1 <= original.length <= 5 * 10^4
* 1 <= original[i] <= 10^5
* 1 <= m, n <= 4 * 10^4

