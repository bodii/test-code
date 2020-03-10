package avl_tree;

import avl_tree.Setable;
import avl_tree.AVLTree;

/**
 * AVLTreeSet class 
 * avl树的set版本
 * 
 * @param <E>
 */
public class AVLTreeSet<E extends Comparable<E>> implements Setable<E> {
    private AVLTree<K, Object> avl;

    public AVLTreeSet() {
        avl = new AVLTree<>();
    }

    /**
     * 向avl tree添加元素
     * 
     * @param e 添加当前元素
     */
    @Override
    public void add(E e) {
        avl.add(e, null);
    }

    /**
     * 从avl tree中删除元素 
     * 
     * @param e 移除当前元素
     */
    @Override
    public void remove(E e) {
        avl.remove(e);
    }

    /**
     * 检测某个元素是否包含指定的元素
     * 
     * @param e 指定查询的元素
     * @return 是否包含
     */
    @Override
    public boolean contains(E e) {
        return avl.contains(e);
    }

    /**
     * 获取当前avl tree的节点个数
     * 
     * @return 元素个数
     */
    @Override
    public int getSize() {
        return avl.getSize();
    }

    /**
     * 查询当前avl tree是否为空
     * 
     * @return 是否为空
     */
    @Override
    public boolean isEmpty() {
        return avl.isEmpty();
    }
}