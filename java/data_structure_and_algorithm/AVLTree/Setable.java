package avl_tree;

/**
 * 集合接口类
 * 
 * @param <E>
 */
public interface Setable<E> {
    /**
     * 向集合中添加元素
     * 
     * @param e 要添加的元素
     */
    void add(E e);

    /**
     * 从集合中删除某个元素
     * 
     * @param e 要删除的元素
     */
    void remove(E e);

    /**
     * 该集合是否包含某个元素
     * 
     * @param e 要比较的元素
     * @return 是否包含
     */
    boolean contains(E e);

    /**
     * 获取当前集合的大小
     * 
     * @return 大小
     */
    int getSize();

    /**
     * 查看当前集合是否为空
     * 
     * @return 是否为空
     */
    boolean isEmpty();
}