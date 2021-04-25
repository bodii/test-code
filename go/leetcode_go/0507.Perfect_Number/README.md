### 507. Perfect Number

A perfect number is a positive integer that is equal to the sum of its <b>positive divisors</b>, excluding the number itself. A divisor of an integer `x` is an integer that can divide `x` evenly.

Given an integer`n`, return `true` if `n` is a perfect number, otherwise return `false`.

 

Example 1:

    Input: num = 28
    Output: true
    Explanation: 28 = 1 + 2 + 4 + 7 + 14
    1, 2, 4, 7, and 14 are all divisors of 28.


Example 2:

    Input: num = 6
    Output: true

Example 3:

    Input: num = 496
    Output: true

Example 4:

    Input: num = 8128
    Output: true

Example 5:

    Input: num = 2
    Output: false

 

Constraints:
* 1 <= num <= 108

----------------

### 507. 完美数
对于一个<b>正整数</b>，如果它和除了它自身以外的所有 <b>正因子</b> 之和相等，我们称它为 「完美数」。

给定一个 <b>整数</b> `n`， 如果是完美数，返回 `true`，否则返回 `false`

 

示例 1：

输入：28
输出：True
解释：28 = 1 + 2 + 4 + 7 + 14
1, 2, 4, 7, 和 14 是 28 的所有正因子。

示例 2：

    输入：num = 6
    输出：true

示例 3：

    输入：num = 496
    输出：true

示例 4：

    输入：num = 8128
    输出：true

示例 5：

    输入：num = 2
    输出：false

 

提示：

* 1 <= num <= 108