### 面试题 03.01. Three in One LCCI
Describe how you could use a single array to implement three stacks.

You should implement push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum) methods. stackNum is the index of the stack. value is the value that pushed to the stack.

The constructor requires a stackSize parameter, which represents the size of each stack.

Example1:

	 Input:
	["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
	[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
	 Output:
	[null, null, null, 1, -1, -1, true]
	Explanation: When the stack is empty, `pop, peek` return -1. When the stack is full, `push` does nothing.

Example2:

	 Input:
	["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
	[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
	 Output:
	[null, null, null, null, 2, 1, -1, -1]


----

### 面试题 03.01. 三合一
三合一。描述如何只用一个数组来实现三个栈。

你应该实现push(stackNum, value)、pop(stackNum)、isEmpty(stackNum)、peek(stackNum)方法。stackNum表示栈下标，value表示压入的值。

构造函数会传入一个stackSize参数，代表每个栈的大小。

示例1:

 输入：
	["TripleInOne", "push", "push", "pop", "pop", "pop", "isEmpty"]
	[[1], [0, 1], [0, 2], [0], [0], [0], [0]]
	 输出：
	[null, null, null, 1, -1, -1, true]
	说明：当栈为空时`pop, peek`返回-1，当栈满时`push`不压入元素。

示例2:

	 输入：
	["TripleInOne", "push", "push", "push", "pop", "pop", "pop", "peek"]
	[[2], [0, 1], [0, 2], [0, 3], [0], [0], [0], [0]]
	 输出：
	[null, null, null, null, 2, 1, -1, -1]

提示：

* 0 <= stackNum <= 2


