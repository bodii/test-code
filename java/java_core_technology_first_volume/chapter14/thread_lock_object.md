#### 锁对象
有两种机制防止代码块受并发访问的干扰。
Java语句提供了一个`synchroinzed`关键字达到这一目的，并且Java SE 5.0引入了ReentrantLock类。
synchronized关键字自动提供一个锁以及相关的“条件”，对于大多数需要显式锁的情况，这是很便利的。

用ReentrantLock保护代码块的基本结构如下:
```java
myLock.lock(); // a ReentrantLock object
try {
	critical section
} finally {
	myLock.unlock(); // make sure the lock is unlocked even if an exception is thrown
}
```
这地结构确保任何时候只有一个线程进入临界区。一旦一个线程封锁了锁对象，其他任何线程都无法
通过lock语句。当其他线程调用lock时，它们被阻塞，直到第一个线程释放锁对象。

> 把解锁操作括在finally子句之内是至关重要的。如果在临界区的代码抛出异常，锁必须被释放。
否则，其他线程将永远阻塞。

如果使用锁，就不能使用带资源的try语句。首先，解锁方法名不是close。不过，即使将它重命名，
带资源的try语句也无法正常工作。它的首部希望声明一个新变量。但是如果使用一个锁，你可能想
使用多个线程共享的那个变量(而不是新变量)。

使用一个锁来保护Bank类的transfer方法:
```java
public class Bank {
	private Lock bankLock = new ReentrantLock();
	...
	public void transfer(int from, int to, int amount) {
		bankLock.lock();
		try {
			System.out.print(Thread.currentThread());
			accounts[from] -= amount;
			System.out.print(" %10.2f from %d to %d", amount, from, to);
			accounts[to] += amount;
			System.out.print(" Total Balance: %10.2f%n", getTotalBalance());
		} finally {
			bankLock.unlock();
		}
	}
}
```
假定一个线程调用transfer,在执行结束前被剥夺了运行权。假定第二个线程也调用transfer,由于第
二个线程不能获得锁，将在调用lock方法时被除害。它必须等待第一个线程完成transfer方法的执行
之后才能再度激活。当第一个线程释放锁时，那么第二个线程才能开始运行。
