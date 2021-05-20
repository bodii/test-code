### 1502. Can Make Arithmetic Progression From Sequence
A sequence of numbers is called an arithmetic progression if the difference between any two consecutive elements is the same.

Given an array of numbers arr, return true if the array can be rearranged to form an arithmetic progression. Otherwise, return false.



Example 1:

	Input: arr = [3,5,1]
	Output: true
	Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with differences 2 and -2 respectively, between each consecutive elements.

Example 2:

	Input: arr = [1,2,4]
	Output: false
	Explanation: There is no way to reorder the elements to obtain an arithmetic progression.



Constraints:

* 2 <= arr.length <= 1000
* -106 <= arr[i] <= 106

----

### 1502. 判断能否形成等差数列
给你一个数字数组 arr 。

如果一个数列中，任意相邻两项的差总等于同一个常数，那么这个数列就称为 等差数列 。

如果可以重新排列数组形成等差数列，请返回 true ；否则，返回 false 。



示例 1：

	输入：arr = [3,5,1]
	输出：true
	解释：对数组重新排序得到 [1,3,5] 或者 [5,3,1] ，任意相邻两项的差分别为 2 或 -2 ，可以形成等差数列。

示例 2：

	输入：arr = [1,2,4]
	输出：false
	解释：无法通过重新排序得到等差数列。



提示：

* 2 <= arr.length <= 1000
* -10^6 <= arr[i] <= 10^6

