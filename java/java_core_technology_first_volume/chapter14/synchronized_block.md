#### 同步阻塞
> 每一个Java对象有一个锁。
线程可以通过调用同步方法获得锁。还有另一种机制可以获得锁，通过进入一个同步阻塞。
当线程进入如下形式的阻塞:
```java
synchronized (obj) { // this is the syntax for a synchronized block
	critical section
}
```
于是它获得obj的锁。
有时会发现“特殊的”锁，例如：
```java
public class Bank {
	private double[] accounts;
	private Object lock = new Object();

	...
	public void transfer(int from, int to, int amount) {
		synchronized (lock) { // an ad-hoc lock
			accounts[from] -= amount;
			accounts[to] += amount;
		}
		System.out.println(...);
	}
}
```
在此，lock对象被创建仅仅是用来使用每个Java对象持有的锁。
有时程序员使用一个对象的锁来实现额外的原子操作，实际上称为`客户端锁定(clientside locking)`。

例如，考虑Vector类，一个列表，它的方法是同步的。现在，假定在Vector<Double>中存储银行余额。
这圼有一个transfer方法的原始实现:
```java
public void transfer(Vector<Double> accounts, int from, int to, int amount) { // Error
	accounts.set(from, accounts.get(from) - amount);
	accounts.set(to, accounts.et(to) + amount);
	System.out.print(...);
}
```
Vector类的get和set方法是同步的，但是，这对于我们没有什么帮助。在第一次对get的调用已经完成
之后，一个线程完全可能在transfer方法中被剥夺运行权。于是，另一个线程可能在相同的存储位置存
入不同的值。但是，我们可以截获这个锁:
```java
public void transfer(Vector<Double> accounts, int from, int to, int ammount) {
	synchronized(accounts) {
		accounts.set(from, accounts.get(from) - amount);
		accounts.set(to, accounts.get(to) + amount);
	}
	System.out.println(...);
}
```
这个方法可以工作，但是它完全依赖于这样一个事实，Vector类对自己的所有可修改方法都使用内部锁。
然而，这是真的吗？Vector类的文档没有给出这样的承诺。不得不仔细研究源代码并希望将来的版本能
介绍非同步的可修改方法。如你所见，客户端锁定是非常脆弱的，通常不推荐使用。


