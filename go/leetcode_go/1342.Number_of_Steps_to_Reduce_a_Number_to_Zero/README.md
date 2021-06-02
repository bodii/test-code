### 1342. Number of Steps to Reduce a Number to Zero
Given a non-negative integer num, return the number of steps to reduce it to zero. If the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.



Example 1:

	Input: num = 14
	Output: 6
	Explanation:
	Step 1) 14 is even; divide by 2 and obtain 7.
	Step 2) 7 is odd; subtract 1 and obtain 6.
	Step 3) 6 is even; divide by 2 and obtain 3.
	Step 4) 3 is odd; subtract 1 and obtain 2.
	Step 5) 2 is even; divide by 2 and obtain 1.
	Step 6) 1 is odd; subtract 1 and obtain 0.

Example 2:

	Input: num = 8
	Output: 4
	Explanation:
	Step 1) 8 is even; divide by 2 and obtain 4.
	Step 2) 4 is even; divide by 2 and obtain 2.
	Step 3) 2 is even; divide by 2 and obtain 1.
	Step 4) 1 is odd; subtract 1 and obtain 0.

Example 3:

	Input: num = 123
	Output: 12



Constraints:

* 0 <= num <= 10^6

----

### 1342. 将数字变成 0 的操作次数
给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。



示例 1：

	输入：num = 14
	输出：6
	解释：
	步骤 1) 14 是偶数，除以 2 得到 7 。
	步骤 2） 7 是奇数，减 1 得到 6 。
	步骤 3） 6 是偶数，除以 2 得到 3 。
	步骤 4） 3 是奇数，减 1 得到 2 。
	步骤 5） 2 是偶数，除以 2 得到 1 。
	步骤 6） 1 是奇数，减 1 得到 0 。

示例 2：

	输入：num = 8
	输出：4
	解释：
	步骤 1） 8 是偶数，除以 2 得到 4 。
	步骤 2） 4 是偶数，除以 2 得到 2 。
	步骤 3） 2 是偶数，除以 2 得到 1 。
	步骤 4） 1 是奇数，减 1 得到 0 。

示例 3：

	输入：num = 123
	输出：12



提示：

* 0 <= num <= 10^6

