```java
public interface Queue<E> {
	void add(E element);
	E remove();
	int size();
}

public class CircularArrayQueue<E> implements Queue<E> {
	private int head;
	private int tail;

	CircularArrayQueue(int capacity) { ... }
	public void add(E element) { ... }
	public E remove() { ... }
	public int size() { ... }
	private E[] elements;
}

public class LinkedListQueue<E> implements Queue<E> {
	private Link head;
	private Link tail;

	LinkedListQueue() { ... }
	public void add(E element) { ... }
	public E remove() { ... }
	public int size() { ... }
}
```

#### 队列与双端队列
有两个端头的队列，即`双端队列`， 可以让人们有效地在头部和尾部同时添加或删除元素。
ArrayDeque 和 LinkedList类实现。这个两个类都提供了双端队列，而且在必要时可以增加队列
的长度。


#### java.util.Queue<E>
* boolean add(E element)
* boolean offer(E element)
	如果队列没有满，将给定的元素添加到这个双端队列的尾部并返回true。
	如果队列满了，第一个方法将抛出一个IllegalStateException, 而第二个方法返回false。

* E remove()
* E poll()
	假如队列不空，删除并返回这个队列头部的元素。如果队列是空的，第一个方法抛出NoSuchElementException,
	而第二个方法返回null。

* E element()
* E peek()
	如果队列不空，返回这个队列头部的元素，但不删除。如果队列空，第一个方法将抛出一个NoSuchElementException,
	而第二个方法返回null.

#### java.util.Deque<E>
* void addFirst(E element)
* void addLast(E element)
* boolean offerFirst(E elment)
* boolean offerLast(E element)
	将给定的对象添加到双端队列的头部或尾部。如果队列满了，前面两个方法将抛出一个IllegalStteException,
	而后面两个方法返回false。

* E remvoeFirst()
* E remvoeLast()
* E pollFirst()
* E pollLast()
	如果队列不空，删除并返回队列头部的元素。如果队列为空，前面两个方法将抛出一个NoSuchElementException,
	而后面两个方法返回null。

* E getFirst()
* E getLast()
* E peekFirst()
* E peekLast()
	如果队列非空，返回队列头部的元素，但不删除。如果队列空，前面两个方法将抛出一个NoSuchElementException,
	而后面两个方法返回null.

#### java.util.ArrayDeque<E>
* ArrayDeque()
* ArrayDeque(int initialCapacity)
	用初始容易16或给定的初始变量构造一个无限双端队列。
