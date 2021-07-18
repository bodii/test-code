### 1114. Print in Order 
Suppose we have a class:

```
public class Foo {
	public void first() { print("first"); }
	public void second() { print("second"); }
	public void third() { print("third"); }
}
```

The same instance of Foo will be passed to three different threads. Thread A will call first(), thread B will call second(), and thread C will call third(). Design a mechanism and modify the program to ensure that second() is executed after first(), and third() is executed after second().

Note:

We do not know how the threads will be scheduled in the operating system, even though the numbers in the input seem to imply the ordering. The input format you see is mainly to ensure our tests' comprehensiveness.



Example 1:

	Input: nums = [1,2,3]
	Output: "firstsecondthird"
	Explanation: There are three threads being fired asynchronously. The input [1,2,3] means thread A calls first(), thread B calls second(), and thread C calls third(). "firstsecondthird" is the correct output.

Example 2:

	Input: nums = [1,3,2]
	Output: "firstsecondthird"
	Explanation: The input [1,3,2] means thread A calls first(), thread B calls third(), and thread C calls second(). "firstsecondthird" is the correct output.

----

### 1114. 按序打印 
我们提供了一个类：

public class Foo {
	public void first() { print("first"); }
	public void second() { print("second"); }
	public void third() { print("third"); }
}

三个不同的线程 A、B、C 将会共用一个 Foo 实例。

一个将会调用 first() 方法
一个将会调用 second() 方法
还有一个将会调用 third() 方法

请设计修改程序，以确保 second() 方法在 first() 方法之后被执行，third() 方法在 second() 方法之后被执行。



示例 1:

	输入: [1,2,3]
	输出: "firstsecondthird"
	解释:
	有三个线程会被异步启动。
	输入 [1,2,3] 表示线程 A 将会调用 first() 方法，线程 B 将会调用 second() 方法，线程 C 将会调用 third() 方法。
	正确的输出是 "firstsecondthird"。

示例 2:

	输入: [1,3,2]
	输出: "firstsecondthird"
	解释:
	输入 [1,3,2] 表示线程 A 将会调用 first() 方法，线程 B 将会调用 third() 方法，线程 C 将会调用 second() 方法。
	正确的输出是 "firstsecondthird"。

