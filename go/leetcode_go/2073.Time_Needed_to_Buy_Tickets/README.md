### 2073. Time Needed to Buy Tickets
There are n people in a line queuing to buy tickets, where the 0th person is at the front of the line and the (n - 1)th person is at the back of the line.

You are given a 0-indexed integer array tickets of length n where the number of tickets that the i<sup>th</sup> person would like to buy is tickets[i].

Each person takes exactly 1 second to buy a ticket. A person can only buy 1 ticket at a time and has to go back to the end of the line (which happens instantaneously) in order to buy more tickets. If a person does not have any tickets left to buy, the person will leave the line.

Return the time taken for the person at position k (0-indexed) to finish buying tickets.



Example 1:

	Input: tickets = [2,3,2], k = 2
	Output: 6
	Explanation:
	- In the first pass, everyone in the line buys a ticket and the line becomes [1, 2, 1].
	- In the second pass, everyone in the line buys a ticket and the line becomes [0, 1, 0].
	The person at position 2 has successfully bought 2 tickets and it took 3 + 3 = 6 seconds.

Example 2:

	Input: tickets = [5,1,1,1], k = 0
	Output: 8
	Explanation:
	- In the first pass, everyone in the line buys a ticket and the line becomes [4, 0, 0, 0].
	- In the next 4 passes, only the person in position 0 is buying tickets.
	The person at position 0 has successfully bought 5 tickets and it took 4 + 1 + 1 + 1 + 1 = 8 seconds.



Constraints:

* n == tickets.length
* 1 <= n <= 100
* 1 <= tickets[i] <= 100
* 0 <= k < n

----

### 2073. 买票需要的时间
有 n 个人前来排队买票，其中第 0 人站在队伍 最前方 ，第 (n - 1) 人站在队伍 最后方 。

给你一个下标从 0 开始的整数数组 tickets ，数组长度为 n ，其中第 i 人想要购买的票数为 tickets[i] 。

每个人买票都需要用掉 恰好 1 秒 。一个人 一次只能买一张票 ，如果需要购买更多票，他必须走到  队尾 重新排队（瞬间 发生，不计时间）。如果一个人没有剩下需要买的票，那他将会 离开 队伍。

返回位于位置 k（下标从 0 开始）的人完成买票需要的时间（以秒为单位）。



示例 1：

	输入：tickets = [2,3,2], k = 2
	输出：6
	解释：
	- 第一轮，队伍中的每个人都买到一张票，队伍变为 [1, 2, 1] 。
	- 第二轮，队伍中的每个都又都买到一张票，队伍变为 [0, 1, 0] 。
	位置 2 的人成功买到 2 张票，用掉 3 + 3 = 6 秒。

示例 2：

	输入：tickets = [5,1,1,1], k = 0
	输出：8
	解释：
	- 第一轮，队伍中的每个人都买到一张票，队伍变为 [4, 0, 0, 0] 。
	- 接下来的 4 轮，只有位置 0 的人在买票。
	位置 0 的人成功买到 5 张票，用掉 4 + 1 + 1 + 1 + 1 = 8 秒。



提示：

* n == tickets.length
* 1 <= n <= 100
* 1 <= tickets[i] <= 100
* 0 <= k < n

