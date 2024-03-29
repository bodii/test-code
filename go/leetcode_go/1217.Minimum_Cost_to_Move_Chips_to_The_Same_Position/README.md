### 1217. Minimum Cost to Move Chips to The Same Position
We have n chips, where the position of the ith chip is position[i].

We need to move all the chips to the same position. In one step, we can change the position of the ith chip from position[i] to:

* position[i] + 2 or position[i] - 2 with cost = 0.
* position[i] + 1 or position[i] - 1 with cost = 1.

Return the minimum cost needed to move all the chips to the same position.



Example 1:

	Input: position = [1,2,3]
	Output: 1
	Explanation: First step: Move the chip at position 3 to position 1 with cost = 0.
	Second step: Move the chip at position 2 to position 1 with cost = 1.
	Total cost is 1.

Example 2:

	Input: position = [2,2,2,3,3]
	Output: 2
	Explanation: We can move the two chips at position  3 to position 2. Each move has cost = 1. The total cost = 2.

Example 3:

	Input: position = [1,1000000000]
	Output: 1



Constraints:

* 1 <= position.length <= 100
* 1 <= position[i] <= 10^9

---- 

### 1217. 玩筹码
轴上放置了一些筹码，每个筹码的位置存在数组 chips 当中。

你可以对 任何筹码 执行下面两种操作之一（不限操作次数，0 次也可以）：

* 将第 i 个筹码向左或者右移动 2 个单位，代价为 0。
* 将第 i 个筹码向左或者右移动 1 个单位，代价为 1。

最开始的时候，同一位置上也可能放着两个或者更多的筹码。

返回将所有筹码移动到同一位置（任意位置）上所需要的最小代价。



示例 1：

	输入：chips = [1,2,3]
	输出：1
	解释：第二个筹码移动到位置三的代价是 1，第一个筹码移动到位置三的代价是 0，总代价为 1。

示例 2：

	输入：chips = [2,2,2,3,3]
	输出：2
	解释：第四和第五个筹码移动到位置二的代价都是 1，所以最小总代价为 2。



提示：

* 1 <= chips.length <= 100
* 1 <= chips[i] <= 10^9

