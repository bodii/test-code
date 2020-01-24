package loop_queue;

import loop_queue.Queueable;

/**
 * LoopQueue class
 * 循环链表类
 * 
 * @param <E>
 */
public class LoopQueue<E> implements Queueable<E> {

    private E[] data;
    private int front, tail;
    private int size;
    private final int DEFAULT_CAPACITY = 10; // 默认链表的容量

    /**
     * 构造函数
     * 初始化链表信息
     */
    public LoopQueue() {
        data = (E[])new Object[DEFAULT_CAPACITY + 1];
        size = 0;
        front = 0;
        tail = 0;
    }

    /**
     * 获取链表的大小
     * 
     * @return int 大小
     */
    @Override public int getSize() {
        return size;
    }

    /**
     * 获取链表的容量
     * 
     * @return int 容量
     */
    public int getCapacity() {
        return data.length - 1;
    }

    /**
     *  查询链表是否为空
     * 
     * @return boolean 是否为空
     */
    @Override public boolean isEmpty() {
        return front == tail;
    }

    /**
     * 获取队首元素
     * 
     * @return E 获取的元素
     */
    @Override public E getFront() {
        nonEmpty();

        return data[front];
    }

    /**
     * 向循环队列中添加元素
     * 
     * @param E 要添加的元素
     */
    @Override public void enqueue(E e) {
        enableCapacity();

        data[tail] = e;
        tail = (tail + 1) % data.length;
        size ++;
    }

    /**
     * 把循环队列中头部的元素删除
     * 
     * @return E 返回删除的元素
     */
    @Override public E dequeue() {
         nonEmpty();
         enableCapacity();

         E result = data[front];
         data[front] = null;
         front = (front + 1) % data.length;
         size --;

         return result;
    }

    /**
     * 保证容量是否够用
     * 扩容 or  缩容
     */
    private void enableCapacity() {
        // 扩容
        if ((tail + 1) % data.length == front) 
            resize(getCapacity() * 2);

        // 缩容
        if (size == getCapacity() / 4 && getCapacity() / 2 != 0)
            resize(getCapacity() / 2);
    }

    /**
     * 重置循环队列的容量
     * 
     * @param int newCapacity
     */
    private void resize(int newCapacity) {
        E[] newData = (E[]) new Object[newCapacity + 1];
        for (int i = 0; i < size; i++)
            newData[i] = data[(i + front) % data.length];

        data = newData;
        front = 0;
        tail = size;
    }

    /**
     * 检测循环队列是否为空
     * 
     * @throws IllegalArgumentException
     */
    private void nonEmpty() {
        if (isEmpty())
            throw new IllegalArgumentException("Cannot loop queue from an empty queue.");
    }

    /**
     * 将该循环队列对象转换成字符串
     */
    @Override public String toString() {
        StringBuilder result = new StringBuilder();
        result.append(String.format("Loop queue: size = %d, capacity = %d\n", size, getCapacity()));
        result.append("front [");
        for (int i = front; i != tail; i = (i + 1) % data.length) 
            result.append(data[i] + ", ");
        result.replace(result.length() - 2, result.length(), "] tail");

        return result.toString();
    }
}