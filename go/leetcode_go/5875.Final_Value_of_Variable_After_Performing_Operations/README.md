### 5875. Final Value of Variable After Performing Operations
There is a programming language with only four operations and one variable X:

* ++X and X++ increments the value of the variable X by 1.
* --X and X-- decrements the value of the variable X by 1.

Initially, the value of X is 0.

Given an array of strings operations containing a list of operations, return the final value of X after performing all the operations.

 

Example 1:

	Input: operations = ["--X","X++","X++"]
	Output: 1
	Explanation: The operations are performed as follows:
	Initially, X = 0.
	--X: X is decremented by 1, X =  0 - 1 = -1.
	X++: X is incremented by 1, X = -1 + 1 =  0.
	X++: X is incremented by 1, X =  0 + 1 =  1.

Example 2:

	Input: operations = ["++X","++X","X++"]
	Output: 3
	Explanation: The operations are performed as follows:
	Initially, X = 0.
	++X: X is incremented by 1, X = 0 + 1 = 1.
	++X: X is incremented by 1, X = 1 + 1 = 2.
	X++: X is incremented by 1, X = 2 + 1 = 3.

Example 3:

	Input: operations = ["X++","++X","--X","X--"]
	Output: 0
	Explanation: The operations are performed as follows:
	Initially, X = 0.
	X++: X is incremented by 1, X = 0 + 1 = 1.
	++X: X is incremented by 1, X = 1 + 1 = 2.
	--X: X is decremented by 1, X = 2 - 1 = 1.
	X--: X is decremented by 1, X = 1 - 1 = 0.

 

Constraints:

* 1 <= operations.length <= 100
*  operations[i] will be either "++X", "X++", "--X", or "X--".

----
### 5875. 执行操作后的变量值
存在一种仅支持 4 种操作和 1 个变量 X 的编程语言：

* ++X 和 X++ 使变量 X 的值 加 1
* --X 和 X-- 使变量 X 的值 减 1

最初，X 的值是 0

给你一个字符串数组 operations ，这是由操作组成的一个列表，返回执行所有操作后， X 的 最终值 。

 

示例 1：

	输入：operations = ["--X","X++","X++"]
	输出：1
	解释：操作按下述步骤执行：
	最初，X = 0
	--X：X 减 1 ，X =  0 - 1 = -1
	X++：X 加 1 ，X = -1 + 1 =  0
	X++：X 加 1 ，X =  0 + 1 =  1

示例 2：

	输入：operations = ["++X","++X","X++"]
	输出：3
	解释：操作按下述步骤执行： 
	最初，X = 0
	++X：X 加 1 ，X = 0 + 1 = 1
	++X：X 加 1 ，X = 1 + 1 = 2
	X++：X 加 1 ，X = 2 + 1 = 3

示例 3：

	输入：operations = ["X++","++X","--X","X--"]
	输出：0
	解释：操作按下述步骤执行：
	最初，X = 0
	X++：X 加 1 ，X = 0 + 1 = 1
	++X：X 加 1 ，X = 1 + 1 = 2
	--X：X 减 1 ，X = 2 - 1 = 1
	X--：X 减 1 ，X = 1 - 1 = 0

 

提示：

* 1 <= operations.length <= 100
* operations[i] 将会是 "++X"、"X++"、"--X" 或 "X--"

