### 575. Distribute Candies
Alice has n candies, where the i<sup>th</sup> candy is of type candyType[i]. Alice noticed that she started to gain weight, so she visited a doctor.

The doctor advised Alice to only eat n / 2 of the candies she has (n is always even). Alice likes her candies very much, and she wants to eat the maximum number of different types of candies while still following the doctor's advice.

Given the integer array candyType of length n, return the maximum number of different types of candies she can eat if she only eats n / 2 of them.



Example 1:

	Input: candyType = [1,1,2,2,3,3]
	Output: 3
	Explanation: Alice can only eat 6 / 2 = 3 candies. Since there are only 3 types, she can eat one of each type.

Example 2:

	Input: candyType = [1,1,2,3]
	Output: 2
	Explanation: Alice can only eat 4 / 2 = 2 candies. Whether she eats types [1,2], [1,3], or [2,3], she still can only eat 2 different types.

Example 3:

	Input: candyType = [6,6,6,6]
	Output: 1
	Explanation: Alice can only eat 4 / 2 = 2 candies. Even though she can eat 2 candies, she only has 1 type.



Constraints:

* n == candyType.length
* 2 <= n <= 10^4
* n is even.
* -10^5 <= candyType[i] <= 10^5

----

### 575. 分糖果
Alice 有 n 枚糖，其中第 i 枚糖的类型为 candyType[i] 。Alice 注意到她的体重正在增长，所以前去拜访了一位医生。

医生建议 Alice 要少摄入糖分，只吃掉她所有糖的 n / 2 即可（n 是一个偶数）。Alice 非常喜欢这些糖，她想要在遵循医生建议的情况下，尽可能吃到最多不同种类的糖。

给你一个长度为 n 的整数数组 candyType ，返回： Alice 在仅吃掉 n / 2 枚糖的情况下，可以吃到糖的最多种类数。



示例 1：

	输入：candyType = [1,1,2,2,3,3]
	输出：3
	解释：Alice 只能吃 6 / 2 = 3 枚糖，由于只有 3 种糖，她可以每种吃一枚。

示例 2：

	输入：candyType = [1,1,2,3]
	输出：2
	解释：Alice 只能吃 4 / 2 = 2 枚糖，不管她选择吃的种类是 [1,2]、[1,3] 还是 [2,3]，她只能吃到两种不同类的糖。

示例 3：

	输入：candyType = [6,6,6,6]
	输出：1
	解释：Alice 只能吃 4 / 2 = 2 枚糖，尽管她能吃 2 枚，但只能吃到 1 种糖。



提示：

* n == candyType.length
* 2 <= n <= 10^4
* n 是一个偶数
* -10^5 <= candyType[i] <= 10^5

