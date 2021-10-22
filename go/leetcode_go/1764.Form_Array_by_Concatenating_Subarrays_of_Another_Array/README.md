### 1764. Form Array by Concatenating Subarrays of Another Array
You are given a 2D integer array groups of length n. You are also given an integer array nums.

You are asked if you can choose n disjoint subarrays from the array nums such that the ith subarray is equal to groups[i] (0-indexed), and if i > 0, the (i-1)th subarray appears before the ith subarray in nums (i.e. the subarrays must be in the same order as groups).

Return true if you can do this task, and false otherwise.

Note that the subarrays are disjoint if and only if there is no index k such that nums[k] belongs to more than one subarray. A subarray is a contiguous sequence of elements within an array.



Example 1:

	Input: groups = [[1,-1,-1],[3,-2,0]], nums = [1,-1,0,1,-1,-1,3,-2,0]
	Output: true
	Explanation: You can choose the 0th subarray as [1,-1,0,1,-1,-1,3,-2,0] and the 1st one as [1,-1,0,1,-1,-1,3,-2,0].
	These subarrays are disjoint as they share no common nums[k] element.

Example 2:

	Input: groups = [[10,-2],[1,2,3,4]], nums = [1,2,3,4,10,-2]
	Output: false
	Explanation: Note that choosing the subarrays [1,2,3,4,10,-2] and [1,2,3,4,10,-2] is incorrect because they are not in the same order as in groups.
	[10,-2] must come before [1,2,3,4].

Example 3:

	Input: groups = [[1,2,3],[3,4]], nums = [7,7,1,2,3,4,7,7]
	Output: false
	Explanation: Note that choosing the subarrays [7,7,1,2,3,4,7,7] and [7,7,1,2,3,4,7,7] is invalid because they are not disjoint.
	They share a common elements nums[4] (0-indexed).



Constraints:

    groups.length == n
    1 <= n <= 10^3
    1 <= groups[i].length, sum(groups[i].length) <= 10^3
    1 <= nums.length <= 10^3
    -10^7 <= groups[i][j], nums[k] <= 10^7

----
### 1764. 通过连接另一个数组的子数组得到一个数组
给你一个长度为 n 的二维整数数组 groups ，同时给你一个整数数组 nums 。

你是否可以从 nums 中选出 n 个 不相交 的子数组，使得第 i 个子数组与 groups[i] （下标从 0 开始）完全相同，且如果 i > 0 ，那么第 (i-1) 个子数组在 nums 中出现的位置在第 i 个子数组前面。（也就是说，这些子数组在 nums 中出现的顺序需要与 groups 顺序相同）

如果你可以找出这样的 n 个子数组，请你返回 true ，否则返回 false 。

如果不存在下标为 k 的元素 nums[k] 属于不止一个子数组，就称这些子数组是 不相交 的。子数组指的是原数组中连续元素组成的一个序列。



示例 1：

	输入：groups = [[1,-1,-1],[3,-2,0]], nums = [1,-1,0,1,-1,-1,3,-2,0]
	输出：true
	解释：你可以分别在 nums 中选出第 0 个子数组 [1,-1,0,1,-1,-1,3,-2,0] 和第 1 个子数组 [1,-1,0,1,-1,-1,3,-2,0] 。
	这两个子数组是不相交的，因为它们没有任何共同的元素。

示例 2：

	输入：groups = [[10,-2],[1,2,3,4]], nums = [1,2,3,4,10,-2]
	输出：false
	解释：选择子数组 [1,2,3,4,10,-2] 和 [1,2,3,4,10,-2] 是不正确的，因为它们出现的顺序与 groups 中顺序不同。
	[10,-2] 必须出现在 [1,2,3,4] 之前。

示例 3：

	输入：groups = [[1,2,3],[3,4]], nums = [7,7,1,2,3,4,7,7]
	输出：false
	解释：选择子数组 [7,7,1,2,3,4,7,7] 和 [7,7,1,2,3,4,7,7] 是不正确的，因为它们不是不相交子数组。
	它们有一个共同的元素 nums[4] （下标从 0 开始）。



提示：

    groups.length == n
    1 <= n <= 10^3
    1 <= groups[i].length, sum(groups[i].length) <= 10^3
    1 <= nums.length <= 10^3
    -10^7 <= groups[i][j], nums[k] <= 10^7

