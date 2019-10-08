#### synchronized关键字
有关锁和条件的关键之处:
* 锁用来保护代码片段，任何时刻只能有一个线程执行被保护的代码。
* 锁可以管理试图进入被保护代码段的线程。
* 锁可以拥有一个或多个相关的条件对象。
* 每个条件对象管理那些已经进入被保护的代码段但还不能运行的线程。

> Lock和Condition接口为程序设计人员提供了高度的锁定控制。
从1.0版本开始，Java中的每个对象都有一个内部锁。
> 如果一个方法用synchronized关键字声明, 那么对象的锁将保护整个方法。也就是说，要调用该方法，线程必须获得内部的对象锁。

换句话说:
```java
public synchronized void method() {
    method body
}

// 等价于:
public void method() {
    this.intrinsicLock.lock();
    try {
        method body
    } finally {
        this.intrinsicLock.unlock();
    }
}
```

内部对象锁只有一个相关条件。wait方法添加一个线程到等待集中，notifyAll/notify 方法解除等待线程的阻塞状态。换句话说，调用wait或notifyAll等价于
```java
intrinsicCondition.await();
intrinsicCondition.signalAll();
```

wait、notifyAll以及notify方法是Object类的final方法。Condition方法必须被命名为await、signalAll和signal以便它们不会与那些方法发生冲突。

```java
public class Bank {
    private double[] accounts;

    public synchronized void transfer(int from, int to, double amount) throws InterruptedException {
        while (accounts[fomr] <  amount)
            wait();

        accounts[from] -= amount;
        accounts[to] += amount;
        notfiyAll(); // notify all threads  waiting on the condition
    }

    public synchronized double getTotalBalance() { ... }
}
```

> 将静态方法声明为synchronized也是合法的。

内部锁和条件存在一些局限，包括：
* 不能中断一个正在试图获得锁的线程。
* 试图获得锁时不能设定超时。
* 每个锁仅有单一的条件，可能是不够的。
* 最好既不使用Lock/Condition也不使用synchronized关键字。在许多情况下你可以使用java.util.concurrent包中的一种机制，它会为你处理所有的加锁。
* 如果synchronized关键字适合你的程序，那么请尽量使用它，这样可以减少编写的代码数量，减少出错的几率。
* 如果特别需要Lock/Condition结构提供的独有特性时，才使用Lock/Condition。