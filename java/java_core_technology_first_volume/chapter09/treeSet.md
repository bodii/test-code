#### 树集
TreeSet类与散列集十分类似，不过，这比散列集有所改过。树集是一个有序集合(sorted collection)
可以以任意顺序将元素插入到集合中。在对集合进行遍历时，每个值将自动地按照排序后的顺序呈现。
```java
SortedSet<String> sorter = new TreeSet<>();
sorter.add("Bob");
sorter.add("Amy");
sorter.add("Carl");
for (String s : sorter) System.println(s);
```

正如TreeSet类名所示，排序是用树结构完成的(当前实现使用的是红黑树(red-black tree))
每次将一个元素添加到树中时，都被放置在正确的排序位置上。因此，迭代器总是以排好序的
顺序访问每个元素。

将一个元素添加到树中要比添加到散列表中慢，但是，与检查数组或链表中的重复元素相比还是快
很多。如果树中包含n个元素，查找新元素的正确位置平均需要log2n次比较。例如，如果一棵树包含
了1000个元素，添加一个新元素大约需要比较10次。

> 要使用树集，必须能够比较元素。这些元素必须实现Comparable接口，或者构造集是必须提供一个
Comparator。

