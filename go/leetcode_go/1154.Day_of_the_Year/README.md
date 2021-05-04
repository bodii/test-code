### 1154. Day of the Year

Given a string date representing a Gregorian calendar date formatted as YYYY-MM-DD, return the day number of the year.

Example 1:

    Input: date = "2019-01-09"
    Output: 9
    Explanation: Given date is the 9th day of the year in 2019.

Example 2:

    Input: date = "2019-02-10"
    Output: 41

Example 3:

    Input: date = "2003-03-01"
    Output: 60

Example 4:

    Input: date = "2004-03-01"
    Output: 61

Constraints:

- date.length == 10
- date[4] == date[7] == '-', and all other date[i]'s are digits
- date represents a calendar date between Jan 1st, 1900 and Dec 31, 2019.

---

### 1154. 一年中的第几天

给你一个按 YYYY-MM-DD 格式表示日期的字符串 date，请你计算并返回该日期是当年的第几天。

通常情况下，我们认为 1 月 1 日是每年的第 1 天，1 月 2 日是每年的第 2 天，依此类推。每个月的天数与现行公元纪年法（格里高利历）一致。

示例 1：

    输入：date = "2019-01-09"
    输出：9

示例 2：

    输入：date = "2019-02-10"
    输出：41

示例 3：

    输入：date = "2003-03-01"
    输出：60

示例 4：

    输入：date = "2004-03-01"
    输出：61

提示：

- date.length == 10
- date[4] == date[7] == '-'，其他的 date[i] 都是数字。
- date 表示的范围从 1900 年 1 月 1 日至 2019 年 12 月 31 日。
