package linked_list_queue;

import linked_list_queue.Queueable;

/**
 * 链表队列类
 * 
 * @param <E>
 */
public class LinkedListQueue<E>  implements Queueable<E> {

    /**
     * 链表队列类的内部节点类
     * 也可以声明一个Node类的单独文件
     */
    private class Node {
        public E e;
        public Node next;

        public Node(E e, Node next) {
            this.e = e;
            this.next = next;
        }

        public Node(E e) {
            this(e, null);
        }

        public Node() {
            this(null, null);
        }

        @Override public String toString() {
            return e.toString();
        }
    }

    private int size; // 链表的大小
    private Node head /* 头节点 */, tail /* 尾节点 */;

    /**
     * 构造函数
     */
    public LinkedListQueue() {
        size = 0;
        head = null;
        tail = null;
    }

    /**
     * 获取链表队列的大小
     * 
     * @return int 大小
     */
    @Override public int getSize() {
        return size;
    }

    /**
     * 检测链表队列是否为空
     * 
     * @return boolean 是否为空
     */
    @Override public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 向链表的一端压入元素
     * 尾部压入
     * 
     * @param E e 元素
     */
    @Override public void enqueue(E e) {
        if (tail == null) {
            tail = new Node(e);
            head = tail;
        } else {
            tail.next = new Node(e);
            tail = tail.next;
        }

        size ++;
    }

    /**
     * 从链表另一端取出元素
     * 头部取出
     * 
     * @return E 返回取出的元素
     */
    @Override public E dequeue() {
        nonEmpty();
        
        Node deNode = head;
        head = head.next;
        deNode.next = null;

        // 如果出队后没有其它元素，则尾节点也为空
        if (head == null)
            tail = null;

        size --;
        
        return deNode.e;
    }

    /**
     * 获取链表的头部元素
     * 
     * @return E 元素
     */
    @Override public E getFront() {
        nonEmpty();
         
        return head.e;
    }

    /**
     * 检测链表是否为空，如果为空时则抛出异常
     */
    private void nonEmpty() {
        if (isEmpty())
            throw new IllegalArgumentException("Failed. Queue is empty.");
    }

    /**
     * 将该链表队列转换成字符串
     */
    @Override public String toString() {

        StringBuilder result = new StringBuilder();
        result.append(String.format("Linked List Queue  size = %d\n", size));
        result.append("front ");

        for (Node current = head; current != null; current = current.next)
            result.append(current + "-->");

        result.append("NULL tail");

        return result.toString();
    }

}