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

这一调用重新激活因为这一条件而等待的所有线程。当这些线程从等待集当中移出时，它们再次成为可
运行的，调度器将现次激活它们。同时，它们将试图重新进入该对象。一旦锁成为可用的，它们中的某
个将从await调用返回，获得该锁并从被阻塞的地方继续执行。
此时，线程应该再次测试该条件。由于无法确保该条件被满足－－signalAll方法仅仅是通知正在等待的
线程：此时有可能已经满足条件，值得再次去检测该条件。

通常，对await的调用应该如下形式的循环体中:
```java
while (!(ok to proceed))
	condition.await();
```
至关重要的是最终需要某个其他线程调用signalAll方法。当一个线程调用await时，它没有办法重新激活
自身。它寄希望于其他线程。如果没有其他线程来重新激活等待的线程，它就永远不再运行了。这将导
致令人不快的`死锁(deadlock)`现象。如果所有其他线程被阻塞，最后一个活动线程在解除其他线程的
阻塞状态之前就调用await方法，那么它也被阻塞。没有任何线程可以解除其他线程的阻塞，那么该程序
就挂起了。

应该何时调用signalAll？经验上讲，在对象的状态有利于等待线程的方向改变时调用signalAll。例如，
当一个账户余额发生改变时，等待的线程会应该有机会检查余额。
```java
public void transfer(int from, int to, int amount) {
	bankLock.lock();
	try {
		while (accounts[from] < amount)
			sufficientFunds.await();
		...
		sufficientFunds.signalAll();
	} finally {
		bankLock.unlock();
	}
}
```
注意调用signalAll不会立即激活一个等待线程。它仅仅解除等待线程的阻塞，以便这些线程可以在当
前线程退出同步方法之后，通过竞争实现对对象的访问。
另一个方法signal,则是随机解除等待集中某个线程的阻塞状态。这比解除所有线程的阻塞更加有效，但
也存在危险。如果随机选择的线程发现自己仍然不能运行，那么它再次被阻塞。如果没有其他线程再次
调用signal,那么系统就死锁了。

> 当一个线程拥有某个条件的锁时，它仅仅可以在该条件上调用await、signalAll或signal方法。
