### 1395. Count Number of Teams
There are n soldiers standing in a line. Each soldier is assigned a unique rating value.

You have to form a team of 3 soldiers amongst them under the following rules:

* Choose 3 soldiers with index (i, j, k) with rating (rating[i], rating[j], rating[k]).
* A team is valid if: (rating[i] < rating[j] < rating[k]) or (rating[i] > rating[j] > rating[k]) where (0 <= i < j < k < n).

Return the number of teams you can form given the conditions. (soldiers can be part of multiple teams).



Example 1:

	Input: rating = [2,5,3,4,1]
	Output: 3
	Explanation: We can form three teams given the conditions. (2,3,4), (5,4,1), (5,3,1).

Example 2:

	Input: rating = [2,1,3]
	Output: 0
	Explanation: We can't form any team given the conditions.

Example 3:

	Input: rating = [1,2,3,4]
	Output: 4



Constraints:

* n == rating.length
* 3 <= n <= 1000
* 1 <= rating[i] <= 105
* All the integers in rating are unique.

----

### 1395. 统计作战单位数
n 名士兵站成一排。每个士兵都有一个 独一无二 的评分 rating 。

每 3 个士兵可以组成一个作战单位，分组规则如下：

    从队伍中选出下标分别为 i、j、k 的 3 名士兵，他们的评分分别为 rating[i]、rating[j]、rating[k]
    作战单位需满足： rating[i] < rating[j] < rating[k] 或者 rating[i] > rating[j] > rating[k] ，其中  0 <= i < j < k < n

请你返回按上述条件可以组建的作战单位数量。每个士兵都可以是多个作战单位的一部分。

 

示例 1：

	输入：rating = [2,5,3,4,1]
	输出：3
	解释：我们可以组建三个作战单位 (2,3,4)、(5,4,1)、(5,3,1) 。

示例 2：

	输入：rating = [2,1,3]
	输出：0
	解释：根据题目条件，我们无法组建作战单位。

示例 3：

	输入：rating = [1,2,3,4]
	输出：4

 

提示：

* n == rating.length
* 3 <= n <= 1000
* 1 <= rating[i] <= 10^5
* rating 中的元素都是唯一的

