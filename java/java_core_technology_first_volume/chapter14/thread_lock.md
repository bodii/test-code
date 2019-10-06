#### 锁
锁是`可重入`的，因为线程可以重复地获得已经持有的锁。锁保持一个`持有计数(hold count)`来跟踪
对lock方法的嵌套调用。线程在每一次调用lock都要调用unlock来释放锁。由于这一特性，被一个锁保
护的代码可以调用另一个使用相同的锁的方法。

例如，transfer方法调用getTotalBalance方法，这也会封锁bankLock对象，此时bankLock对象的持有
计数为2。当getTotalBalance方法退出的时候，持有计数变回1。当transfer方法退出的时候，持有计
数变为0, 线程释放锁。

> 要留心临界区中的代码，不要因为异常的抛出而跳出临界区。如果在临界区代码结束之前抛出了异常，
finally子句将释放锁，但会使对象可能处于一种受损状态。

#### java.util.concurrent.locks.Lock
* void lock()
	获取这个锁；如果锁同时被另一个线程拥有则发生阻塞。
* void unlock()
	释放这个锁。


#### java.util.concurrent.locks.ReentrantLock
* ReentrantLock()
	构建一个可以被用来保护临界区的可重入锁。

* ReentrantLock(boolean fair)
	构造一个带有公平策略的锁。一个公平锁偏爱等待时间最长的线程。但是，这一公平的保证将大
	大降低性能。所以，默认情况下，锁没有被强制为公平的。


#### 条件对象
通常，线程进入临界区，却发现在某一条件满足之后它才能执行。要使用一个条件对象来管理那些已经
获得了一个锁但是却不能做有用工作的线程。条件对象经常被称为`条件变量(coniditional variable)`

一个锁对象可以有一个或多个相关的条件对象。你可以用newCondition方法获得一个条件对象。习惯上
给每一个条件对象命名为可以反映它所表达的条件的名字。
例如:
```java
class Bank {
	private Condition sufficientFnds;
	...
	public Bank() {
		...
		sufficientFunds = bankLock.newCondition();
	}
}
```
如果transfer方法发现余额不足，它调用
```java
sufficientFunds.await();
```
当前线程现在被阻塞了，并放弃了锁。我们希望这样可以使得另一个线程可以进行增加账户余额的操作。
等待获得锁的线程和调用await方法的线程存在本质上的不同。一旦一个线程调用await方法，它进入该
条件的等待集。当锁可用时，该线程不能马上解除阻塞。相反，它处于阻塞状态，直到另一个线程调用
同一条件上的signalAll方法时为止。
当另一个线程转账时，它应该调用:
```java
sufficientFunds.signalAll();
```
