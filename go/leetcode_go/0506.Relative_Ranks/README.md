### 506. Relative Ranks
You are given an integer array score of size n, where score[i] is the score of the ith athlete in a competition. All the scores are guaranteed to be unique.

The athletes are placed based on their scores, where the 1st place athlete has the highest score, the 2nd place athlete has the 2nd highest score, and so on. The placement of each athlete determines their rank:

* The 1st place athlete's rank is "Gold Medal".
* The 2nd place athlete's rank is "Silver Medal".
* The 3rd place athlete's rank is "Bronze Medal".
* For the 4th place to the nth place athlete, their rank is their placement number (i.e., the xth place athlete's rank is "x").

Return an array answer of size n where answer[i] is the rank of the ith athlete.



Example 1:

	Input: score = [5,4,3,2,1]
	Output: ["Gold Medal","Silver Medal","Bronze Medal","4","5"]
	Explanation: The placements are [1st, 2nd, 3rd, 4th, 5th].

Example 2:

	Input: score = [10,3,8,9,4]
	Output: ["Gold Medal","5","Bronze Medal","Silver Medal","4"]
	Explanation: The placements are [1st, 5th, 3rd, 2nd, 4th].



Constraints:

* n == score.length
* 1 <= n <= 10^4
* 0 <= score[i] <= 10^6
* All the values in score are unique.

----

### 506. 相对名次
给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold Medal", "Silver Medal", "Bronze Medal"）。

(注：分数越高的选手，排名越靠前。)

示例 1:

	输入: [5, 4, 3, 2, 1]
	输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
	解释: 前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal" and "Bronze Medal").
	余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。

提示:

* N 是一个正整数并且不会超过 10000。
* 所有运动员的成绩都不相同。

