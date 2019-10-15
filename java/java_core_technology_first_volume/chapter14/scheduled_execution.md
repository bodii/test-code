#### 预定执行
ScheduledExecutorService接口具有为预定执行(Scheduled Execution) 或重复执行任务而设计的方法。
它是一种允许使用线程池机制的java.util.Timer的泛化。
Exceutors类的newScheduledThreadPool和newSingleThreadScheduledExecutor方法将返回实现了Sched-
uledExecutorService接口的对象。
可以预定的Runnable或Callable在初始的延迟之后只运行一次。也可以预定一个Runnable对象周期性地
运行。

#### java.util.concurrent.Executors
* ScheduledExecutorService newScheduledThreadPool(int threads)
	返回一个线程池，它使用给定的线程数来调试任务。

* ScheduledExecutorService newSingleThreadScheduledExecutor()
	返回一个执行器，它在一个单独线程中调度任务。

#### java.util.concurrent.ScheduledExecutorService
* ScheduledFuture<V> schedule(Callable<V> task, long time, TimeUnit unit)
* ScheduledFuture<?> schedule(Runnable task, long time, TimeUnit unit)
	预定在指定的时间之后执行任务

* ScheduledFuture<?> scheduleAtFixedRate(Runnable task, long initialDelay, long period,
		TimeUnit unit)
	预定在初始的延迟结束后，周期性地运行给定的任务，周期长度是period。

* ScheduledFuture<?> scheduleWithFixedDelay(Runnable task, long inititalDelay, long delay,
		TimeUnit unit)
	预定在初始的延迟结束后周期性地运行给定的任务，在一次调用完成和下一次调用开始之间有长度
	为delay的延迟。

