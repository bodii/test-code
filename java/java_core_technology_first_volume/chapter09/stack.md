#### 栈
Stack类扩展为Vector类，它可以让栈使用不属于栈操作的insert和remvoe方法，即可以在
任何地方进行插入或删除操作，而不仅仅是在栈顶。

#### java.util.Stack
* E push(E item)
	将item压入栈并返回item.

* E pop()
	弹出并返回栈顶的item.如果栈为空，请不要调用这个方法。

* E peek()
	返回栈顶元素，但不弹出。如果栈为空，请不要调用这个方法。
