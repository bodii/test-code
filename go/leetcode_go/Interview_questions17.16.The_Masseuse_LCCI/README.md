### 面试题 17.16. The Masseuse LCCI
A popular masseuse receives a sequence of back-to-back appointment requests and is debating which ones to accept. She needs a break between appointments and therefore she cannot accept any adjacent requests. Given a sequence of back-to-back appoint­ ment requests, find the optimal (highest total booked minutes) set the masseuse can honor. Return the number of minutes.

Note: This problem is slightly different from the original one in the book.



Example 1:

	Input:  [1,2,3,1]
	Output:  4
	Explanation:  Accept request 1 and 3, total minutes = 1 + 3 = 4

Example 2:

	Input:  [2,7,9,3,1]
	Output:  12
	Explanation:  Accept request 1, 3 and 5, total minutes = 2 + 9 + 1 = 12

Example 3:

	Input:  [2,1,4,5,3,1,1,3]
	Output:  12
	Explanation:  Accept request 1, 3, 5 and 8, total minutes = 2 + 4 + 3 + 3 = 12

----

### 面试题 17.16. 按摩师
一个有名的按摩师会收到源源不断的预约请求，每个预约都可以选择接或不接。在每次预约服务之间要有休息时间，因此她不能接受相邻的预约。给定一个预约请求序列，替按摩师找到最优的预约集合（总预约时间最长），返回总的分钟数。

注意：本题相对原题稍作改动



示例 1：

	输入： [1,2,3,1]
	输出： 4
	解释： 选择 1 号预约和 3 号预约，总时长 = 1 + 3 = 4。

示例 2：

	输入： [2,7,9,3,1]
	输出： 12
	解释： 选择 1 号预约、 3 号预约和 5 号预约，总时长 = 2 + 9 + 1 = 12。

示例 3：

	输入： [2,1,4,5,3,1,1,3]
	输出： 12
	解释： 选择 1 号预约、 3 号预约、 5 号预约和 8 号预约，总时长 = 2 + 4 + 3 + 3 = 12。

