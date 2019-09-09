#### 链表
从数组的中间位置删除一个元素要付出很大的代价，其原因是数组中处于被删除元素
之后的所有元素都要向数组的前端移动。在数组中间的位置上插入一个元素也是如些。
而`链表(linked list)`解决了这个问题。
尽管数组在连续的存储位置上存放对象引用，但链表却将每个对象存放在独立的结点中。
每个结点还存放着序列中下一个结点的引用。
在Java程序设计语言中，所有链表实际上都是`双向链接的(doubly linked)`，即每个
结点还存放着指向前驱结点的引用。
链表是一个有序集合(ordered collection),每个对象的位置十分重要。
链表不支持快速快随机访问。如果要查看链表中第n个元素，就必须从头开始，越过n-1个
元素。没有捷径可走。鉴于这个原因。在程序需要采用整数索引访问元素时，通常不选
用链表。

尽管如此，LinkedList类提供了一个用来访问某个特定元素的get方法:
```java
LinkedList<string> list = ...;
String obj = list.get(n);
```

如果链表中包含一个等于"Harry"的字符串，调用staff.contains("Harry")后将会返回true.

下面这段代码的效率极低:
```java
for (int i = 0; i < list.size(); i++)
	do something with list.get(i);
```