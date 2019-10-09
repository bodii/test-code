#### 监视器概念
锁和条件是线程同步的强大工具，但是，严格地讲，它们不是面向对象的。在不需要加锁的情况下，
保证多线程的安全性。
> 最成功的解决方案之一是`监视器(monitor)`，监视器具有如下特性:
* 监视器是只包含私有域的类。
* 每个监视器对所有的方法进行回锁。换句话说，如果客户端调用obj.method(),那么obj对象的锁是在
方法调用开始时自动获得，并且当方法返回时自动释放该锁。因为所有的域是私有的，这样的安排可以
确保一个线程在对对象操作时，没有其他线程能访问该域。
* 该锁可以有任意多个相关条件。

监视器的早期版本只有单一的条件，使用一种很优雅的句法。可以简单地调用await accounts[from] >
= balance而不使用任何显式的条件变量。然而，研究表明盲目地重新测试条件是低效的。显式的条件
变量解决了这一问题。每一个条件变量管理一个独立的线程程集。
Java中的每个对象有一个内部的锁和内部的条件。

如果一个方法用synchronized关键字声明，那么，它表现的就像是一个监视器方法。通过调用wait/not-
ifyAll/notify来访问条件变量。
然而，在下述的3个方面Java对象不同于监视器，从而使得线程的安全性下降:
* 域不要求必须是private。
* 方法不要求必须是synchronized。
* 内部锁对客户是可用的。