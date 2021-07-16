### 155. Min Stack
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

* MinStack() initializes the stack object.
* void push(val) pushes the element val onto the stack.
* void pop() removes the element on the top of the stack.
* int top() gets the top element of the stack.
* int getMin() retrieves the minimum element in the stack.



Example 1:

	Input
	["MinStack","push","push","push","getMin","pop","top","getMin"]
	[[],[-2],[0],[-3],[],[],[],[]]

	Output
	[null,null,null,null,-3,null,0,-2]

	Explanation
	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin(); // return -3
	minStack.pop();
	minStack.top();    // return 0
	minStack.getMin(); // return -2



Constraints:

* -2^31 <= val <= 2^31 - 1
* Methods pop, top and getMin operations will always be called on non-empty stacks.
* At most 3 * 104 calls will be made to push, pop, top, and getMin.

----

### 155. 最小栈
设计一个支持 push ，pop ，top 操作，并能在常数时间内检索到最小元素的栈。

* push(x) —— 将元素 x 推入栈中。
* pop() —— 删除栈顶的元素。
* top() —— 获取栈顶元素。
* getMin() —— 检索栈中的最小元素。



示例:

	输入：
	["MinStack","push","push","push","getMin","pop","top","getMin"]
	[[],[-2],[0],[-3],[],[],[],[]]

	输出：
	[null,null,null,null,-3,null,0,-2]

	解释：
	MinStack minStack = new MinStack();
	minStack.push(-2);
	minStack.push(0);
	minStack.push(-3);
	minStack.getMin();   --> 返回 -3.
	minStack.pop();
	minStack.top();      --> 返回 0.
	minStack.getMin();   --> 返回 -2.

