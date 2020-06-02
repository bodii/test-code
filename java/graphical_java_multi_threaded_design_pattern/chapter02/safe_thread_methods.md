### 线程安全的方法
* synchronizedCollection 方法 
* synchronizedList 方法 
* synchronizedMap 方法 
* synchronizedSet 方法
* synchronizedSortedMap 方法 
* synchronizedSortedSet 方法


### 生存性与死锁
死锁是指两个线程分别持有着锁，并相互等待对方释放锁的现象。发生死锁的线程都无法再继续运行，程序也就失去了生存性。


### 在Single Threaded Execution 模式中，满足下列条件时，死锁就会发生。
1. 存在多个SharedResource角色
2. 线程在持有着某个SharedResource角色的锁的同时，还想获取其他SharedResource角色的锁
3. 获取SharedResource角色的锁的顺序并不固定(SharedResource角色是对称的)


### 原子操作
synchronized 方法只允许一个线程同时执行。如果某个线程正在执行synchronized方法，其他线程就无法时入该方法。也就是说，从多线程的观点来看，这个synchronized方法执行的操作是"不可分割的操作"。这种不可分割的操作通常称为原子(atomic)操作。

### long与double的操作不是原子的
由于long和double的赋值和引用是非原子操作，所以如果long字段和double字段是线程共享的，那么在对该字段执行操作时，就必须使用Single Threaded Execution模式。最简单的方法就是在synchronized方法中执行操作。

> `总结如下：`
* 基本类型、引用类型的赋值和引用是原子操作
* 但long和double的赋值和引用是非原子操作
* long或double在线程间共享时，需要将其放入synchronized中操作，或者声明为`volatile` 
> volatile 的作用并不仅仅是将long、double的赋值和引用操作变为原子的。

> java.util.concurrent.atomic 包提供了便于原子操作编程的类，如AtomicInteger、AtomicLong、
AtomicIntegerArray和AtomicLongArray等，这是通过封装volatile功能而得到的类库。

