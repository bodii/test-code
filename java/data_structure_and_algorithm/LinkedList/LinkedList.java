package linked_list;

/**
 * 链表类
 * 
 * @param <E>
 */
public class LinkedList<E> {

    /**
     * 内部节点类
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

    private Node head; // 链表的头节点
    public int size; // 链表的大小

    /**
     * 初始化一个空链表
     */
    public LinkedList() {
        size = 0;
        head = null;
    }

    /**
     * 获取链表的长度
     * 
     * @return int 长度
     */
    public int getSize() {
        return size;
    }

    /**
     * 检测链表是否为空
     * @return boolean 是否是空
     */
    public boolean isEmpty() {
        return size == 0;
    }

    /**
     * 在链表的头部添加元素
     * 
     * @param e 元素
     */
    public void addFirst(E e) {
        // Node node = new Node(e);
        // node.next = head;
        // head = node;
        head = new Node(e, head);
        
        size ++;
    }

    /**
     * 在链表的尾部添加元素
     * 
     * @param e 元素
     */
    public void addLast(E e) {
        add(size - 1, e);
    }

    /**
     * 在链表的指定索引位置添加元素
     * 
     * @param index int 索引
     * @param e 元素
     */
    public void add(int index, E e) {
        ensureIndex(index);

        if (0 == index)
            addFirst(e);
        else {
            Node prev = head;
            for (int i = 0; i < index - 1; i ++) 
                prev = prev.next;
            // Node current = Node(e);
            // current.next = prev.next;
            // prev.next = current;
            prev.next = new Node(e, prev.next);

            size ++;
        }
    }

    /**
     * 检测传入的链表索引是否合法
     * 如果不合法则抛出异常
     * 
     * @param index 
     */
    private void ensureIndex(int index) {
        if (0 > index || size < index)
            throw new IllegalArgumentException("Add failed. Illegal index.");
    }

}