import java.util.EmptyStackException;

import javax.swing.border.EmptyBorder;

/**
 * 链式队列的操作
 */

public class LinkedQueue implements Queue {
	private Node[] qList;
	private int maxLength;
	private int size;

	private int front = 0; // 队头
	private int rear = 0; // 队尾

	public LinkedQueue(int length) {
		if (length <= 0) {
			length = 10;
		}
		// this.qList = new Object[length];
		this.maxLength = length;
	}

	// 向队列未尾添加元素
	public void push(Object e) {
		checkSize();

		Node element = new Node(e);
		if (size > 1) {
			qList[rear - 1].setNext(element);
		}
		qList[rear] = element;
		size++;
		rear++;
	}

	// 从队列头部取出元素
	public Object pop() {
		checkIsEmpty();

		Node topNode = qList[front];
		Object top = topNode.getNodeValue();
		qList[front] = null;

		size--;
		front++;

		return top;
	}

	// 返回队列头部的元素
	public Object peek() {
		checkIsEmpty();

		return (Object) qList[rear].getNodeValue();
	}

	// 判断队列是否为空
	public boolean isEmpty() {
		if (size == 0 || front == rear) {
			return true;
		} else {
			return false;
		}
	}

	// 返回队列的大小
	public int size() {
		return size;
	}

	private void checkSize() {
		if (size + 1 >= maxLength) {
			throw new IndexOutOfBoundsException();
		}
	}

	private void checkIsEmpty() {
		if (isEmpty()) {
			throw new EmptyStackException();
		}
	}
}
