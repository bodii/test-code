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
