#### java.util.List
* ListIterator<E> listIterator()
	返回一个列表迭代器，以便用来访问列表中的元素。

* ListIterator<E> listIterator(int index)
	返回一个列表迭代器，以便用来访问列表中的元素，这个元素是第一次调用next返回
	的给定索引的元素。

* void add(int i, E element)
	在给定位置添加一个元素。

* void addAll(int i, Collection<? extends E> elments)
	将某个集合中的所有元素添加到给定位置。

* E remove(int i)
	删除给定位置的元素并返回这个元素。

* E get(int i) 
	获取给定位置的元素。

* E set(int i, E element)
	用新元素取代给定位置的元素，并返回原来那个元素。

* int indexOf(Object element)
	返回与指定元素相等的元素在列表中第一次出现的位置，如果没有这样的元素将返回-1。

* int lastIndexOf(Object element)
	返回与指定元素相等的元素在列表中最后一次出现的位置，如果没有这样的元素将返回-1。


#### java.util.ListIterator
* void add(E newElement)
	在当前位置前添加一个元素。

* void set(E newElement)
	用新元素取代next或previous上次访问的元素。如果在next或previous上次调用之后列表结
	构被修改了，将抛出一个IllegalStateException异常。

* boolean hasPrevious()
	当反向迭代列表时，还有可供访问的元素，返回true.

* E previous()
	返回前一个对象。如果已经到达了列表的头部、就抛出一个NoSuchElmentException异常。

* int nextIndex()
	返回下一次调用next方法时将返回的元素索引。

* int previousIndex()
	返回下一次调用previous方法时将返回的元素索引。


#### java.util.LinkedList
* LinkedList()
	构造一个空链表

* LinkedList(Collection<? extends E> elements)
	构造一个链表，并将集合中所有的元素添加到这个链表中。

* void addFirst(E element)
* void addLast(E element)
	将某个元素添加到列表的头部和尾部。

* E getFirst()
* E getLast()
	返回列表头部或尾部的元素。

* E removeFirst()
* E removeLast()
	删除并返回列表头部或尾部的元素。
