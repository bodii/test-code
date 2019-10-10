#### 原子性
java.util.concurrent.atomic包中有很多类使用了很高效的机器级指令(而不是使用锁)来保证其他操
作的原子性。例如，AtomicInteger类提供了方法incrementAndGet和decrementAndGet，它们分别以原
子方式将一个整数自增或自减。例如，可以安全地生成一个数值序列，如下所示:
```java
public static AtomicLong nextNumber = new AtomicLong();
long id = nextNumber.incrementAndGet();
```
incrementAndGet方法以原子方式将AtomicLong自增，并返回自增后的值。也就是说，获得值、增1并设
置然后生成新值的操作不会中断。可以保证即使是多个线程并发地访问同一个实例，也会计算并返回正
确的值。
有很多方法可以以原子方式设置和增减值，不过，如果希望完成更复杂的更新，就必须使用compareAnd-
Set方法。例如，假设希望跟踪不同线程观察的最大值。下面的代码是不可行的：
```java
public static AtomicLong largest = new AtomicLong();
largest.set(Math.max(largest.get(), observed));
```
如果另一个线程也在更新larget,就可能阻止这个线程更新。这样一来，compareAndSet会返回false,
而不会设置新值。在这种情况下，循环会更次尝试，读取更新后的值，并尝试修改。最终，它会成功
地用新值替换原来的值。这听上去有些麻烦，不过compareAndSet方法会映射到一个处理器操作，比使
用锁速度更快。
lambda表达式更新变更，它会为你完成更新。对于这个例子，我们可以调用:
```java
largest.updateAndGet(x -> Math.max(x, observed));
// or
largest.accumulateAndGet(observed, Math::max);
```
accumulateAndGet方法利用一个二元操作符来合并原子值和所提供的参数。

