#### 线程池
newCachedThreadPool方法构建了一个线程池，对于每个任务，如果有空闭线程可用，立即让它执行任务，如果没有可用的空闲线程，则创建一个新线程。

newFixedThreadPool方法构建一个具有固定大小的线程池。如果提交的任务数多于空闲的线程数，那么把得不到服务的任务放置到队列中。当其他任务完成以后再运行它们。

newSingleThreadExecutor是一个退化了的大小为1的线程池：由一个线程执行提交的任务，一个接着一个。这3个方法返回实现了ExecutorService接口的ThreadPoolExecutor类的对象。
可用下面的方法之一将一个Runnable对象或Callable对象提交给ExecutorService:
```java
Future<?> submit(Runnable task)
Future<T> submit(Runnable task, T result)
Future<T> submit(Callable<T> task)
```
该池会在方便的时候尽早执行提交的任务。调用submit时，会得到一个Future对象，可用来查询该任务的状态。
第一个submit方法返回一个奇怪样子的Future<?>。可以使用这样一个对象来调用isDone、cancel或isCancelled。但是，get方法在完成的时候只是简单地返回null。
第二个版本的submit也提交一个Runnable,并且Future的get方法在完成的时候返回指定的result对象。
第三个版本的submit提交一个Callable,并且返回的Future对象将在计算结果准备好的时候得到它。
当用完一个线程池的时候，调用shutdown。该方法启动该池的关闭序列。被关闭的执行器不再接受新的任务。当所有任务都完成以后，线程池中的线程死亡。另一种方法是调用shutdownNow。该池取消尚未开始的所有任务并试图中断正在运行的线程。

下面总结了在使用连接池时应该做的事:
1. 调用Executors类中静态的方法newCachedThreadPool或newFixedThreadPool。
2. 调用submit提交Runnable或Callable对象。
3. 如果想要取消一个任务，或如果提交Callable对象，那就要保存好返回的Future对象。
4. 当不再提交任何任务时，调用shutdown。