### 面试题 05.01. Insert Into Bits LCCI
You are given two 32-bit numbers, N and M, and two bit positions, i and j. Write a method to insert M into N such that M starts at bit j and ends at bit i. You can assume that the bits j through i have enough space to fit all of M. That is, if M = 10011, you can assume that there are at least 5 bits between j and i. You would not, for example, have j = 3 and i = 2, because M could not fully fit between bit 3 and bit 2.

Example1:

	Input: N = 10000000000, M = 10011, i = 2, j = 6
	Output: N = 10001001100

Example2:

	Input:  N = 0, M = 11111, i = 0, j = 4
	Output: N = 11111

----

### 面试题 05.01. 插入
给定两个整型数字 N 与 M，以及表示比特位置的 i 与 j（i <= j，且从 0 位开始计算）。

编写一种方法，使 M 对应的二进制数字插入 N 对应的二进制数字的第 i ~ j 位区域，不足之处用 0 补齐。具体插入过程如图所示。

题目保证从 i 位到 j 位足以容纳 M， 例如： M = 10011，则 i～j 区域至少可容纳 5 位。



示例1:

	输入：N = 1024(10000000000), M = 19(10011), i = 2, j = 6
	输出：N = 1100(10001001100)

示例2:

	输入： N = 0, M = 31(11111), i = 0, j = 4
	输出：N = 31(11111)

