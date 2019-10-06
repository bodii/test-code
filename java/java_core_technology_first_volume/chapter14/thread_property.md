#### 线程属性
线程的各种属性：线程优先级、守护线程、线程组以及处理未捕获异常的处理器。

#### 线程优先级
在Java程序设计语言中，每一个线程有一个优先级。默认情况下，一个线程继承它的父线程的优先级。
可以用setPriority(在Thread类中定义为1)与MAX_PRIORITY(定义为10)之间的任何值。NORM_PRIORITY
被定义为5。
每当线程调度器有机会选择新线程时，它首先选择具有较高优先级的线程。但是，线程优先级是高度依
赖于系统的。当虚拟机信赖于宿主机平台的线程实现机制时，Java线程的优先级被映射到宿主机平台的
优先级上，优先级个数也许更多，也许更少。
例如，Windows有7个优先级别。一些Java优先级将映射到相同的操作系统优先级。在Oracle为Linux提供
的Java虚拟机中，线程的优先级被忽略－－所有线程具有相同的优先级。
初级程序员常常过度使用线程优先级。为优先级而烦恼是事出有因的。不要将程序构建为功能的正确性
依赖于优先级。

> 如果确实要使用优先级，应该避免初学者常规的一个错误。如果有几个高优先级的线程没有进入非活
动状态，低优先级的线程可能永远也不能执行。每当调度器决定运行一个新线程时，首先会在具有高优
先级的线程中进行选择，尽管这样会使低优先级的线程完全饿死。

#### java.lang.Thread
* void setPriority(int newPriority)
	设置线程的优先级。优先级必须在Thread.MIN_PRIORITY与Thread.MAX_PRIORITY之间。一般使用
	Thread.NORM_PRIORITY优先级。

* static int MIN_PRIORITY
	线程的最小优先级。最小优先级的值为1。

* static int NORM_PRIORITY
	线程的默认优先级。默认优先级为5。

* static int MAX_PRIORITY
	线程的最高优先级。最高优先级的值为10。

* static void yield()
	导致当前执行线程处于让步状态。如果有其他的可运行线程具有至少与此线程同样高的优先级，那
	么这些线程接下来会被调度。注意，这是一个静态的方法。


#### 守护线程
可以通过调用：
```java
t.setDaemon(true);
```
将线程转换为`守护线程(daemon thread)`。这样一个线程没有什么神奇。守护线程的唯一用途是为其
他线程提供服务。计时线程就是一个例子，它定时地发送“计时器嘀嗒”信号给其他线程或清空过时的高
速缓存项的线程。当只剩下守护线程时，虚拟机就退出了，由于如果只剩下守护线程，就没必要继续运
行程序了。
守护线程有时会被初学者错误地使用，他们不打算考虑关机(shutdown)动作。但是，这是很危险的。守
护线程应该永远不去访问固有资源，如文件、数据库，因为它会在任何时候甚至在一个操作的中间发生
中断。

#### java.lang.Thread
* void setDaemon(boolean isDaemon)
	标识该线程为守护线程或用户线程。这一方法必须在线程启动之前调用。


#### 未捕获异常处理器
线程的run方法不能抛出任何爱查异常，但是，非受查异常会导致线程终止。在这种情况下，线程就死亡
了。
但是，不需要任何catch子句一来处理可以被传播的异常。相反，就在线程死亡之前，异常被传递到一
个用于未捕获异常的处理器。
该处理器必须属于一个实现Thread.UncajughtExceptionHandler接口的类。这个接口只有一个方法。
```java
void uncaughtException(Thread t, Throwable e)
```
可以用setUncaughtExceptionHandler方法为任何线程安装一个处理器。也可以用Thread类的静态方法
setDefaultUncaughtExceptionHandler为所有线程安装一个默认的处理器。替换处理器可以使用日志
API发送未捕获异常的报告到日志文件。
如果不安装默认的处理器，默认的处理器为空。但是，如果不为独立的线程安装处理器，此时的处理器
就是该线程的ThreadGroup对象。

> 线程组是一个可以统一管理的线程集合。默认情况下，创建的所有线程属于相同的线程组，但是，也
可能会建立其他的组。现在引入了更好的特性用于线程集合的操作，所以建议不要在自己的程序中使用
线程组。

ThreadGroup类实现Thread.UncaughtExceptionHandler接口。它的uncaughtException方法如下操作:
1. 如果该线程有父线程组，那么父线程组的uncaughtException方法被调用。
2. 否则，如果Thread.getDefaultExceptionHandler方法返回一个非空的处理器，则调用该处理器。
3. 否则，如果Throwable是ThreadDeath的一个实例，什么都不做。
4. 否则，线程的名字以及Throwable的栈轨迹被输出到System.err上。
这是你在程序中肯定看到过许多次的栈轨迹。


#### java.lang.ThreadGroup
* void uncaughtException(Thread t, Throwable e)
	如果有父线程组，调用父线程组的这一方法；或者，如果Thread类有默认处理器，调用该处理器，
	否则，输出栈轨迹到标准错误流上(但是，如果e是一个ThreadDeath对象，栈轨迹是被禁用的。
	ThreadDeath对象由stop方法产生，而该方法已经过时)。

