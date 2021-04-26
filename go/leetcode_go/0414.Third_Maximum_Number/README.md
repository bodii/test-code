### 414. Third Maximum Number

Given integer array nums, return the third maximum number in this array. If the third maximum does not exist, return the maximum number.

 
Example 1:

    Input: nums = [3,2,1]
    Output: 1
    Explanation: The third maximum is 1.

Example 2:

    Input: nums = [1,2]
    Output: 2
    Explanation: The third maximum does not exist, so the maximum (2) is returned instead.

Example 3:

    Input: nums = [2,2,3,1]
    Output: 1
    Explanation: Note that the third maximum here means the third maximum distinct number.
    Both numbers with value 2 are both considered as second maximum.

 

Constraints:

* 1 <= nums.length <= 104
* -231 <= nums[i] <= 231 - 1

---------------

### 414. 第三大的数

给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。

 

示例 1：

    输入：[3, 2, 1]
    输出：1
    解释：第三大的数是 1 。

示例 2：

    输入：[1, 2]
    输出：2
    解释：第三大的数不存在, 所以返回最大的数 2 。

示例 3：

    输入：[2, 2, 3, 1]
    输出：1
    解释：注意，要求返回第三大的数，是指在所有不同数字中排第三大的数。
    此例中存在两个值为 2 的数，它们都排第二。在所有不同数字中排第三大的数为 1 。

 

提示：

* 1 <= nums.length <= 104
* -231 <= nums[i] <= 231 - 1