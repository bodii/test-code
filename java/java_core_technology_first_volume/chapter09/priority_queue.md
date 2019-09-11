#### 优先级队列
优先级队列(priority queue)中的元素可以按照任意的顺序插入，却总是按照排序的顺序进行检索。
也就是说，无论何时调用remove方法，总会获得当前优先级队列中最小的元素。然而，优先级队列并
没有对所有元素时行排序。如果用迭代的方式处理这些元素，并不需要对它们进行排序。

优先级队列使用了一个优雅且高效的数据结构，称为`堆(heap)`.
堆是一个可以自我调整的二叉树，对树执行添加(add)和删除(remove)操作，可以让最小的元素移动到
根，而不必花费时间对元素进行排序。


#### java.util.PriorityQueue
* PriorityQueue()
* PriorityQueue(int initialCapacity)
	构造一个用于存放Comparable对象的优先级队列。

* PriorityQueue(int initialCapacity, Comparator<? super E> c)
	构造一个优先级队列，并用指定的比较器对元素进行排序。
