### 1475. Final Prices With a Special Discount in a Shop
Given the array prices where prices[i] is the price of the ith item in a shop. There is a special discount for items in the shop, if you buy the ith item, then you will receive a discount equivalent to prices[j] where j is the minimum index such that j > i and prices[j] <= prices[i], otherwise, you will not receive any discount at all.

Return an array where the ith element is the final price you will pay for the ith item of the shop considering the special discount.



Example 1:

	Input: prices = [8,4,6,2,3]
	Output: [4,2,4,2,3]
	Explanation:
	For item 0 with price[0]=8 you will receive a discount equivalent to prices[1]=4, therefore, the final price you will pay is 8 - 4 = 4.
	For item 1 with price[1]=4 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 4 - 2 = 2.
	For item 2 with price[2]=6 you will receive a discount equivalent to prices[3]=2, therefore, the final price you will pay is 6 - 2 = 4.
	For items 3 and 4 you will not receive any discount at all.

Example 2:

	Input: prices = [1,2,3,4,5]
	Output: [1,2,3,4,5]
	Explanation: In this case, for all items, you will not receive any discount at all.

Example 3:

	Input: prices = [10,1,1,6]
	Output: [9,0,1,6]



Constraints:

* 1 <= prices.length <= 500
* 1 <= prices[i] <= 10^3

----

### 1475. 商品折扣后的最终价格
给你一个数组 prices ，其中 prices[i] 是商店里第 i 件商品的价格。

商店里正在进行促销活动，如果你要买第 i 件商品，那么你可以得到与 prices[j] 相等的折扣，其中 j 是满足 j > i 且 prices[j] <= prices[i] 的 最小下标 ，如果没有满足条件的 j ，你将没有任何折扣。

请你返回一个数组，数组中第 i 个元素是折扣后你购买商品 i 最终需要支付的价格。



示例 1：

	输入：prices = [8,4,6,2,3]
	输出：[4,2,4,2,3]
	解释：
	商品 0 的价格为 price[0]=8 ，你将得到 prices[1]=4 的折扣，所以最终价格为 8 - 4 = 4 。
	商品 1 的价格为 price[1]=4 ，你将得到 prices[3]=2 的折扣，所以最终价格为 4 - 2 = 2 。
	商品 2 的价格为 price[2]=6 ，你将得到 prices[3]=2 的折扣，所以最终价格为 6 - 2 = 4 。
	商品 3 和 4 都没有折扣。

示例 2：

	输入：prices = [1,2,3,4,5]
	输出：[1,2,3,4,5]
	解释：在这个例子中，所有商品都没有折扣。

示例 3：

	输入：prices = [10,1,1,6]
	输出：[9,0,1,6]



提示：

* 1 <= prices.length <= 500
* 1 <= prices[i] <= 10^3
