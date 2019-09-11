#### java.util.NavigableSet
* E higher(E value)
* E lower(E value)
	返回大于value的最小元素或小于value的最大元素，如果没有这样的元素则返回null。

* E ceiling(E value)
* E floor(E value)
	返回大于等于value的最小元素或小于等于value的最大元素，如果没有这样的元素则返回null。

* E pollFirst()
* E pollLast()
	删除并返回这个集中的最大元素或最小元素，这个集为空时返回null

* Iterator<E> descendingIterator()
	返回一个按照递减顺序遍历集中元素的迭代器。

