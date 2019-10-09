#### final 变量
还有一种情况可以安全地访问一个共享域，即这个域声明为final时，考虑以下声明:
```java
final Map<String, Double> accounts = new HashMap<>();
```
其他线程会在构造函数完成构造之后才看到这个accounts变量。
如果不使用final,就不能保证其他线程看到的是accounts更新后的值，它们可能都只是看到null,而不
是新构造的HashMap。
当然，对这个映射表的操作并不是线程安全的。如果多个线程在读写这个映射表，仍然需要进行同步。

