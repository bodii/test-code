#### Volatile域
* 多处理器的计算机能够暂时在寄存器或本地内存缓冲区中保存内存中的值。结果是，运行在不同处
理器上的线程可能在同一内存位置位置取到不同的值。
* 编译器可以改变指令执行的顺序以使吞吐量最大化。这种顺序上的变化不会改变代码语义，但是编
译器假定内存的值仅仅在代码中有显式的修改指令时才会改变。然而，内存的值可以被另一个线程改
变!

同步格言: 如果向一个变量写入值，而这个变量接下来可能会被另一个线程读取，或者，从一个变量
读值，而这个变量可能是之前被另一个线程写入的，此时必须使用同步。

volatile关键字为实例域的同步访问提供了一种免锁机制。如果声明一个域为volatile,那么编译器和
虚拟机就知道该域是可能被另一个线程并发更新的。
例如，假定一个对象有一个布尔标记done,它的值被一个线程设置却被另一个线程查询，如同我们讨论
的那样，你可以使用锁:
```java
private boolean done;
public synchronized boolean isDone() { return done; }
public synchronized void setDone() { done = true; }
```
或许使用内部锁不是个好主意。如果另一个线程已经对该对象加锁，isDone和setDone方法可能阻塞。
如果注意到这个方面，一个线程可以为这一变量使用独立的Lock。但是，这也会带来许多麻烦。
在这种情况下，将域声明为volatile是合理的:
```java
private volatile boolean done;
public boolean isDone() { return done; }
public void setDone() { done = true; }
```
Volatile变量不能提供原子性。例如，方法：
```java
public void flipDone() { done = !done; }
```
不能确保翻转域中的值。不能保证读取、翻转和写入不被中断。

