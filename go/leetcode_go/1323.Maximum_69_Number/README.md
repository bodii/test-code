### 1323. Maximum 69 Number

Given a positive integer num consisting only of digits 6 and 9.

Return the maximum number you can get by changing at most one digit (6 becomes 9, and 9 becomes 6).

 

Example 1:

    Input: num = 9669
    Output: 9969
    Explanation: 
    Changing the first digit results in 6669.
    Changing the second digit results in 9969.
    Changing the third digit results in 9699.
    Changing the fourth digit results in 9666. 
    The maximum number is 9969.

Example 2:

    Input: num = 9996
    Output: 9999
    Explanation: Changing the last digit 6 to 9 results in the maximum number.

Example 3:

    Input: num = 9999
    Output: 9999
    Explanation: It is better not to apply any change.

 

Constraints:

* 1 <= num <= 10^4
* num's digits are 6 or 9.

-------------------

### 1323. 6 和 9 组成的最大数字

给你一个仅由数字 6 和 9 组成的正整数 num。

你最多只能翻转一位数字，将 6 变成 9，或者把 9 变成 6 。

请返回你可以得到的最大数字。

 

示例 1：

    输入：num = 9669
    输出：9969
    解释：
    改变第一位数字可以得到 6669 。
    改变第二位数字可以得到 9969 。
    改变第三位数字可以得到 9699 。
    改变第四位数字可以得到 9666 。
    其中最大的数字是 9969 。

示例 2：

    输入：num = 9996
    输出：9999
    解释：将最后一位从 6 变到 9，其结果 9999 是最大的数。

示例 3：

    输入：num = 9999
    输出：9999
    解释：无需改变就已经是最大的数字了。

 

提示：

* 1 <= num <= 10^4
* num 每一位上的数字都是 6 或者 9 。