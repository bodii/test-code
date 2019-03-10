import java.util.EmptyStackException;

import javax.naming.SizeLimitExceededException;

/**
 * 顺序栈的实现
 */

public class SequenceStack implements Stack {
	private Object[] sList;
	private int size;
	private int maxLength;

	public SequenceStack(int length) {
		if (length <= 0) {
			this.sList = new Object[10];
			this.maxLength = 10;
		} else {
			this.sList = new Object[length];
			this.maxLength = length;
		}
	}

	public void push(Object e) {
		checkSize();

		sList[size] = e;
		size++;
	}

	public Object pop() {
		checkIsEmpty();

		Object top = sList[size];
		sList[size] = null;
		size--;

		return top;
	}

	public Object peek() {
		checkIsEmpty();

		return (Object) sList[size];
	}

	public boolean isEmpty() {
		if (size <= 0) {
			return true;
		} else {
			return false;
		}
	}

	public int size() {
		return size;
	}

	// 检测是否为空,如果是空则抛出异常
	private void checkIsEmpty() {
		if (isEmpty()) {
			throw new EmptyStackException();
		}
	}

	// 检测栈的大小是否越界，如果是则抛出异常
	private void checkSize() {
		if (size >= maxLength) {
			throw new IndexOutOfBoundsException("max length:  " + maxLength + "  size:" + size);
		}
	}
}
