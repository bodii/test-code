package chapter05;

public class StackG<E> {
	private E[] elements;
	private int size;
	private static final int DEFAULT_INITIAL_CAPACITY = 16;

	@SuppressWarnings("unchecked")
	public StackG() {
		elements = (E[])new Object[DEFAULT_INITIAL_CAPACITY];
	}

	public void push(E e) {
		ensureCapacity();
		elements[size++] = e;
	}

	public E pop() {
		if (size == 0) 
			throw new EmptyStackException();

		@SuppressWarnings("unchecked")
		E result = (E)elements[--size];
		elements[size] = null;
		return result;
	}

	public int size() {
		return size;
	}

	public boolean isEmpty() {
		return size == 0;
	}

	private void ensureCapacity() {
		if (size == elements.length)
			elements = Arrays.copyOf(elements, size * 2 + 1);
	}
}
