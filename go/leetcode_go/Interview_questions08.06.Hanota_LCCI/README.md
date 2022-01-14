### 面试题 08.06. Hanota LCCI
In the classic problem of the Towers of Hanoi, you have 3 towers and N disks of different sizes which can slide onto any tower. The puzzle starts with disks sorted in ascending order of size from top to bottom (i.e., each disk sits on top of an even larger one). You have the following constraints:

(1) Only one disk can be moved at a time.
(2) A disk is slid off the top of one tower onto another tower.
(3) A disk cannot be placed on top of a smaller disk.

Write a program to move the disks from the first tower to the last using stacks.

Example1:

	Input: A = [2, 1, 0], B = [], C = []
	Output: C = [2, 1, 0]

Example2:

	Input: A = [1, 0], B = [], C = []
	Output: C = [1, 0]

Note:

* A.length <= 14

----
### 面试题 08.06. 汉诺塔问题
在经典汉诺塔问题中，有 3 根柱子及 N 个不同大小的穿孔圆盘，盘子可以滑入任意一根柱子。一开始，所有盘子自上而下按升序依次套在第一根柱子上(即每一个盘子只能放在更大的盘子上面)。移动圆盘时受到以下限制:
(1) 每次只能移动一个盘子;
(2) 盘子只能从柱子顶端滑出移到下一根柱子;
(3) 盘子只能叠在比它大的盘子上。

请编写程序，用栈将所有盘子从第一根柱子移到最后一根柱子。

你需要原地修改栈。

示例1:

	输入：A = [2, 1, 0], B = [], C = []
	输出：C = [2, 1, 0]

示例2:

	输入：A = [1, 0], B = [], C = []
	输出：C = [1, 0]

提示:

* A中盘子的数目不大于14个。

