#### 多线程(multithreaded)
AWT事件分派线程(event dispatch thread)将一直地并行运行，以处理用户界面的事件。

下面是在一个单独的线程中执行一个任务的简单过程:
1. 将任务代码移到实现了Runnable接口的类的run方法中:
```java
@FunctionalInterface
public interface Runnable {
	void run();
}
```
由于Runnable是一个函数式接口，可以用lambda表达式建立一个实例:
```java
Runnable r = () -> { task code };
```

2. 由Runnable创建一个Thread对象:
```java
Thread t = new Thread(r);
```

3. 启动线程:
```java
t.start();
```


> 也可心通过构造一个Thread类的子类定义一个线程:
```java
class MyThread extends Thread {
	public void run() {
		task code
	}
}
```
然后，构造一个子类的对象，并调用start方法。不过，这种方法已不再推荐。应该将要并行
运行的任务与运行机制耦合。如果有很多任务，要为每个任务创建一个独立的线程所付出的代
价太大了。可以使用线程池来解决这个问题。


#### java.lang.Thread
* Thread(Runnable target) 
	构造一个新线程，用于调用给定目标的run()方法

* void start()
	启动这个线程，将引发调用run()方法。这个方法将立即返回，并且新线程将并发运行。

* void run()
	调用关联Runnable的run方法。


#### java.lang.Runnable
* void run()
	必须覆盖这个方法，并在这个方法中提供要执行的任务指令。


#### 中断线程
没有可以强制线程终止的方法。然而，interrupt方法可以用来请求终止终止线程。
当对一个线程调用interrupt方法时，线程的`中断状态`将被置位。这是每一个线程都具有的boolean标
志。每个线程都应该不时检查这个标志，以判断线程是否被中断。
要想弄清中断状态是否被置位，首先调用静态的`Thread.currentThread`方法获得当前线程,然后调
用isInterrupt方法:
```java
while (!Thread.currentThread().isInterrupted() && more work to do) {
	do more work
}
```
但是，如果线程被阻塞，就无法检测中断状态。这是产生InterruptedException异常的地方。当在一个
被阻塞的线程(调用sleep或wait)上调用interrupt方法时，阻塞调用将会被Interrupted Exception
异常中断。（存在不能被中断的阻塞I/O调用，应该考虑选择可中断的调用)
没有任何语言方面的需求要求一个被中断的线程应该终止。中断一个线程不过是引起它的注意。被中
断的线程可以决定如何响应中断。某些线程是如此重要以至于应该处理完异常后，继续执行，而不理
会中断。但是，更普遍的情况是，线程将简单地将中断作为一个终止的请求

这种线程的run方法具有如下形式:
```java
Runnable r = () -> {
	try {
		...
		while (!Thread.currentThread().isInterrupted() && more work to do) {
			do more work
		}
	} catch (InterruptedException e) {
		thread was interrupted during sleep or wait
	} finally {
		cleanup, if required
	}
	// exiting the run method terminates the thread
}
```

如果在每次工作迭代之后都调用sleep方法（或者其他的可中断方法), isInterrupted检测既没有必
要也没有用处。如果在中断状态被置位时调用sleep方法，它不会休眠。相反，它将清除这一状态(!)
并抛出InterruptedException。因此，如果你的循环调用sleep,不会检测中断状态。相反，要如下
所示捕获InterruptedExccption异常:
```java
Runnable r = () -> {
	try {
		while (more work to do) {
			do more work
			Thread.sleep(delay);
		}
	} catch (InterruptedException e) {
		// thread was interrupted during sleep
	} finally {
		cleanup, if required
	}
	// exiting the run method terminates the thread
}
```

在很多发布的代码中会发现InterruptedException异常被抑制在很低的层次上，像这样:
```java
void mySubTask() {
	...
	try { sleep(delay); }
	catch (InterruptedException e) {} // Don't ignore!
	...
}
```

不要这样做！如果不认为在catch子句中做这一处理有什么好处的话，仍然有两种合理选择:

* 在catch 子句中调用Thread.currentThread().interrupt()来设置中断状态。于是，调用者可以对
其进行检测。
void mySubTask() {
	try { sleep(delay); }
	catch (InterruptedException e) { Thread.currentThread().interrupt(); }
	...
}

* 或者，更好的选择是，用throws InterruptedException标记你的方法，不采用try语句块捕获异常。
于是，调用者(或者，最终在run方法)可以捕获这一异常。
```java
void mySubTask() throws InterruptedException {
	...
	sleep(delay);
	...
}
```


#### java.lang.Thread
* void interrupt()
	向线程发送中断请求。线程的中断状态将设置为true。如果目前该线程被一个sleep调用阻塞，那
	么，InterruptedException异常被抛出。

* static boolean interrupted()
	测试当前线程（即正在执行这一命令的线程)是否被中断。注意，这是一个静态方法。
	这一调用会产生副作用--它将当前线程的中断状态重置为false。

* boolean isInterrupted()
	测试线程是否被终止。不像静态的中断方法，这一调用不改变线程的中断状态。

* static Thread currentThread()
	返回代表当前执行线程的Thread对象。
