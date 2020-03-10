package avl_tree;

import avl_tree.Mapable;
import avl_tree.AVLTree;

/**
 * AVLTreeMap class
 * avl树的map版本
 * 
 * @param <K>
 * @param <V>
 */
class AVLTreeMap<K extends Comparable<K>, V> implements Mapable<K, V> {
    private AVLTree<K, V> avl;

    public AVLTreeMap() {
        avl = new AVLTree<>();
    }

    /**
     * 获取AVLTree Map的元素个数
     * 
     * @return 节点数
     */
    @Override
    public int getSize() {
        return avl.getSize();
    }

    /**
     * 检测AVLtree Map的节点是否为空
     * 
     * @return 是否为空
     */
    @Override
    public boolean isEmpty() {
        return avl.isEmpty();
    }

    /**
     * 检测AVLTree Map的元素中是否包含要查询的key
     * 
     * @param key
     * @return 是否包含
     */
    @Override
    public boolean contains(K key) {
        return avl.contains(key);
    }

    /**
     * 添加元素
     * @param key
     * @param val
     */
    @Override
    public void add(K key, V val) {
        avl.add(key, val);
    }

    /**
     * 删除元素 
     * 
     * @param key
     * @return
     */
    @Override
    public V remove(K key) {
        return avl.remove(key);
    }

    /**
     * 获取指定key对应的值
     * 
     * @param key
     * @return
     */
    @Override
    public V get(K key) {
        return avl.get(key);
    }

    /**
     * 设置指定元素的key对应的值
     * 
     * @param k
     * @param v
     */
    @Override
    public void set(K k, V v) {
        avl.set(k, v);
    }

}