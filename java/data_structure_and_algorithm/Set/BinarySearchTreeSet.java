package set;

import set.BinarySearchTree;
import set.Setable;

/**
 * 二分搜索树集合类
 * 
 * @param <E>
 */
public class BinarySearchTreeSet<E extends Comparable<E>> implements Setable<E> {
    private BinarySearchTree<E> bst;

    /**
     * 二分搜索树集合类构造函数
     */
    public BinarySearchTreeSet() {
        bst = new BinarySearchTree<>();
    }

    /**
     * 向二分搜索树集合中添加元素
     * 
     * @param e 要添加的元素
     */
    @Override 
    public void add(E e) {
        bst.add(e);
    }

    /**
     * 从二分搜索树集合中删除元素
     * 
     * @param e 要删除的元素
     */
    @Override
    public void remove(E e) {
        bst.remove(e);
    }

    /**
     * 查询二分搜索树集合中是否包含指定的元素
     * 
     * @param e 要查询比较的元素
     */
    @Override
    public boolean contains(E e) {
        return bst.contains(e);
    }

    /**
     * 获取当前二分搜索树集合的大小
     * 
     * @return 大小
     */
    @Override
    public int getSize() {
        return bst.getSize();
    }

    /**
     * 查看当前二分搜索树集合是否为空
     * 
     * @return 是否为空
     */
    @Override
    public boolean isEmpty() {
        return bst.isEmpty();
    }

}