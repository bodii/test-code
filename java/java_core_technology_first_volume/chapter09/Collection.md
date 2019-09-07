#### java.util.Collection<E>
* Iterator<E> iterator()
	返回一个用于访问集合中每个元素的迭代器。

* int size()
	返回当前存储在集合中的元素个数。

* boolean isEmpty()
	如果集合中没有元素，返回true.

* boolean contains(Object obj) 
	如果集合中包含一个与obj相等的对象，返回true.

* boolean containsAll(Collection<?> other)
	如果这个集合包含other集合中的所有元素，返回true.

* boolean add(Object element)
	将一个元素添加到集合中。如果由于这个调用改变了集合，返回true.

* boolean addAll(Collection<? extends E> other)
	将other集合中的所有元素添加到这个集合。如果由于这个调用改变了集合，返回true.

* boolean remove(Object obj)
	从这个集合中删除等于obj的对象。如果有匹配的对象被删除，返回true.

* boolean removeAll(Collection<?> other)
	从这个集合中删除other集合中存在的所有元素。如果由于这个调用改变了集合，返回true.

* default boolean removeIf(Predicate<? super E> filter)
	从这个集合删除filter返回true的所有元素。如果由于这个调用改变了集合，则返回true.

* void clear()
	从这个集合中删除所有的元素。

* boolean retainAll(Collection<?> other)
	从这个集合中删除所有与other集合中的元素不同的元素。如果由于这个调用改变了集合，
	返回true.

* Object[] toArray()
	返回这个集合的对象数组.

* <T> T[] toArray(T[] arrayToFill)
	返回这个集合的对象数组。如果arrayToFill足够大，就将集合中的元素填入这个数组中。
	剩余空间填补null;否则，分配一个新数组，其成员类型与arrayToFill的成员类型相同，
	其长度等于集合的大小，并填充集合元素。
