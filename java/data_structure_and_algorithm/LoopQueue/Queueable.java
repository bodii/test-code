package loop_queue;

/**
 * Queueable Interface class
 * 队列接口类
 * @param <E>
 */
public interface Queueable<E> {
    // 向一端压入元素
    void enqueue(E e);

    // 向另一端取出元素
    E dequeue();

    // 获取队首元素
    E getFront();

    // 获取队列长度
    int getSize();

    // 判断队列是否为空
    boolean isEmpty();
}