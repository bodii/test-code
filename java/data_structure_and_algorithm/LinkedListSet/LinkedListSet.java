package linked_list_set;

import linked_list_set.Setable;
import linked_list_set.LinkedList;

/**
 * 链表集合类
 * 
 * @param <E>
 */
public class LinkedListSet<E> implements Setable<E> {
    private LinkedList<E> linked;

    /**
     * 链表集合类构造函数
     */
    public LinkedListSet() {
        linked = new LinkedList<>();
    }

    /**
     * 在链表集合类中添加元素
     * 
     * @param e 要添加的元素
     */
    @Override
    public void add(E e) {
        if (linked.contains(e))  // 不包含重复元素
            return;

        linked.addFirst(e);
    }

    /**
     * 从链表集合中删除元素
     * 
     * @param e 要删除的元素
     */
    @Override
    public void remove(E e) {
        linked.removeElement(e);
    }

    /**
     * 查看某个元素是否存在于链表集合中
     * 
     * @param e 要查看的元素
     * @return 是否包含
     */
    @Override 
    public boolean contains(E e) {
        return linked.contains(e);
    }

    /**
     * 获取链表集合的大小
     * 
     * @return 大小
     */
    @Override
    public int getSize() {
        return linked.getSize();
    }

    /**
     * 查看链表集合是否为空
     * 
     * @return 是否为空
     */
    @Override
    public boolean isEmpty() {
        return linked.isEmpty();
    }

    /**
     * 将当前链表集合转换成字符串
     * 
     * @return 返回转换后的字符串
     */
    @Override
    public String toString() {
        return linked.toString();
    }
}