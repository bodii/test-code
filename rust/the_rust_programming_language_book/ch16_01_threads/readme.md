### 使用线程同时运行代码
在大部分现代操作系统中，已执行程序的代码在一个`进程(process)`中运行，操作系统则负责管理多个进程。在程序内部，也可以拥有多个同时运行的独立部分。运行这些独立部分的功能被称为`线程(threads)`。

将程序中的计算拆分进多个线程可以改善性能，因为程序可以同时进行多个任务，不过这也会增加复杂性。因为线程是同时运行的，所以无法预先保证不同线程中的代码的执行顺序。这会导致诸如此类的问题:
* 竞争状态(Race conditions)，多个线程以不一致的顺序访问数据或资源
* 死锁 (DeadLocks)，两个线程相互等待对方停止使用其所拥有的资源，这会阻止它们继续运行
* 只会发生在特定情况且难以稳定重现和修复的bug

### 使用spawn创建新线程
为了创建一个新线程，需要调用`thread::spawn`函数并传递一个闭包，并在其中包含希望的新线程运行的代码。

默认当主线程结束时，新线程也会结束，而不管其是否执行完毕。

### 使用join等待所有线程结束
可以通过将thread::spawn 的返回值储存在变量中来修复新建线程部分没有执行或者完全没有执行的问题。
thread::spawn的返回值类型是JoinHandle。JoinHandle是一个拥有所有权的值，当对其调用join方法时，它会等待其线程结束。

通过调用handle的join会阻塞当前线程直到handle所代表的线程结束。阻塞(Blocking)线程意味着阻止该线程执行工作或退出。

### 线程与move闭包
在参数列表前使用move关键字强制闭包获取其使用的环境值的所有权。这个技巧在创建新线程将值的所有权从一个线程移动到另一个线程时最为实用。

Rust是保守的并只会为线程借用v，这意味着主线程理论上可能使新建线程的引用无效。
move关键字覆盖了Rust默认保守的借用，但它不允许我们违反所有权规则
