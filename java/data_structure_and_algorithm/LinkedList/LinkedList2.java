package linked_list;

/**
 * 链表类2
 * 添加虚拟头节点
 * 
 * @param <E>
 */
public class LinkedList2<E> {

    /**
     * 链表类2 的内部节点类
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

    // 虚拟头节点
    private Node dummyHead;
    // 链表的大小
    private int size;

    // 构造方法
    public LinkedList2() {
        dummyHead = new Node();
        size = 0;
    }

    /**
     * 获取链表的大小
     * 
     * @return int 大小
     */
    public int getSize() {
        return size;
    }

    /**
     * 检测链表是否为空
     * 
     * @return boolean 是否为空
     */
    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 向链表的合法位置添加元素
     * 
     * @param index int 索引
     * @param e 元素
     */
    public void add(int index, E e) {
        ensureIndex(index);

        Node prev = dummyHead;
        for (int i = 0; i < index; i ++) 
            prev = prev.next;

        prev.next = new Node(e, prev.next);
        size ++;
    }

    /**
     * 在链表的头部添加元素
     * 
     * @param e 元素
     */
    public void addFirst(E e) {
        add(0, e);
    }

    /**
     * 在链表的尾部添加元素
     * 
     * @param e 元素
     */
    public void addLast(E e) {
        add(size -1, e);
    }

    /**
     * 获取指定索引位置的元素
     * 
     * @param index int 索引
     * @return 元素
     */
    public E get(int index) {
        ensureIndex(index);

        Node current = dummyHead.next;
        for (int i = 0; i < index; i ++) 
            current = current.next;

        return current.e;
    }

    /**
     * 获取链表的头元素
     * 
     * @return 元素
     */
    public E getFirst() {
        return get(0);
    }

    /**
     * 获取链表的尾部元素
     * 
     * @return 元素
     */
    public E getLast() {
        return get(size -1);
    }

    /**
     * 设置链表指定索引的值
     * 
     * @param index int 索引
     * @param e 元素
     */
    public void set(int index, E e) {
        ensureIndex(index);

        Node current = dummyHead.next;
        for (int i = 0; i < index; i ++)
            current = current.next;

        current.e = e;
    }

    /**
     * 删除链表中指定索引的元素
     * @param index
     * @return 删除的元素
     */
    public E remove(int index) {
        ensureIndex(index);

        Node prev = dummyHead;
        for (int i = 0; i < index; i++) 
            prev = prev.next;
        
        Node delNode = prev.next;
        prev.next = delNode.next;
        delNode.next = null;
        size --;

        return delNode.e;
    }

    /**
     * 删除链表中的头元素
     * 
     * @return 删除的元素
     */
    public E removeFirst() {
        return remove(0);
    }

    /**
     * 删除链表尾部的元素
     * 
     * @return 删除的元素
     */
    public E removeLast() {
        return remove(size -1);
    }

    /**
     * 检测元素是否存在于链表中
     * 
     * @param e 元素
     * @return boolean 是否存在
     */
    public boolean contains(E e) {

        Node current = dummyHead.next;
        // for (int i = 0; i < size; i ++) {
        //     if (current.e.equals(e)) 
        //         return true;
        
        //     current = current.next;
        // }
        while (current != null) {
            if (current.e.equals(e)) return true;

            current = current.next;
        }

        return false;
    }

    /**
     * 检测传入的索引的合法性
     * 
     * @param index int 索引
     */
    private void ensureIndex(int index) {
        if (0 > index || size < index)
            throw new IllegalArgumentException("Failed. Illegal index.");
    }

    /**
     * 将当前链表转换成字符串
     * 
     */
    @Override public String toString() {
        StringBuilder result = new StringBuilder();

        // Node current = dummyHead.next;
        // while( current != null ) {
        //     result.append(current + "-->");
        //     current = current.next;
        // }

        for (Node current = dummyHead.next; current != null; current = current.next)
            result.append(current + "-->");

        result.append("NULL");

        return result.toString();
    }

}