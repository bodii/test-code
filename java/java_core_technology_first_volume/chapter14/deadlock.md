#### 死锁
有可能会因为每一个线程要等待更多的钱款存入而导致所有线程都被阻塞。这样的状态称为`死锁
(deadlock)`。
如果修改run方法，把每次转账至多$1000的限制去掉，死锁很快就会发生。
导致死锁的另一种途径是让第i个线程负责向第i个账户存钱，而不是从第i个账户取钱。这样一来，有
可能将所有的线程都集中到一个账户上，每一个线程都试图从这个账户中取出大于该账户余额的钱。
试试看。在SynchBankTest程序中，转用TransferRunnable类的run方法。在调用transfer时，交换
fromAccount和toAccount。运行该程序并查看它为什么会立即死锁。
还有一种很容易导致死锁的情况: 在SynchBankTest程序中，将signalAll方法转换为signal,会发现该
程序最终会挂起。signalAll通知所有等待增加资金的线程，与从不同的是signal方法仅仅对一个线程
解锁。如果该线程不能继续运行，所有的线程可能都被阻塞
