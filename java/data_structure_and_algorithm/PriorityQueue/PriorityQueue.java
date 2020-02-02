package priority_queue;

import priority_queue.Queueable;
import priority_queue.MaxHeap;
/**
 * 优先队列类
 * @param <E>
 */
public class PriorityQueue<E extends Comparable<E>> implements Queueable<E> {
    private MaxHeap<E> heap;  // 最大二叉堆的优先队列对象

    /**
     * 构造函数，用于实例化一个最大二叉堆对象
     */
    public PriorityQueue() {
        heap = new MaxHeap<>();
    }

    /**
     * 获取优先队列的元素个数
     * 
     * @return 返回元素个数
     */
    @Override
    public int getSize() {
        return heap.size();
    }

    /**
     * 检测优先队列是否为空
     * 
     * @return 是否为空
     */
    @Override 
    public boolean isEmpty() {
        return heap.isEmpty();
    }

    /**
     * 向优先队列中添加元素
     * 
     * @param e 要添加的元素
     */
    @Override
    public void enqueue(E e) {
        heap.add(e);
    }

    /**
     * 查看优先队列的头元素，也就是最大元素
     * 
     * @return 头部元素
     */
    @Override
    public E getFront() {
        return heap.findMax();
    }

    /**
     * 从优先队列中取出头元素
     * 
     * @return 头元素
     */
    @Override
    public E dequeue() {
        return heap.extractMax();
    }
}