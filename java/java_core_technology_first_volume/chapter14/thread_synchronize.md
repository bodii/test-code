#### 同步
在大多数实际的多线程应用中，两个或两个以上的线程需要共享对同一数据的存取。如果两个线程存
取相同的对象，并且每一个线程都调用了一个修改该对象状态的方法，将会发生什么呢？可以想象，线
程彼此踩了对方的脚。根据各线程访问数据的次序，可能会产生讹误的对象。这样一个情况通常称为
`竞争条件(race condition)`。


#### 竞争条件的一个例子
为了避免多线程引起的对共享数据的讹误，必须学习如何同步存取。

transfer方法的Bank类从一个账户转移一定数目的钱款到另一个账户(还没有考虑负的账户余额)。
```java
public void transfer(int from, int to, double amount) {
	System.out.print(Thread.currentThread());
	accounts[from] -= amount;
	System.out.printf(" %10.2f from %d to %d", amount, from, to);
	account[to] += amount;
	System.out.printf(" Total Balance: %10.2f%n", getTotalBalance());
}
```
这圼是Runnable类的代码。它的run方法不断地从一个固定的银行账户取出钱款。在每一次迭代中，
run方法随机选择一个目标账户和一个随机账户，调用bank对象的transfer方法然后眨眼。
```java
Runnable r = () -> {
	try {
		while (true) {
			int toAccount = (int) (bank.size() * Math.random());
			double amount = MAX_AMOUNT * Math.random();
			bank.transfer(fromAccount, toAccount, amount);
			Thread.sleep((int) (DELAY * Math.random()));
		}
	} catch (InterruptedException e) {

	}
}
```
当这个模拟程序运行时，不清楚在某一时刻某一银行账户中有多少钱。但是，知道所有账户的总金额应
该保持不变，因为所做的一切不过是从一个账户转移钱到另一个账户。
在每一次交易的结尾，transfer方法重新计算总值并打印出来。
本程序永远不会结束。只能按ctrl + c来终止这个程序。
在最初的交易中，银行的余额保持在$100 000，这是正确的，因为共100个账户，每个账户$1 000。但
是，过一段时间，余额问题有轻微的变化。当运行这个程序的时候，会发现有时候很快就出错了，有时
很长的时间后余额发生混乱。
