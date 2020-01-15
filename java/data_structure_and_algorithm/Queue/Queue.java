package queue;

import queue.Array;
import queue.Queueable;

public class Queue<E> implements Queueable<E> {
    // 初始化一个队列
    Array<E> queue;

    // 构造函数
    public Queue() {
        queue = new Array<>();
    }

    /**
     * 获取队列的长度
     * @return int 长度
     */
    @Override public int getSize() {
        return queue.getSize();
    }

    /**
     * 判断队列是否为空
     * @return boolean 真假
     */
    @Override public boolean isEmpty() {
        return queue.isEmpty();
    }

    /**
     * 获取当前队列的容量
     * 
     * @return int 容量值
     */
    public int getCapacity() {
        return queue.getCapacity();
    }

    /**
     * 在队列的尾部压入一个元素
     */
    @Override public void enqueue(E e) {
        queue.addLast(e);
    }

    /**
     * 在队列头部取出一个元素
     * @return E 元素
     */
    @Override public E dequeue() {
        return queue.removeFirst();
    }

    /**
     * 获取队列头部元素
     * @return E 元素
     */
    @Override public E getFront() {
        return queue.get(0);
    }

    @Override public String toString() {
        StringBuilder result = new StringBuilder();
        result.append("Queue:  front [");
        for (int i = 0; i < getSize(); i++)
            result.append(queue.get(i) + ", ");
        result.replace(result.length() -2, result.length(), "] tail");

        return result.toString();
    }
}