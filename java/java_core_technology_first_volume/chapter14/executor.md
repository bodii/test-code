#### 执行器
构建一个新的线程是有一定代价的，因为涉及与操作系统的交互。如果程序中创建了大量的生命期很短的线程，应该使用`线程池(thread pool)`。一个线程池中包含许多准备运行的空闲线程。将Runnable对象交给线程池，就会有一个线程调用run方法。当run方法退出时，线程不会死亡，而是在池中准备为下一个请求提供服务。

另一个使用线程池的理由是减少并发线程的数目。创建大量线程会大大降低性能甚至使虚拟机崩溃。如果有一个会创建许多线程的算法，应该使用一个线程数"固定的"线程池以限制并发线程的总数。

`执行器(Executor)`类有许多静态工厂方法用来创建线程池:

newCachedThreadPool                 必要时创建新线程；空间线程会被保留60秒
newFixedThreadPool                      该池包含固定数量的线程；空闲线程会一直被保留
newSingleThreadExecutor             只有一个线程的"池", 该线程顺序执行每一个提交的任务(类似于Swing事件分配线程)
newScheduledThreadPool              用于预定执行而构建的固定线程池，替代java.util.Timer
newSingleThreadScheduledExecutor       用于预定执行而构建的单线程"池"