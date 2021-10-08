### 面试题 10.05. Sparse Array Search LCCI
Given a sorted array of strings that is interspersed with empty strings, write a method to find the location of a given string.

Example1:

	Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
	Output: -1
	Explanation: Return -1 if s is not in words.

Example2:

	Input: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
	Output: 4

Note:

* 1 <= words.length <= 1000000

----

### 面试题 10.05. 稀疏数组搜索
稀疏数组搜索。有个排好序的字符串数组，其中散布着一些空字符串，编写一种方法，找出给定字符串的位置。

示例1:

	输入: words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ta"
	输出：-1
	说明: 不存在返回-1。

示例2:

	输入：words = ["at", "", "", "", "ball", "", "", "car", "", "","dad", "", ""], s = "ball"
	输出：4

提示:

* words的长度在[1, 1000000]之间

