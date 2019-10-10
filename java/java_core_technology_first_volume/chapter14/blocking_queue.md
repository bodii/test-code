#### 阻塞队列
对于许多线程问题，可以通过使用一个或多个队列以优雅且安全的方式将其形式化。生产者线程向队列
插入元素，消息者线程则取出它们。使用队列，可以安全地从一个线程向另一个线程传递数据。
当试图向队列添加元素而队列已满，或是想从队列移出元素而队列为空的时候，`阻塞队列(blocking 
queue)`导致线程阻塞。在协调多个线程之间的合作时，阻塞队列是一个有用的工具。
工作者线程可以周期性地将中间结果存储在阻塞队列中。其他的工作者线程移出中间结果并进一步加以
修改。队列会自动地平衡负载。如果第一个线程集运行得比第二个慢，第二个线程集在等待结果时会阻
塞。如是第一个线程集运行得快，它将等待第二个队列集赶上来。

阻塞队列方法：
add		添加一个元素
element	返回队列的头元素
offer	添加一个元素并返回true
peek	返回队列的头元素
poll	移出并返回队列的头元素
put		添加一个元素
remove	移出并返回头元素
take	移出并返回头元素

阻塞队列方法分为以下3类,这取决于当队列满或空时它们的响应方式。如果将队列当作线程管理工具来
使用，将要用到put或take方法。当试图向满的队列中添加或从空的队列中移出元素时，add、remove和
element操作抛出异常。当然，在一个多线程程序中，队列会在任何时候空或满，因此，一定要使用off-
er、poll和peek方法作为替代。这些方法如果不能完成任务，只是给出一个错误提示而不会抛出异常。

带有超时的offer方法和poll方法的变体:
```java
boolean success = q.offer(x, 100, TimeUnit.MILLISECONDS);
```
尝试在100毫秒的时间内在队列的尾部插入一个元素。如果成功返回true;否则，达到超时时，返回false.
类似地，下面的调用:
```java
Object head = q.poll(100, TimeUnit.MILLISECONDS);
```
尝试用100毫秒的时间移除队列的头元素；如果成功返回头元素，否则，达到在超时时，返回null。
如果队列满，则put方法阻塞；如果队列空，则take方法阻塞。在不带超时参数时，offer和poll方法等
效。
java.util.concurrent包提供了阻塞队列的几个变种。默认情况下，LinkedBlockingQueue的容量是没有
上边界的，但是，也可以选择指定最大容量。LinkedBlockingDeque是一个双端的版本。ArrayBlocking-
Queue在构造时需要指定容量，并且有一个可选的参数来指定是否需要公平性。若设置了公平参数，则
那么等待了最长时间的线程会优先得到处理。通常，公平性会降低性能，只有在确实非常需要时才使用
它。
PriorityBlockingQueue是一个带优先级的队列，而不是先进先出队列。元素按照它们的优先级顺序被
移出。该队列是没有容量上限，但是，如果队列是空的，取元素的操作会阻塞。
最后，DelayQueue包含实现Delayed接口的对象:
```java
interface Delayed extends Comparable<Delayed> {
	long getDelay(TimeUnit unit);
}
```
getDelay方法返回对象的残留延迟。负值表示延迟已经结束。元素只有在延迟用完的情况下才能从Delay-
Queue移除。还必须实现compareTo方法。DelayQueue使用该方法对元素时行排序。
Java SE 7增加了一个TransferQueue接口，允许生产者线程等待，直到消费者准备就绪可以接收一个元
素。
```java
q.transfer(item);
```
这个调用会阻塞，直到另一个线程将元素(item)删除。LinkedTransferQueue类实现了这个接口。
