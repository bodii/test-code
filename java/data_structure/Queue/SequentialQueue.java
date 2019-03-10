import java.util.EmptyStackException;

import javax.swing.border.EmptyBorder;

/**
 * 顺序队列的实现操作
 */

public class SequentialQueue implements Queue {

	private Object[] qList;
	private int maxLength;
	private int size;

	private int front; // 队头
	private int rear; // 队尾

	public SequentialQueue(int length) {
		if (length <= 0) {
			this.qList = new Object[10];
			this.maxLength = 10;
		} else {
			this.qList = new Object[length];
			this.maxLength = length;
		}
	}

	// 向队列未尾添加元素
	public void push(Object e) {
		checkSize();

		if (size == 0) {
			front = size;
		}

		qList[size] = e;
		rear = size;
		size++;
	}

	// 从队列头部取出元素
	public Object pop() {
		checkIsEmpty();

		Object frontVal = qList[front];
		qList[front] = null;
		front++;
		size--;
		return frontVal;
	}

	// 返回队列头部的元素
	public Object peek() {
		checkIsEmpty();

		return (Object) qList[front];
	}

	// 判断队列是否为空
	public boolean isEmpty() {
		if (size == 0) {
			return true;
		} else {
			return false;
		}
	}

	// 返回队列的大小
	public int size() {
		return size;
	}

	private void checkIsEmpty() {
		if (isEmpty()) {
			throw new EmptyStackException();
		}
	}

	private void checkSize() {
		if (size >= maxLength) {
			String message = "max length:  " + maxLength + "  size:" + size;
			throw new IndexOutOfBoundsException(message);
		}
	}
}
