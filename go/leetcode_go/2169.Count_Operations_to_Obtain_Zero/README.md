### 2169. Count Operations to Obtain Zero
You are given two non-negative integers num1 and num2.

In one operation, if num1 >= num2, you must subtract num2 from num1, otherwise subtract num1 from num2.

* For example, if num1 = 5 and num2 = 4, subtract num2 from num1, thus obtaining num1 = 1 and num2 = 4. However, if num1 = 4 and num2 = 5, after one operation, num1 = 4 and num2 = 1.

Return the number of operations required to make either num1 = 0 or num2 = 0.



Example 1:

	Input: num1 = 2, num2 = 3
	Output: 3
	Explanation:
	- Operation 1: num1 = 2, num2 = 3. Since num1 < num2, we subtract num1 from num2 and get num1 = 2, num2 = 3 - 2 = 1.
	- Operation 2: num1 = 2, num2 = 1. Since num1 > num2, we subtract num2 from num1.
	- Operation 3: num1 = 1, num2 = 1. Since num1 == num2, we subtract num2 from num1.
	Now num1 = 0 and num2 = 1. Since num1 == 0, we do not need to perform any further operations.
	So the total number of operations required is 3.

Example 2:

	Input: num1 = 10, num2 = 10
	Output: 1
	Explanation:
	- Operation 1: num1 = 10, num2 = 10. Since num1 == num2, we subtract num2 from num1 and get num1 = 10 - 10 = 0.
	Now num1 = 0 and num2 = 10. Since num1 == 0, we are done.
	So the total number of operations required is 1.



Constraints:

*  0 <= num1, num2 <= 10^5

----

### 2169. 得到 0 的操作数
给你两个 非负 整数 num1 和 num2 。

每一步 操作 中，如果 num1 >= num2 ，你必须用 num1 减 num2 ；否则，你必须用 num2 减 num1 。

* 例如，num1 = 5 且 num2 = 4 ，应该用 num1 减 num2 ，因此，得到 num1 = 1 和 num2 = 4 。然而，如果 num1 = 4且 num2 = 5 ，一步操作后，得到 num1 = 4 和 num2 = 1 。

返回使 num1 = 0 或 num2 = 0 的 操作数 。



示例 1：

	输入：num1 = 2, num2 = 3
	输出：3
	解释：
	- 操作 1 ：num1 = 2 ，num2 = 3 。由于 num1 < num2 ，num2 减 num1 得到 num1 = 2 ，num2 = 3 - 2 = 1 。
	- 操作 2 ：num1 = 2 ，num2 = 1 。由于 num1 > num2 ，num1 减 num2 。
	- 操作 3 ：num1 = 1 ，num2 = 1 。由于 num1 == num2 ，num1 减 num2 。
	此时 num1 = 0 ，num2 = 1 。由于 num1 == 0 ，不需要再执行任何操作。
	所以总操作数是 3 。

示例 2：

	输入：num1 = 10, num2 = 10
	输出：1
	解释：
	- 操作 1 ：num1 = 10 ，num2 = 10 。由于 num1 == num2 ，num1 减 num2 得到 num1 = 10 - 10 = 0 。
	此时 num1 = 0 ，num2 = 10 。由于 num1 == 0 ，不需要再执行任何操作。
	所以总操作数是 1 。



提示：

* 0 <= num1, num2 <= 10^5

