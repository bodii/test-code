#### 并发集视图
假设你想要的是一个大的线程安全的集而不是映射。并没有一个ConurrentHashSet类，而且你肯定不想
自己创建这样一个类。当然，可以使用ConcurrentHashMap（包含"假"值），不过这会得到一个映射而
不是集，而且不能应用Set接口的操作。

静态newKeySet方法会生成一个Set<K>, 这实际上是ConcurrentHashMap<K, Boolean> 的一个包装器。
（所有映射值都为Boolean.TRUE, 不过因为只是要把它用作一个集，所以并不关心具体的值。)
```java
Set<String> words = ConcurrentHashMap.<String>newKeySet();
```
当然，如果原来有一个映射，keySet方法可以生成这个映射的键集。这个集是可变的。如果删除这个
集的元素，这个键(以及相应的值)会从映射中删除。不过，不能向键集增加元素，因为没有相应的值
可以增加。
Java SE 8 为ConcurrentHashMap增加了第二个keySet方法，包含一个默认值，可以在为集增加元素时
使用:
```java
Set<String> words = map.keySet(1L);
words.add("Java");
```
如果“Java"在words中不存在，现在它会有一个值1。
