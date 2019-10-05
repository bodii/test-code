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

