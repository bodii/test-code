### 728. Self Dividing Numbers
A self-dividing number is a number that is divisible by every digit it contains.

* For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0, and 128 % 8 == 0.

A self-dividing number is not allowed to contain the digit zero.

Given two integers left and right, return a list of all the self-dividing numbers in the range [left, right].



Example 1:

	Input: left = 1, right = 22
	Output: [1,2,3,4,5,6,7,8,9,11,12,15,22]

Example 2:

	Input: left = 47, right = 85
	Output: [48,55,66,77]



Constraints:

* 1 <= left <= right <= 104

----

### 728. 自除数
自除数 是指可以被它包含的每一位数除尽的数。

例如，128 是一个自除数，因为 128 % 1 == 0，128 % 2 == 0，128 % 8 == 0。

还有，自除数不允许包含 0 。

给定上边界和下边界数字，输出一个列表，列表的元素是边界（含边界）内所有的自除数。

示例 1：

输入：
	上边界left = 1, 下边界right = 22
	输出： [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]

注意：

* 每个输入参数的边界满足 1 <= left <= right <= 10000。

