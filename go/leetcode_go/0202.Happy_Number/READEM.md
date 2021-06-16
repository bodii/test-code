### 202. Happy Number
Write an algorithm to determine if a number n is happy.

A happy number is a number defined by the following process:

* Starting with any positive integer, replace the number by the sum of the squares of its digits.
* Repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1.
* Those numbers for which this process ends in 1 are happy.

Return true if n is a happy number, and false if not.



Example 1:

	Input: n = 19
	Output: true
	Explanation:
	12 + 92 = 82
	82 + 22 = 68
	62 + 82 = 100
	12 + 02 + 02 = 1

Example 2:

	Input: n = 2
	Output: false



Constraints:

* 1 <= n <= 2<sup>31</sup> - 1

----

### 202. 快乐数
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」定义为：

* 对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
* 然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
* 如果 可以变为  1，那么这个数就是快乐数。

如果 n 是快乐数就返回 true ；不是，则返回 false 。

 

示例 1：

	输入：19
	输出：true
	解释：
	12 + 92 = 82
	82 + 22 = 68
	62 + 82 = 100
	12 + 02 + 02 = 1

示例 2：

	输入：n = 2
	输出：false

 

提示：

* 1 <= n <= 2<sup>31</sup> - 1

